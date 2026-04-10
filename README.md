# Планировщик задач на Go

## Описание проекта

Данный проект представляет собой веб-приложение планировщика задач, реализованное на языке Go.  

Приложение позволяет:
- добавлять задачи;
- редактировать задачи;
- удалять задачи;
- отмечать задачи как выполненные;
- работать с повторяющимися задачами;
- хранить данные в базе SQLite;
- получать список задач через REST API.

Фронтенд уже предоставлен, серверная часть реализована на Go.

## Выполненные задания со звёздочкой

Задания повышенной трудности не выполнялись.

## Локальный запуск проекта

go run main.go

После запуска откройте в браузере:

http://localhost:7540/

## Запуск тестов

Запуск всех тестов:

go test ./tests

Запуск отдельных тестов:

go test -run ^TestApp$ ./tests
go test -run ^TestDB$ ./tests
go test -run ^TestNextDate$ ./tests
go test -run ^TestAddTask$ ./tests
go test -run ^TestTasks$ ./tests
go test -run ^TestTask$ ./tests
go test -run ^TestEditTask$ ./tests
go test -run ^TestDone$ ./tests
go test -run ^TestDelTask$ ./tests

## Docker
Сборка Docker-образа

docker build -t todo-app .

Запуск контейнера

docker run -p 7540:7540 todo-app

После запуска приложение будет доступно в браузере:

http://localhost:7540/