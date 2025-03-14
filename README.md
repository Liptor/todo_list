# REST API для управления задачами (TODO-лист)

## Описание
Этот сервис предоставляет REST API для создания и управления задачами. Включает возможности добавления, удаления, изменения и получения массива задач.

## Функциональность
1. **REST API методы:**
   - Получение данных задач.
   - Удаление задач.
   - Изменение данных задачи.
   - Добавление новой задачи в формате JSON:
     ```json
     {
       "title": "Заголовок",
       "description": "Описание",
       "status": "new"
     }
     ```
2. **Хранение данных:**
   - Полученные данные обогащаются и сохраняются в базе данных PostgreSQL.
   - Структура БД создаётся с помощью миграций при старте сервиса.

3. **Логирование:**
   - Код покрыт `debug`- и `info`-логами.
   
4. **Конфигурация:**
   - Все настройки вынесены в `.env`-файл.
   
## Запуск проекта
### 1. Клонирование репозитория
```sh
git clone https://github.com/Liptor/todo_list.git
cd todo_list
```

### 2. Настройка окружения
Создайте `.env` файл в корневой директории и добавьте в него настройки для базы данных и API.

Пример:
```env
DATABASE_URL="postgres://username:password@localhost:5432/todo_list?sslmode=disable"
PORT=3030

```

### 3. Установка зависимостей
```sh
go mod tidy
```

### 4. Запуск миграций
В качестве утилиты для миграций применяется CLI golang/migrate

```sh
migrate -database "postgres://username:password@localhost:5432/todo_list?sslmode=disable" -path internal/database/migration up 
```

### 5. Запуск сервера
Для локального запуска достаточно ввести команду

```sh
go run main.go
```

Чтобы создать Docker Container 

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

docker run -it todo_list
```

## Используемые технологии
- **Go (Golang)** – основной язык разработки
- **Pgx** - в качестве драйвера для доступа к базе данных
- **Fiber** – веб-фреймворк
- **PostgreSQL** – база данных
- **Docker (опционально)** – контейнеризация сервиса


## Скриншоты работоспособности сервиса
Добавление элемента через POST запрос:
![alt text](<images/add_task.png>)
Элемент в базе:
![alt text](<images/add_task_db.png>)

Получение списка всех задач
![alt text](<images/task_list.png>)
Наличие этих задач в БД
![alt text](<images/task_list_db.png>)

Изменение одной задачи
![alt text](<images/change_task.png>)
Измененная задача в БД
![alt text](<images/change_task_db.png>)

Удаление задачи 
![alt text](<images/delete_task.png>)
Задача в БД удалена
![alt text](<images/delete_task_db.png>)