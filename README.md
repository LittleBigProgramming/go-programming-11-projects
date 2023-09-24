# go-programming-11-projects

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
