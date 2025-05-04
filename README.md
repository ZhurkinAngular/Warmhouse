# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside>

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удаленно включать и выключать отопление в своих домах
- Пользователи могут выставлять необходимую температуру для отопления

**Мониторинг температуры:**

- Система получает данные о температуре с датчиков, установленных в домах. Пользователи могут просматривать текущую  
температуру дома через веб интерфейс.

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go
- БД: PostgreSQL
- Архитектура: Монолитная, все компоненты системы находятся в рамках одного приложения
- Взаимодействие: Синхронное, запросы обрабатываются последовательно
- Масштабируемость: Ограничена, так как монолит сложно масштабировать по частям
- Развертывание: Требует остановки всего приложения


### 3. Определение доменов и границы контекстов

- **Домен**: Управление устройствами
  - **Поддомен** управления отоплением
    - Контекст: Изменение температуры
  - **Поддомен** проверки температуры
    - Контекст: Отображение температуры

### **4. Проблемы монолитного решения**

- Ограниченная функциональность - работаем только с датчиками температуры и отоплением
- Неудобство добавления новых устройств
- Неэффективная утилизация ресурсов из-за синхронной обработки запросов
- Сложность развертывания - для обновления части системы нужно перезапускать весь сервис
- Низкая отказоустойчивость - упал монолит = упало все и сразу
- Вся коммуникация через один сервер - боттлнек
- Невозможность масштабироваться

### 5. Визуализация контекста системы — диаграмма С4

[Диаграмма контекста](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/context.puml)

# Задание 2. Проектирование микросервисной архитектуры

**Диаграмма контейнеров (Containers)**

[Диаграмма контейнеров](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/containers.puml)

**Диаграмма компонентов (Components)**

- [Диаграмма компонентов микросервиса отопления](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/components_heating.puml)
- [Диаграмма компонентов микросервиса видеонаблюдения](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/components_cameras.puml)
- [Диаграмма компонентов микросервиса дверей и ворот](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/components_doors.puml)
- [Диаграмма компонентов микросервиса света](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/components_light.puml)
- [Диаграмма компонентов микросервиса сценариев](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/components_scenarios.puml)

**Диаграмма кода (Code)**

- [Диаграмма кода ядра микросервиса сценариев](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/code_scenario_engine.puml)
- [Диаграмма потока обработки данных в компоненте консьюмера](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/code_kafka_consumer_flow.puml)

# Задание 3. Разработка ER-диаграммы

[ER-диаграмма](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/diagrams/er_diagram.puml)

# Задание 4. Создание и документирование API

### 1. Тип API

Для взаимодействия юзер - микросервер используется синхронное API, так как оно позволяет немедленно возвращать ответ.  
Для взаимодействия между микросервисами будет использовано асинхронное API (для Kafka).

### 2. Документация API

- [Синхронное API для микросервиса отопления](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/apis/rest/heating-service.openapi.yaml)
- [Синхронное API для микросервиса сценариев](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/apis/rest/scenario-service.openapi.yaml)
- [Асинхронное API для микросервиса отопления](https://github.com/justaleaf/architecture-warmhouse/blob/warmhouse/apis/async/kafka-events.asyncapi.yaml)

# Задание 5. Работа с docker и docker-compose

1) Реализовано приложение на Python + FastAPI для выдачи температуры согласно требуемому формату запроса.

2) Приложение упаковано в Docker и добавлено в docker-compose. Порт выставлен 8081

3) Добавлен в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql
