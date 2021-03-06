# Car CRUD Service
CRUD (REST API) для модели автомобиля

## Запуск
Сервис можно собрать и запустить с помощью `make` (см. [Makefile](Makefile)).

Сборка cервиса:
```bash
make build
```
Запуск сервиса:
```bash
make run
```

## Запуск в Docker
Так же сервис можно запустить в Docker контейнере

Сборка образа:
```bash
make docker-build
```

Запуск контейнера:
```bash
make docker-run
```

## Rest API
Сервис принимает запросы на 5 эндпоинтов:
+ Создать автомобиль
+ Получить автомобиль по его UUID
+ Получить срез автомобилей по любому из полей
+ Обновить любое из полей автомобиля
+ Удалить автомобиль

По умолчанию сервер запускается на порту 80, но порт можно поменять в [configs](deployment/configs.toml) фаиле
### Создать автомобиль 
Для создание автомобиля необходимо отправить `POST` запрос на эндпоинт */car/create*

Пример тела запроса:

```json
{
  "UUID": "ad17237d-a706-467f-bc2c-e76954421c",
  "Brand": "Toyota",
  "Model": "land-cruiser-prado",
  "Price": 3000000,
  "State": "in-stock",
  "Mileage": 10000
}
```

Если поле **UUID** придёт пустым, то uuid автомобиля сгенерируется автоматически 
с помощью библиотеки [uuid](https://github.com/google/uuid)

В ответ вернёт структуру с автомобилем в формате *json*

### Получить автомобиль по его UUID
Чтобы получить автомобиль по его uuid, необходим отправить `GET` запрос на эндпоинт */car/:uuid*

Например:
```
/car/ad17237d-a706-467f-bc2c-e76954421c
```

### Получить срез автомобилей по любому из полей
Для получения среза автомобиля по любому из полей необходимо отправить `GET` запрос на эндпоинт */car*

Например, для того, чтобы получить срез всех автомобилей марки Nissan, необходима в теле запроса отправить:
```json
{
  "Brand": "Nissan"
}
```

В ответ придёт массив структур с найденными автомобилями, либо null, если автомобили по фильтру не были найдены.

### Обновить любое из полей автомобиля
Для обновления информации об автомобиле, отправляется `PUT` запрос на эндпоинт */car/update*
В теле запроса отправляется uuid автомобили, и все поля которые необходимо обновить.

Пример тела запроса, для обновления статуса автомобиля:
```json
{
  "UUID": "ad17237d-a706-467f-bc2c-e76954421c",
  "State": "on-way"
}
```

### Удалить автомобиль
Чтобы удалить автомобиль, отправляется `DELETE` запрос на эндпоинт */car/delete/:uuid*

Пример:
```
/car/delete/ad17237d-a706-467f-bc2c-e76954421c
```
