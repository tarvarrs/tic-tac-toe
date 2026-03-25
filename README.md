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

4. API будет доступен по адресу `http://localhost:8080`. Также можно воспользоваться UI через файл `index.html` или использовать Swagger UI по адресу `http://localhost:8080/swagger`

## Архитектура

```
.
├── cmd
│   └── app
│       └── main.go           # Точка входа и di-контейнер
├── go.mod                    # Управление зависимостями
├── go.sum                    # Контрольные суммы 
├── index.html                # UI
├── internal
│   ├── game                  # Слой бизнес-логики игры
│   │   ├── models.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── http                  # Слой HTTP
│   │   ├── game_handler.go
│   │   └── swagger.go
│   └── storage               # Слой работы с данными
│       └── memory            # Реализация in-memory хранилища
│           ├── game_storage.go
│           └── models.go
├── openapi.yaml              # Спецификация API
└── README.md                 # Описание проекта
```
