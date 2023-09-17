# EM-Service
Test task for ***Effective mobile*** company.

## User's guide
Write `make help` in the project root to get info about project main commands.

## curl test queries
### REST
* **Get FIO with filters and pagination**
```
curl --location --request GET 'http://localhost:3000/fio' \
--header 'Content-Type: application/json' \
--data '{
    "limit":20,
    "offset":1,
    "minAge":20,
    "order":"DESC"
  }'
  ```
* **Save FIOs to the kafka `FIO` topic**
```
curl --location 'http://localhost:3000/fio' \
--header 'Content-Type: application/json' \
--data '{
    "fios":[
        {
            "name":"Matvey",
            "surname":"Sizov",
            "patronymic":"Alekseevich"
        },
        {
            "name":"Dmitriy",
            "surname":"Moskalenko",
            "patronymic":"Alekseevich"
        }
    ]
}'
```
* **Update already enriched FIO**
```
curl --location --request PATCH 'http://localhost:3000/fio' \
--header 'Content-Type: application/json' \
--data '{
    "id":2,
    "name":"Misha",
    "surname":"Borov",
    "patronymic":"Vasilevich",
    "age":20,
    "gender":"male",
    "country":"RU"
}'
```
* **Delete enriched FIO**
```
curl --location --request DELETE 'http://localhost:3000/fio' \
--header 'Content-Type: application/json' \
--data '{
    "id":2
}'
```
### GraphQL
* **Get FIO with filters and pagination**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"query get($filter: Filter!){\r\n
    getEnrichedFIO(filter: $filter) {\r\n
        id,\r\n
        name,\r\n
        surname,\r\n
        patronymic,\r\n
        age,\r\n
        gender,\r\n
        country\r\n
    }\r\n
}","variables":{
    "filter":{
        "limit":20,
        "offset":1,
        "minAge":20,
        "order":"DESC"
    }
}}'
```
* **Save FIOs to the kafka `FIO` topic**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation create($fios: [Fio!]!) {\r\n
    createFIO(fios: $fios) {\r\n
        name,\r\n
        surname,\r\n
        patronymic\r\n
    }\r\n
}","variables":{"fios":
    [
        {
            "name":"Matvey",
            "surname":"Sizov",
            "patronymic":"Alekseevich"
        },
        {
            "name":"Dmitriy",
            "surname":"Moskalenko",
            "patronymic":"Alekseevich"
        }
    ]
}}'
```
* **Update already enriched FIO**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation update($enrichedFIO: EnrichedFio!) {\r\n
    updateEnrichedFIO(enrichedFIO: $enrichedFIO) {\r\n
        id,\r\n
        name,\r\n
        surname,\r\n
        patronymic,\r\n
        age,\r\n
        gender,\r\n
        country\r\n
    }\r\n
}","variables":{
    "enrichedFIO":{
        "id":2,
        "name":"Misha",
        "surname":"Borov",
        "patronymic":"Vasilevich",
        "age":20,
        "gender":"male",
        "country":"RU"
    }
}}'
```
* **Delete enriched FIO**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation delete($id: Int!) {\r\n
    deleteEnrichedFIO(id: $id) {\r\n
        id,\r\n
        name,\r\n
        surname,\r\n
        patronymic,\r\n
        age,\r\n
        gender,\r\n
        country\r\n
    }\r\n
}","variables":{
    "id":2
}}'
```
