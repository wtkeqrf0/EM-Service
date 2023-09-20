# EM-Service
Test task for ***Effective mobile*** company.

## User's guide
* Write `make help` in the project root to get info about project main commands.
* To activate *debug mode*, set `PROD` variable to `false` in [.env](configs/.env) file.

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
--data '{"query":"query get($req: GetEnrichedFioRequest!){\r\n    getEnrichedFio(req: $req) {\r\n        users {\r\n            id\r\n            name\r\n            surname\r\n            patronymic\r\n            age\r\n            gender\r\n            country\r\n        }\r\n    }\r\n}","variables":{"req":{"filter":{"limit":20,"offset":1,"minAge":20,"order":"DESC"}}}}'
```
* **Save FIOs to the kafka `FIO` topic**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation create($req: CreateFioRequest!) {\r\n    createFio(req: $req) {\r\n        failedFIOs {\r\n            name\r\n            surname\r\n            patronymic\r\n        }\r\n    }\r\n}","variables":{"req":{"FIOs":[{"name":"Matvey","surname":"Sizov","patronymic":"Alekseevich"},{"name":"Dmitriy","surname":"Moskalenko","patronymic":"Alekseevich"}]}}}'
```
* **Update already enriched FIO**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation update($req: UpdateEnrichedFioRequest!) {\r\n    updateEnrichedFio(req: $req) {\r\n        user {\r\n            id\r\n            name\r\n            surname\r\n            patronymic\r\n            age\r\n            gender\r\n            country\r\n        }\r\n    }\r\n}","variables":{"req":{"enrichedFio":{"id":10,"name":"Dmitriyesfasfb","surname":"Borov","patronymic":"Vasilevich","age":20,"gender":"male","country":"RU"}}}}'
```
* **Delete enriched FIO**
```
curl --location 'http://localhost:3000/query' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation delete($req: DeleteEnrichedFioRequest!) {\r\n    deleteEnrichedFio(req: $req) {\r\n        user {\r\n            id\r\n            name\r\n            surname\r\n            patronymic\r\n            age\r\n            gender\r\n            country\r\n        }\r\n    }\r\n}","variables":{"req":{"id":2}}}'
```

**P.S. Updates were added after the end of the deadline, because I've been working on this project all the last week. My weekend only started on this Monday :D**
