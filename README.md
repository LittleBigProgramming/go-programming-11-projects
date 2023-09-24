# go-programming-11-projects

## Simple Web Server
Simple web server that handles static files and a file submission from `form.html` to the form handler method

## Movies CRUD API

Getting all movies
`curl -X GET http:/localhost:8000/movies -H 'Content-Type: application/json'`

```
 [{"id":"1","isbn":"123456","title":"Movie One","director":{"firstName":"Jane","lastName":"Doe"}},

{"id":"2","isbn":"654321","title":"Movie Two","director":{"firstName":"John","lastName":"Smith"}}]
```

Getting Movie by ID
`curl -X GET http:/localhost:8000/movies/1 -H 'Content-Type: application/json'`

Creating a Movie
```
curl -X POST http:/localhost:8000/movies -H 'Content-Type: application/json' -d 
'{"isbn":"123456","title":"Movie One", "firstName":"Jane","lastName":"Doe"}'
```

Updating a Movie
```
curl -X PUT http:/localhost:8000/movies/1 -H 'Content-Type: application/json' -d 
'{"isbn":"123456","title":"Updated Movie One", "firstName":"Jane","lastName":"Doe"}'
```

Deleting a Movie
```
curl -X DELETE http:/localhost:8000/movies/1 -H 'Content-Type: application/json'
```

## Bookstore API using mysql
Simple bookstore api in mysql, not using a framework, using:
- https://github.com/go-gorm/gorm - ORM library
- https://github.com/gorilla/mux - http router

https://youtu.be/jFfo23yIWac?feature=shared&t=4034 created by following

Only difference is the current configuration uses sqlite to avoid having to
setup a database for an application with simple functionality/foreign keys.
The method of setting up mysql using dsn and config is commented out incase expanded in the future.

More information can be found of this here - https://gorm.io/docs/connecting_to_the_database.html

### Creating a Book
Request:
```
curl --location 'http://localhost/book' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data '{
    "Name": "To Kill A Mockingbird",
    "Author": "Harper Lee",
    "Publication": "Penguin Books"
}'
```

Response
```
{
    "ID": 2,
    "CreatedAt": "2023-09-24T12:09:40.612741479Z",
    "UpdatedAt": "2023-09-24T12:09:40.612741479Z",
    "DeletedAt": null,
    "name": "To Kill A Mockingbird",
    "author": "Harper Lee",
    "publication": "Penguin Books"
}
```

### Get Books
Request:
```
curl --location 'http://localhost/book' \
--header 'Accept: application/json'
```

Response:
```
[
    {
        "ID": 1,
        "CreatedAt": "2023-09-24T12:07:35.182427689Z",
        "UpdatedAt": "2023-09-24T12:07:35.182427689Z",
        "DeletedAt": null,
        "name": "Pride and Prejudice",
        "author": "Jane Austin",
        "publication": "Penguin Books"
    },
    {
        "ID": 2,
        "CreatedAt": "2023-09-24T12:09:40.612741479Z",
        "UpdatedAt": "2023-09-24T12:09:40.612741479Z",
        "DeletedAt": null,
        "name": "To Kill A Mockingbird",
        "author": "Harper Lee",
        "publication": "Penguin Books"
    }
]
```

### Get Book by ID
Request:
```
curl --location 'http://localhost//book/1' \
--header 'Accept: application/json'
```

Response:
```
{
    "ID": 1,
    "CreatedAt": "2023-09-24T12:07:35.182427689Z",
    "UpdatedAt": "2023-09-24T12:07:35.182427689Z",
    "DeletedAt": null,
    "name": "Pride and Prejudice",
    "author": "Jane Austin",
    "publication": "Penguin Books"
}
```

### Update a Book
Request:
```
curl --location --request PUT 'http://localhost/book/1' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data '{
    "Publication": "Another Book Publisher"
}'
```

Response:
```
{
    "ID": 1,
    "CreatedAt": "2023-09-24T12:07:35.182427689Z",
    "UpdatedAt": "2023-09-24T12:07:37.182427689Z",
    "DeletedAt": null,
    "name": "Pride and Prejudice",
    "author": "Jane Austin",
    "publication": "Another Book Publisher"
}
```

### Delete a Book
Request:
```
curl --location --request DELETE 'http://localhost/book/1' \
--header 'Accept: application/json' \
--data ''
```

Response:
```
Empty 200 OK
```
