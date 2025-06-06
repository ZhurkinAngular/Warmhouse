# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут только управлять отоплением в доме и проверять температуру
- Система поддерживает синхронные операции

**Мониторинг температуры:**

- Пользователи могут только проверять температуру
- Система поддерживает специальные датчики и реле

### 2. Анализ архитектуры монолитного приложения

монолит на Go с СУБД Postgres. Всё синхронно.

### 3. Определение доменов и границы контекстов

Домен: Система отопления
Контекст: Управление отоплением
(Выделять Мониторинг, не вижу смысла, т.к. мы работаем с одним домиеном Система отопления в который входит проверка температуры и т.д., а также в коде бизнез логике лежит все вместе)

### **4. Проблемы монолитного решения**

- Трудности масштабирования: Монолитные системы сложно масштабировать по отдельным компонентам, что ограничивает гибкость и эффективность при росте нагрузки.
- Сложность поддержки и обновления: Изменения в одной части системы могут требовать переработки всего приложения, что увеличивает риск ошибок и затраты времени.
- Ограниченная гибкость разработки: Разработчики работают с единой кодовой базой, что усложняет внедрение новых технологий или языков программирования для отдельных модулей.
- Медленная доставка новых функций: Внесение изменений требует пересборки и развертывания всей системы, что замедляет выпуск обновлений.
- Высокая связность компонентов: Тесная интеграция компонентов делает их трудноотделимыми и усложняет тестирование и отладку.
- Уязвимость к сбоям: Проблемы в одной части системы могут привести к отказу всего приложения, снижая его надежность.
- Трудности в командной работе: Большая кодовая база может усложнить координацию между командами и снизить производительность разработки.

### 5. Визуализация контекста системы — диаграмма С4

<details>
  <summary>С4 Diagram</summary>
  
  ```
  @startuml
  title Визуализация контекста системы SmartHome — диаграмма С4

  top to bottom direction

  !includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

  Person(user, "User", "Пользователь системы умного дома")
  System_Ext(sensor, "Sensor", "Датчик температуры")
  System(SmartHome, "Backend", "Backend управления температурой")

  Rel(user, SmartHome, "Uses the system")
  Rel(sensor, SmartHome, "Send data to the system")

  @enduml
  ```
</details>

[Диаграмма](https://www.planttext.com?text=RL7DIWCn4BxdAM9F5RHxwSbJiGWUL4fBpsMx6-pYRh9iPb5lhGKzL12aT_4DbjPYxUyhJ5x19t6oQgqYWSnaClFDVDzKDBW9f578881GS7p4ARwQFcOujnvcvXQdve7Z35UuD3SumPdfaSsuwT5FBmyjp80t8a8wLf7WxzqXnnSgpt1ikdGlw6GSdykCWOevADvM02hYVf083m8b6Ti9f1UclaYJa7S0OdtWE8bxMJeFe9EsKoqIJqaG4agUYfnwuBc9hyu24OlGgPRtQw4heNbwuaIk1f5Gg5LLKQoahOWjGLOJYLQoQA7sUA59fa0MduZOtDmJrnMEYA_bG--_z4YJ1IufPumhosvbEHLsMUEQ-aMjeoieQY6rijYDtCdHXnRHt94OC_w9XMkSM5Ket3TTCzXW53SAKk6Xwrq8wHT-N3l1EfTpHBkPq2nsB_zXqcnlXChg8lnYkurCn3M7Zk0wxqg9DkzxzktCXktgk-3QZVsgO1Nw8b4_0000)

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

<details>
  <summary>Container Diagram</summary>
  
  ```
  @startuml
title Container Diagram

top to bottom direction

!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

Person(user, "Пользователь", "Пользователь умного дома")
System_Ext(sensor, "Устройство", "Модуль для умного дома")

System_Boundary(ui, "User Interface") {
    Container(web, "Web Site")
    Container(mobile_app, "Mobile application")
    Container(desctop, "Desktop application")
}

Container(api, "API", "Go", "Работа с сенсорами")
Container(bff, "BFF", "Go", "Маршрутизирует запросы")

Container(user_service, "Сервис управления пользователями", "Управляет добавлением, авторизаций пользователя")
Container(sensor_service, "Сервис управления устройствами", "Управляет добавлением новых устройств, настройками, состоянием")
Container(monitoring_service, "Управление мониторингом", "Собирает и публикует данные телеметрии")

ContainerDb(db, "База Данных", "PostgreSQL", "Хранит данные о пользователях и устройствах")

Rel(sensor, api, "Использует систему")
Rel(user, ui, "Использует систему")
Rel(ui, bff, "Использует систему")

Rel(bff, user_service, "Использует сервис")
Rel(bff, sensor_service, "Использует сервис")
Rel(bff, monitoring_service, "Использует сервис")

Rel(api, db, "Чтение и запись данных")
Rel(user_service, db, "Чтение и запись данных")
Rel(sensor_service, db, "Чтение и запись данных")
Rel(monitoring_service, db, "Чтение и запись данных")

@enduml
  ```
</details>

[Диаграмма](https://www.planttext.com?text=dLLDJzj04BtxLsnp2b9ABZtrH0kbGe8g1Q4UaUCiOLNRQzdhKbHL0YAgIY15bNhe0QcjKDMheQ2QZuI_iFv7VRFEDqb5iEB4EpllpSoRMMSkKdQeOi-rb50kP_FILxRmUSWMX5qFRS-ob0oOagmgbP8UgucGEqf8tx8U2TznungFGvTjAHL4JylbqDwUhGkr5LVZY8SEsBYlPXtfbLU5OuSr-KhmWBlb-IUFAwxjg_MLvR9dHug7C6tCIo-GFY42Iicgy32IVf6eIgoWl-ksliwEz5_ydkaaszSNj2xyPuzb3NsZMp3-qMscp_5mevF2ZBMsWw3UneltgXXnFv8cm4csb-rdkt2wf2SGjGtvCIpdO292S5ndpOcyNUBdClPhThXJZ0KOrb41Mu8OuQRjyC8C-s0nN3srYzky2hStlChMX8B3oBOdgyBb6tOGm6l5B1WMBZIbNjpnh_78GT_WlC2ZjzJ1OU-FbjLtjWDAyLbbYMfzAKt5FtIYJtKR6YGisyC7ShRmtOOw2KfDGT9dg6vk0lHySN60uHXDsCq-usw09KLZKdhM5za-moBHjoGqP3uarVfSrEmDtE-4mudd9n2xw4HAQJGCAi7ocXA2iSdqxTZMDqsMfgjz4AmcFlNhT907tpSbHXR0gKXAECa-uVToOeGX3V8fcYfpA3CwRydKMJCpXMVPONOmXh14sqdVgg_o42LcchbdgcrskOOAygGlb0o5NnyiwcIq3dt1GDWs9NMbQz6f05sk0YMT6f2f8eMUaECK1Ac-wax4ERPR6B53yENwKaAqPnZJeI5Pg1PhT5xq5w3G9wQ_Tk7P0KMjo4ZLGxxsUjdayCj4DnaEHwApEw6va1EfZciHGY2PLUxstXpv2TBVe6cFhLCOJAa1efgi0I3XyXTQF2K8xlb1kn_6W0nWz4XDWFUcjHFGOE-EzJJeyJDq3mP3OMJD6_tR5DOPjxJp0Z457mrqr7IcAt0_ueCeHWj_4Ccu-gScmZN7_Hh-4Vy1)

**Диаграмма компонентов (Components)**

<details>
  <summary>Component Diagram</summary>
  
  ```
@startuml
title Component Diagram

top to bottom direction

!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

ContainerDb(db, "База Данных", "PostgreSQL", "Хранит данные")

Container(users, "Сервис управления пользователями") {
  Component(user_controller, "Controller", "Точка взаимодействия с пользователеми")
  Component(user_business_logic, "BusinessLogicLayer", "Бизнес логика")
  Component(user_repository, "RepositoryLayer", "Чтение и запись данных")
    
  Rel(user_controller, user_business_logic, "Использует")
  Rel(user_business_logic, user_repository, "Использует")
  
  Rel(user_repository,db,"Чтение и запись данных")
}

Container(sensors, "Устройства") {
  Component(sensor_controller, "Controller", "Точка взаимодействия")
  Component(sensor_business_logic, "BusinessLogicLayer", "Добавление/Редактирование/Удаление датчиков")
  Component(sensor_repository, "RepositoryLayer", "Чтение и запись данных")

  Rel(sensor_controller, sensor_business_logic, "Использует")
  Rel(sensor_business_logic, sensor_repository, "Использует")
  
  Rel(sensor_repository, db, "Чтение и запись данных")
}


Container(monitoring, "Мониторинг") {
  Component(monitoring_controller, "MonitoringController", "Точка взаимодействия")
  Component(monitoring_business_logic, "BusinessLogicLayer", "Бизнес логика")
  Component(monitoring_repository, "RepositoryLayer", "Чтение и запись данных")
  
  Rel(monitoring_controller, monitoring_business_logic, "Использует")
  Rel(monitoring_business_logic, monitoring_repository, "Использует")
  
  Rel(monitoring_repository, db, "Чтение и запись данных")
}

@enduml
  ```
</details>

[Диаграмма](https://www.planttext.com?text=hLL1RzCm5BxdLvYUDWcQ2oTE4zsncqOHvofDhSvIOaUsApGXfBK8SM3ID6a73Xo02N4jtQALbRP_uVaVxTajQLYJZg4U4iL-Vj_tllU-kZjA1r9tuyZJJ4UKr4IS24wv9hiiwCWWzZmj4g85QGcjHKpQJD9GCy4zxn7ZOTHjqwwCoA7MYNhc-p9uNUqmVTXjTHMLeU0QiQgXYFq62mFP5lkC9ZJoQq-V74G1rw_swduSA4qbRZKpzcfYI_9gc1ymJkLkQwlTUamgS0O3k8O1WNFycC34V3JlAnWu44ft97tvecvNyDqSk_Z8z0bS_JaAQMKx1xfbIrJs-5T8CM68ntl4l8EPonx265A7SKfW1bCOcnFadc9WOFeO6fjJ-0sZoZPvuv5bvnniqqgN8egeH89Qjd3LVOEf-G0tLiRGge4HuaonpXH-cHvY3nqf5bFCYu_ZNMN5hgCofPgHwB0GkPulDkfsNG-E5XMS8UKrQaijnnYnBt432oh2b3GHYcaXZp2paIsMO3zSLTWeI0cCY5Cqixqq9xdcuvmiE27uQj1ejKl50k2JwIsxWCD9JTy1PHXtKrQBBiN8m-Goq6eFKlKsRof5kH9pMrtOOQANfjbO1ojccH_VW5tkZ6w1-w-6E4V0dtdJG-h35_oqIc-GO-Hq3EU_AXkySA7niaTNrfvOg3KIdYokPvDcMaole8DbujUQgIIfgF3x35MGCx--7cQgdAjYmIqGunqByncftUM6pZX6i0bShXXhcV6tkVQo_KtOBCUouRidXxpP6sWnevB-h1EqrZvh4ikat6UZahp_iT8EvMtyIxq5)

**Диаграмма кода (Code)**

<details>
  <summary>Code Diagram</summary>
  
  ```
@startuml
title Code Diagram

top to bottom direction

!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

class Sensor {
  	+int id
  	+string Name
  	+string Location
  	+string Description
  	+string Type
  	+string Status
  	+string UpdatedAt
  	+string CreatedAt
}

package SensorController {
  class Handler {
      +List(request)
      +Get(request)
      +Create(request)
      +Update(request)
      +Delete(request)
  }
}

package SensorService {
  class Service {
      +List(offset int, limit int)
      +Get(id int)
      +Create(sensor object)
      +Update(sensor object)
      +Delete(id int)
  }
}

package SensorRepository {
  class Repository {
      +List(offset int, limit int)
      +Get(id int)
      +Create(sensor object)
      +Update(sensor object)
      +Delete(id int)
  }
}

SensorController --> SensorService
SensorService --> SensorRepository
SensorRepository ..> Sensor

@enduml
  ```
</details>

[Диаграмма](https://www.planttext.com?text=pPEzJiCm481d_1Hc0vKc2nC30gKI3AL2BPsHQr_J0yTdx0keGdrtqYPLagOFWAVpvx_lVFP3P1MuBAnWP0io8mDoWYeFgX22oKickI9cAgJ109gHd12NwBGj3PJ1oWspZtTf6jHFaYDlobKP8MXo38uJJKKwHws2eHc21vjcj-DNgnmlNwPfeI93gD1xHeKdjr_XzpP2MnMZN82B5EIlaF9YX8ubcaCOEQ3BvKmLqErFIQk3O8TD8Eg0_XI_RNrlwO8Lbx5Bbjue1lF8NPW5QE1E2A_qfygXiSogZ0DP2xLldS2pSkP8zcqqnSXN0RvAY7nzX4ymPFL90rnR3V04BFJnRcYuWF2D6ZfwNTBgqNeTeRfhnpVIOe67i2UBfeSQrrXNYrOVrJivDJu_s7Yt-vtHdeEdY4nXsp4_WVz8Vl0QnkFxVW54lnpjU9kK6AIU9CTPGZo0CzKl-GC0)

# Задание 3. Разработка ER-диаграммы

<details>
  <summary>ER Diagram</summary>
  
  ```
@startuml
title ER Diagram

top to bottom direction

!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

entity User {
    **id: number**
    --
    email: string
    password: string
    name: string
    phone_number: string
    updated_at: timestampz
    create_at: timestampz
}
entity House {
    **id: number**
    --
    +user_id: number (FK)
    address: string
}

entity Sensor {
    **id: number**
    --
    +house_id: number (FK)
    type: string
  	name: string
  	description: string
  	location: string
  	status: string
  	created_at: timestampz
  	updated_at: timestampz
}

entity TelemetryData {
    +sensor_id: number (FK)
    created_at: timestampz
    value: float
}

User ||--o{ House
House ||--o{ Sensor
Sensor ||--o{ TelemetryData


@enduml
  ```
</details>

[Диаграмма](https://www.planttext.com?text=XLBBRi8m43pZht1lRL38fQUSaA1LfJvKqNAEbdWBbloIlIwYmB_NIIW51C8NMpFsxioiXu70KzIAaII5_672nnBc7ZHZP1qdosUMo6ekfCUQf3MCtKXJgoWmUiKNH2wKHU5XEPXBMiHP3EXhQmWD3MghYucimGlx9j6XAaRt-Ri2GzFNbq938FG9gaPMEskQ5wxHmj99qefFKoc-PZojB9EYv2Rg6VeiQw4yRpVK85N90tbfvYtY88IbzU88DA3n-DOYTQowYaT4T088HGLKSf8QKqJQ_RHKxJ5HfypsJ--JJUQl2kut2LN_BBzzVBvh6H32OmXxDTjz4XzeWhsUHN_HA3XRcrRkq7_l98-Um51xwPe97yBArd2A9TiKmo7IvN8cijw5CF-TVQ92ZUHNOo3O6Uo7rkvP7nTRSVuDAYPFNye2DHtQxxFPvBbTTwDXtO1sK9SfsqMx0u_aCCQ6Q4Jwahy0)

# Задание 4. Создание и документирование API

### 1. Тип API

Для взаимодействия мобильного приложения и backend можно использовать gRPC, gRPC обладает такими плюсами: streaming, http2, нативность. Но т.к. мы будем использовать несколько разных Fronted-ов, буду использовать REST API, как более универсальное решение.

### 2. Документация API

<details>
  <summary>Swagger</summary>

```yaml
basePath: /
definitions:
  main.HTTPError:
    properties:
      code:
        type: integer
      fields: {}
      message:
        type: string
    type: object
  main.Sensor:
    properties:
      createdAt:
        type: string
      id:
        type: integer
      location:
        type: string
      name:
        type: string
      status:
        type: string
      type:
        type: string
      updatedAt:
        type: string
    type: object
host: localhost:8080
info:
  contact:
    email: test@test.test
    name: OWNER
    url: test.ru
  description: This is REST API.
  title: Swagger API
  version: "1.0"
paths:
  /sensors:
    get:
      consumes:
      - application/json
      description: List sensors
      parameters:
      - description: Offset
        in: query
        name: offset
        required: true
        type: integer
      - description: Limit
        in: query
        name: limit
        required: true
        type: integer
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            items:
              $ref: '#/definitions/main.Sensor'
            type: array
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/main.HTTPError'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/main.HTTPError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/main.HTTPError'
      summary: List sensors
      tags:
      - sensor
    post:
      consumes:
      - application/json
      description: Create sensor
      parameters:
      - description: Sensor object
        in: body
        name: data
        required: true
        schema:
          type: object
      produces:
      - application/json
      responses:
        "201":
          description: Created
          schema:
            $ref: '#/definitions/main.Sensor'
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/main.HTTPError'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/main.HTTPError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/main.HTTPError'
      summary: Create sensor
      tags:
      - sensor
  /sensors/:id:
    delete:
      consumes:
      - application/json
      description: Delete sensor by ID
      parameters:
      - description: ID of sensor
        in: path
        name: id
        required: true
        type: integer
      produces:
      - application/json
      responses:
        "204":
          description: No Content
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/main.HTTPError'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/main.HTTPError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/main.HTTPError'
      summary: Delete sensor
      tags:
      - sensor
    get:
      consumes:
      - application/json
      description: Get sensor by ID
      parameters:
      - description: ID of sensor
        in: path
        name: id
        required: true
        type: integer
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/main.Sensor'
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/main.HTTPError'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/main.HTTPError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/main.HTTPError'
      summary: Get sensor
      tags:
      - sensor
    put:
      consumes:
      - application/json
      description: Update sensor
      parameters:
      - description: ID of sensor
        in: path
        name: id
        required: true
        type: integer
      - description: Sensor object
        in: body
        name: data
        required: true
        schema:
          type: object
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/main.Sensor'
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/main.HTTPError'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/main.HTTPError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/main.HTTPError'
      summary: Update sensor
      tags:
      - sensor
swagger: "2.0"
```
</details>

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


