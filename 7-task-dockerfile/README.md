### Book repository

* This is simple API that allows to retrieve and update list of books in CSV format repo. All the main settings of proggram are conducted through config file located at config/config.json
* config.json file will allow to set custom port for communication with API, set Time format for JSON output, select CSV file to use
* If you want to test input capabilities of proggram and ability to update API through PUT endpoint, here is sample command:

```
curl -X PUT \
http://localhost:8000/books \
-H 'Content-Type: application/json' \
-d '[
{
"name": "The Fellowship of the Ring",
"author": "J. R. R. Tolkien",
"year": 1954
},
{
"name": "The Two Towers",
"author": "J. R. R. Tolkien",
"year": 1954
},
{
"name": "The Return of the King",
"author": "J. R. R. Tolkien",
"year": 1955
}
]'
```

### How to run Docker container

```
docker build -t fugas .

docker network create app-network

docker run -itd --name=test_fugas --network=app-network fugas
```
