<h3>Кто я?</h3>

Юманов Никита, студент группы ИСдо-46

<h3>Ссылка на задание</h3>

https://docs.google.com/document/d/1-laS0wKfca9m3r0FOBkMI1GuZ6HSyC73/edit

<h3>Описание проекта</h3>

Данный проект полностью написан на Go с использованием Gorilla Mux для работы с маршрутизацией.

<h3>Функционал</h3>

- Создание автомобиля
- Получение списка автомобилей
- Получение автомобиля по ID
- Полное обновление автомобиля
- Частичное обновление автомобиля
- Удаление автомобиля
- Сохранение json в отдельный файл ```cars.json```

<h3>Использование</h3>

1. Склонируйте репозиторий: ```git clone https://github.com/SamuraiAkira/REST_API```
2. Убедитесь, что у вас установлен Go версии 1.22 и нужные зависимости. Для работы с маршрутизацией используется пакет Gorilla Mux.
   <br>Установите его следующим образом: ```go get "github.com/gorilla/mux"```
3. Инициализируйте проект: ```go mod init REST-API``` (опционально)
4. Запустите проект: ```go run cmd/main.go```


<h3>Команды</h3>

Используя инструмент типа curl или любой HTTP-клиент (например, Postman, Insomnia), вы можете взаимодействовать с сервером:

1. Создание автомобиля (POST запрос):
<br>```curl -X POST http://localhost:8080/cars -H "Content-Type: application/json" -d '{"brandName":"VALUE","modelName":"VALUE","mileage":VALUE,"numberOfOwners":VALUE}'```
2. Получение списка автомобилей (GET запрос):
<br>```curl http://localhost:8080/cars```
3. Получение автомобиля по ID (GET запрос):
<br>```curl http://localhost:8080/cars/ID```
4. Полное обновление автомобиля (PUT запрос):
<br>``` curl -X PUT http://localhost:8080/cars/ID -H "Content-Type: application/json" -d '{"brandName":"VALUE","modelName":"VALUE","mileage":VALUE,"numberOfOwners":VALUE}'```
5. Частичное обновление автомобиля (PATCH запрос):
<br>```curl -X PATCH http://localhost:8080/cars/ID -H "Content-Type: application/json" -d '{"mileage":VALUE}'```
6. Удаление автомобиля (DELETE запрос):
<br>```curl -X DELETE http://localhost:8080/cars/ID```

Вместо ID прописываем конкретный необходиый id. Например, "1", "2", "3" и т.д.
<br>Вместо VALUE прописываем конкретное необходимое значение. Например, "BMW", "e30", "20387" и т.д.
<br><b>Примечания:</b> Убедитесь, что при взаимодействии с сервером указываете правильные заголовки, такие как <br>```Content-Type: application/json``` для POST, PUT и PATCH запросов.
