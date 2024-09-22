##### _____________________________________

## Приложение:

# eShop

##### _____________________________________

eShop — это серверное приложение для управления заказами, продуктами, налогами и поставщиками в контексте магазина (продуктовые магазины, супермаркеты). Оно предоставляет API для работы с товарами, заказами, возвратами, генерацией штрих-кодов и многим другим.

##### _____________________________________

## Функциональные возможности
- Управление товарами (добавление, обновление, удаление, просмотр)
- Работа с заказами (создание заказа, добавление товаров в заказ, удаление товаров)
- Генерация и обработка штрих-кодов
- Поддержка взвешенных товаров с 18-значными штрих-кодами
- Управление поставщиками и категориями товаров
- Ведение учёта налогов и отчетности
- Авторизация и аутентификация на основе JWT

##### _____________________________________

## Требования
- Go 1.16+
- PostgreSQL
- Redis (для кэширования JWT токенов, если используется)
- Swagger (для документирования API)
- Excelize (для работы с Excel файлами, если отчеты включены)
- GORM (для работы с базой данных)
- Docker (**опционально**, *если используете контейнеры для запуска приложения*)

##### _____________________________________

## Стек технологий
* Язык программирования: Go (Golang)
* Веб-фреймворк: Gin
* База данных: PostgreSQL
* ORM: GORM (для взаимодействия с базой данных)
* Миграции базы данных: GORM AutoMigrate
* Аутентификация: JWT (JSON Web Tokens)
* Документация API: Swagger
* Логирование: Стандартное логирование Go (log), плюс библиотека Lumberjack для ротации логов
* Обработка штрих-кодов: Внутренние утилиты для генерации и разбора штрих-кодов
* Работа с файлами Excel: Библиотека excelize для работы с Excel-файлами (для отчетов)
* Работа с архивами: Библиотека archive/zip для обработки ZIP-архивов (для отчетов)
* Хеширование паролей: Использование SHA-256 для хеширования данных (при необходимости можно поменять на более сильные алгоритмы, такие как bcrypt или Argon2)

      ---* Этот стек технологий отражает всё ключевое, используемое в проекте. ​

##### _____________________________________

## Установка

1. Склонируйте репозиторий:

```bash
git clone https://github.com/AndreyBorisovichKoval/eShop.git
cd eShop
```

2. Настройте зависимости:

```bash
go mod tidy
```

3. Приложение автоматически создаёт базу данных и выполняет миграции при первом запуске в случае, если не было создано. Убедитесь, что параметры базы данных указаны правильно в файле конфигурации.

4. Настройте файл конфигурации:

Скопируйте файл `configs/configs.json.example` в `configs/configs.json` и заполните необходимые параметры для подключения к базе данных.

5. Все необходимые таблицы также могут быть созданы автоматически при первом запуске. Дополнительно можно провести миграции:

```bash
go run cmd/app.go migrate
```

##### _____________________________________

## Заполнение тестовыми данными (опционально):

1. Запустите сервер.

2. Воспользуйтесь эндпоинтом для вставки тестовых данных:

```bash
curl -X POST http://localhost:8585/insert-test-data
```

##### _____________________________________

## Запуск приложения

1. Запустите сервер:

```bash
go run main.go
```

2. Приложение будет доступно по умолчанию по адресу: `http://localhost:8585`. 
   Порт можно изменить в файле конфигурации `configs.json`.

##### _____________________________________

## API документация

Документация API доступна по адресу:

```
http://localhost:8585/swagger/index.html
* Порт можно изменить в файле конфигурации `configs.json`.
```

##### _____________________________________

## Тестирование

Вы можете использовать Postman для тестирования API. Коллекция запросов находится в файле:

```
eShop.postman_collection.json
```

##### _____________________________________

## Структура проекта

```
├── cmd                # Точка входа для запуска приложения
├── configs            # Файлы конфигурации
├── db                 # Подключение к базе данных и миграции
├── docs               # Документация Swagger
├── errs               # Пользовательские ошибки
├── logger             # Логирование
├── logs               # Логи приложения
├── models             # Определение моделей базы данных
├── pkg                # Основная логика приложения
│   ├── controllers    # Контроллеры (обработка HTTP-запросов)
│   ├── service        # Бизнес-логика
│   ├── repository     # Логика работы с базой данных (CRUD операции)
├── server             # Настройка сервера
└── utils              # Утилитарные функции
```

##### _____________________________________

## Лицензия

Этот проект распространяется под лицензией MIT.

##### _____________________________________

## Readme.md

* Файл __README.md__ - это текущий файл, который может быть изменен по мере необходимости, включая сам **"код"**.
* В процессе разработки и возможного обновления проекта, **README.md** может подвергаться изменениям, чтобы отражать ***актуальную информацию о проекте***, его ***функциональности***, ***инструкции по установке*** и ***использованию***, а также другую полезную информацию для **разработчиков** и **пользователей**.
* Это важный компонент документации, который помогает улучшить понимание проекта и содействует его успешному использованию.

##### _____________________________________


### Структура базы данных:

![images\models.jpg](images\models.jpg)

##### _____________________________________

**Create** by: **Andrey Koval** (57)

**Date**: 2024-09-25

**Version**: 01.00

##### _____________________________________
