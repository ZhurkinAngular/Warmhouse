Это шаблон для решения **первой части** проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании условия задания. Это нормально.

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут включать/выключать отопление
- Пользователи могут регулировать температуру отопления

**Мониторинг температуры:**

- Система поддерживает актуальные показания с датчиков/приборов
- Пользователи могут получать текущие показания с датчиков (термометров)

### 2. Анализ архитектуры монолитного приложения

Используется ЯП Java, БД Postgres. Домены реализуются единым компонентом (монолитная архитектура), взаимодействие между ними – синхронное.

### 3. Определение доменов и границы контекстов

- Управление устройствами (отопление в данном случае – тоже устройство)
  - Включить/выключить
  - Установка желаемых параметров (например, установка конкретной температуры)

- Мониторинг устройств
  - Получение текущих значений датчиков устройств (например, текущей температуры)

### **4. Проблемы монолитного решения**

В целом актуальны все минусы любого монолита, основные:
- Сложность масштабирования (по сути доступно только вертикальное)
- Сложности и риски (поломки приложения целиком) при добавлении новой бизнес-логики в отдельные домены
- Сложность релиза обновленных версий отдельных частей приложения (например, мониторинга)

### 5. Визуализация контекста системы — диаграмма С4


@startuml
!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

Person(user, "User")
Container(publicInterface, "Public interface")

System(service, "Monolith Heating System Service")
ContainerDb(database, "Service's Database", "Postgres", "Stores info about houses")
System_Ext(devices, "Physical Devices", "External devices placed in houses")

Rel(user, publicInterface, "Change devices properties and monitor their sensors")
Rel(publicInterface, service, "Generate and send to service requests for user's actions", "REST API, HTTP")
Rel(service, database, "CRUD service-important info", "TCP, SSL")
Rel(service, devices, "Operate external devices", "HTTP")
@enduml


# Задание 2. Проектирование микросервисной архитектуры

**Диаграмма контейнеров (Containers)**


@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
Person(user, "User")

Container(mobileApp, "Mobile App")
Container(apiGateway, "API Gateway")

System_Boundary(SmartHomeSystem, "Smart Devices System") {
    Container(devicesService, "Devices Service", "Controls external physical devices")
    Container(userService, "User Service", "Controls users")
    Container(deviceMonitoringService, "Devices Sensors Monitor", "Service for monitoring sensors")
    Container(notificationService, "Notification Service", "Provides users notificating")

    ContainerDb(devicesServiceDB, "Devices DB", "Postgres", "Stores info about devices")
    ContainerDb(userServiceDB, "Users DB", "Postgres", "Stores info about users")
    ContainerDb(deviceMonitoringServiceDB, "Sensors Monitoring DB", "Postgres", "Stores info about sensors events")
    ContainerDb(notificationServiceDB, "Notifications DB", "Postgres", "Stores info about notifications")

    Container(queue, "Message Queue", "Kafka")
}

Rel(user, mobileApp, "Interaction")
Rel(mobileApp, apiGateway, "Send request")
Rel(apiGateway, devicesService, "Redirect")
Rel(apiGateway, userService, "Redirect")
Rel(apiGateway, deviceMonitoringService, "Redirect")
Rel(apiGateway, notificationService, "Redirect")

Rel(devicesService, devicesServiceDB, "CRUD")
Rel(userService, userServiceDB, "CRUD")
Rel(deviceMonitoringService, deviceMonitoringServiceDB, "CRUD")
Rel(notificationService, notificationServiceDB, "CRUD")

Rel(devicesService, queue, "Publish")
Rel(deviceMonitoringService, queue, "Publish")
Rel(notificationService, queue, "Subscribed")

Container_Ext(sensor, "Sensors")
Rel(sensor, devicesService, "Publish events")

@enduml


**Диаграмма компонентов (Components)**


@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

Container_Boundary(generalService, "Devices Controlling Service") {
    Component(apiGateway, "API Gateway")
    Component(controller, "Devices Controller")
    Component(service, "Devices Service", "Handles external physical devices")
    Component(database, "Devices Repository", "Handles high-level requests from Devices Service and direct it as database queries")
    Component(consumer, "Message Queue Consumer", "Receive updates from sensors")
}

Rel(apiGateway, controller, "requests of users")
Rel(controller, service, "operations on devices")
Rel(service, database, "CRUD")
Rel(service, consumer, "Service consumes (reads) messages")
@enduml


**Диаграмма кода (Code)**


@startuml
class DevicesService {
    + linkDevice(id: Uint): void
    + updateDeviceConfig(deviceId: Uint, config: DeviceConfig): void
    + enableDevice(deviceId: Uint): void
    + disableDevice(deviceId: Uint): void
    + unlinkDevice(deviceId: Uint): void
    + listDevices(): List<Device>
}

class DevicesServiceRepository {
    + findById(deviceId: Uint): Device
    + save(device: Device): Device
}

class EventsConsumer {
    + subscribe(deviceID: Uint, eventType: String): void
}

class EventsProducer {
    + produce(event: SensorEvent): void
}

class Device {
    + Uint id
    + String name
    + boolean isEnabled
    + DeviceConfig config
    + List<SensorEvent> generatedDeviceEvents
    + updateConfig(cfg: DeviceConfig): void
}

struct DeviceConfig {
    + String asJson
}
Device --> DeviceConfig

struct SensorEvent {
    + Uint id
    + Uint deviceID
    + String sensorType
    + String eventType
    + Timestamp time
}

DevicesService --> DevicesServiceRepository : CRUD
DevicesService --> Device : Operate
Device --> SensorEvent : Generate and send event
SensorEvent --> EventsProducer: Event is sent
EventsProducer --> EventsConsumer: Message Queue
DevicesService --> EventsConsumer : Subscribe
EventsConsumer --> DevicesService : Send event on subscription
@enduml


# Задание 3. Разработка ER-диаграммы


@startuml
entity User {
    * id : Uint
    --
    + username : String
    + password : String
    + personalData : PersonalData
}

entity Device {
    * id : Uint
    --
    + name : String
    + type : DeviceType
    + config : DeviceConfig
    + houseID : Uint
}

entity DeviceType {
    * id : Uint
    --
    + name : String
    + configTemplate : DeviceConfig
}

entity DeviceController {
    * id : Uint
    --
    + device_id : Uint
    - config : DeviceConfig
    - isSubscribed : boolean
}

entity DeviceCommand {
    * id : Uint
    --
    + deviceID : Uint
    + name : String
    + newDeviceConfig : DeviceConfig
}

entity DeviceEvent {
    * id : Uint
    --
    + deviceID : Uint
    + value : String
    + config : DeviceConfig
    + time : Timestamp
}

User }o--o{ Device
Device ||--o| DeviceType
Device ||--o| DeviceController
Device ||--o{ DeviceCommand
DeviceEvent ||--o{ DeviceController
User ||--o{ DeviceEvent
DeviceCommand ||--o{ DeviceController
@enduml

