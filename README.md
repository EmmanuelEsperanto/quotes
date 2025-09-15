# 📚 Quotes API

Простой REST API на Go для работы с цитатами (in-memory хранилище). Поддерживает добавление, получение, удаление и фильтрацию цитат по автору.

---

## 📁 Структура проекта

```
quotes/
├── cmd/
│   └── apiserver/         # Точка входа (main.go)
├── internal/
│   └── app/
│       ├── apiserver/     # HTTP-обработчики и логика сервера
│       ├── model/         # Модели (Quote и тесты)
│       └── store/
│           ├── mainstore/ # Основная in-memory реализация
│           └── teststore/ # Тестовая in-memory реализация
├── .gitignore
├── go.mod
└── Makefile
```

---

## 🚀 Как запустить

### ⚙️ Без Docker

#### Требования:
- Go 1.24
- Git

#### Запуск

```bash
git clone https://github.com/EmmanuelEsperanto/quotes
cd quotes
go mod tidy
go run cmd/apiserver/main.go
```

По умолчанию сервер стартует на `localhost:8080`.

---

#### Сборка и запуск

```bash
docker buildx build -t quotes-api .
docker run -p 8080:8080 quotes-api
```

---

## ✅ Запуск тестов

```bash
go test ./...
```

---

## 📬 API Эндпоинты

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/quotes` | Добавить новую цитату |
| `GET` | `/quotes` | Получить все цитаты |
| `GET` | `/quotes/random` | Получить случайную цитату |
| `GET` | `/quotes?author={author}` | Получить цитаты по автору |
| `DELETE` | `/quotes/{id}` | Удалить цитату по ID |

---

## 🔁 Примеры запросов

### Добавить цитату

```bash
curl -X POST http://localhost:8080/quotes      -H "Content-Type: application/json"      -d '{"author":"Seneca", "quote":"While we teach, we learn."}'
```

### Получить все цитаты

```bash
curl http://localhost:8080/quotes
```

### Получить случайную цитату

```bash
curl http://localhost:8080/quotes/random
```

### Получить цитату по автору

```bash
curl http://localhost:8080/quotes?author=NAME
```

### Удалить цитату

```bash
curl -X DELETE http://localhost:8080/quotes/1
```

---

## 🛠 Подсказки по разработке

- `teststore` используется для изолированного unit-тестирования без привязки к реальному хранилищу.
- Используется `in-memory` хранилище, данные очищаются при перезапуске.
- Структура проекта разделена по зонам ответственности: HTTP-слой, модели, хранилище.
