# Tic-Tac-Toe API & Minimax Algorithm

Backend для игры в Крестики-нолики, написанный на Go. Проект реализует ИИ с использованием алгоритма Минимакс, следует чистой архитектуре и принципам внедрения зависимостей через `uber-go/fx`.

## Запуск проекта

1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/tarvarrs/tic-tac-toe.git
    cd tic-tac-toe
    ```

2. Загрузите зависимости:

    ```bash
    go mod download
    ```

3. Запустите сервер:

    ```bash
    go run cmd/app/main.go
    ```

4. API будет доступен по адресу `http://localhost:8000`. Также можно воспользоваться UI через файл `index.html`

## Архитектура

```
.
├── cmd
│   └── app
│       └── main.go               # Точка входа, DI-контейнер и запуск сервера
├── go.mod                        # Управление зависимостями
├── go.sum                        # Контрольные суммы зависимостей
├── index.html                    # UI
├── internal
│   ├── datasource                # Слой работы с данными
│   │   ├── mapper                # Конвертация domain моделей в datasource и обратно
│   │   │   └── mapper.go
│   │   ├── models                # Внутренние модели для хранения в базе (sync.Map)
│   │   │   └── current_game.go
│   │   └── repository            # Реализация in-memory базы данных
│   │       ├── game_repo.go
│   │       └── game_storage.go
│   ├── domain                    # Слой бизнес-логики
│   │   ├── models                # Чистые доменные модели
│   │   │   ├── consts.go
│   │   │   ├── current_game.go
│   │   │   └── grid.go
│   │   └── service               # Бизнес-правила, Минимакс, интерфейсы контрактов
│   │       ├── game_repo.go
│   │       ├── game_service.go
│   │       └── manager.go
│   └── web                       # Слой HTTP
│       ├── handler               # Обработка HTTP-запросов (Echo v5)
│       │   └── game_handler.go
│       ├── mapper                # Конвертация Web DTO в Domain модели и обратно
│       │   └── mapper.go
│       └── models                # Структуры Request/Response
│           └── game_dto.go
├── openapi.yaml                  # Спецификация API
└── README.md                     # Описание проекта
```
