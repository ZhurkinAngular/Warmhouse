# Теплый дом

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удалённо включать/выключать отопление в своих домах
- Пользователи не могу самостоятельно подключить свой датчик к системе пользователь не может
- Система поддерживает изменение параметров термостата на основе запросов от пользователей

**Мониторинг температуры:**

- Пользователи могут просматривать текущую температуру в своих домах через веб-интерфейс
- Пользователи не могу самостоятельно подключить свой датчик к системе пользователь не может
- Система поддерживает получение данных о температуре с датчиков, установленных в домах

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go
- База данных: PostgreSQL
- Архитектура: Монолитная, все компоненты системы (обработка запросов, бизнес-логика, работа с данными) находятся в рамках одного приложения
- Взаимодействие: Синхронное, запросы обрабатываются последовательно

### 3. Определение доменов и границы контекстов

1. Домен: Smart Home Control (Управление умным домом) - Весь функционал отопления и температуры реализован в одном логическом блоке
1.1. Поддомен: Heating Management (Управление отоплением)
     Контекст:
     - Вкл / выкл отопления
	 - Сохранение состояния
	 - Связь с устройством
1.2. Temperature Monitoring (Мониторинг температуры)
     Контекст:
	 - Запрос показаний с датчика
	 - Запись в базу

2. Домен: Device (Устройства) - Обособленный модуль параметров устройств и обновления ПО
     Контектст:
     - Параметры устройства (Версия ПО, Характеристики)
     - Расположение
     - Обновление ПО

3. Домен: User Management (Управление пользователями) - Обособленный модуль с логикой авторизации и хранения данных пользователя
   Контектст:
   - Регистрация
   - Авторизация
   - Управление аккаунтом
   - Просмотр устройств

4. Домен: Web Interface (Веб-клиент) - Веб-интерфейс для взаимодействия с серверной частью
4.1. Поддомен: Web Interface (Параметры устройств) - Пользовательский Веб-интерфейс
     Контектст:
     - Отображение текущих параметров отопления
     - Отображение текущих параметров температуры
     - Интрфейс Регистрации / Авторизации / Управления аккаунтом
4.2. Поддомен: Admin Panel (Параметры устройств) - Административный Веб-интерфейс
     Контектст:
	 - Управление ролями пользоателей
     - Добавление / редактирование / удаление устройств

### **4. Проблемы монолитного решения**

- Взаимодействие: Синхронное, запросы обрабатываются последовательно
- Масштабируемость: Ограничена, так как монолит сложно масштабировать по частям
- Высокий риск ошибок: Изменения в одной части приложения могут непредсказуемо влиять на другие части
- Длительные циклы разработки и развёртывания: При каждом изменении приходится тестировать всё приложение целиком. Это замедляет выпуск новых функций
- Трудно управлять командой: В больших командах работа над монолитом часто приводит к конфликтам и задержкам. Изменения, которые вносит одна команда, влияют на работу других команд

### 5. Визуализация контекста системы — диаграмма С4

- [Схема Context](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_context.png)

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

- [Схема C4 Containers](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Containers.png)

**Диаграмма компонентов (Components)**

- [Схема Components Device Service](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_Device_Service.png)
- [Схема Components Heating Service](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_Heating_Service.png)
- [Схема Components Notification Service](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_Notification_Service.png)
- [Схема Components Temperature Service](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_Temperature_Service.png)
- [Схема Components User Service](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_User_Service.png)

**Диаграмма кода (Code)**

Возьмём самый критичный микросервис — Heating Service, так как он содержит бизнес-логику управления отоплением, взаимодействует с устройствами и хранит состояние.
- [Схема Components_Heating_Service_sequence](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_Heating_Service_sequence.png)
- [Схема Components_Heating_Service_class](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_Components_Heating_Service_class.png)

# Задание 3. Разработка ER-диаграммы

- [Общая ER-диаграмма](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/schemas/C4.Visualize_system_ER.png)

# Задание 4. Создание и документирование API

### 1. Тип API

REST API. Он хорошо подходит для синхронного взаимодействия, когда клиенту (в данном случае — другому микросервису) нужен немедленный ответ на запрос
- Все сервисы общаются по синхронным HTTP(S) запросам
- Запросы — CRUD операции, команды управления, получение статусов и данных
- Ошибки и ответы обрабатываются сразу, клиент ждёт подтверждения
- Логирование и мониторинг проще — трассировка цепочек вызовов через HTTP
- Упрощается отладка и тестирование

### 2. Документация API

- [OpenAPI](https://github.com/Mkuzya/architecture-warmhouse/blob/main/apps/smart_home/api/OpenApi.yaml)

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


