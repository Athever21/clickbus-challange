# Clickbus - Challange

###### My attempt to solve https://github.com/RocketBus/quero-ser-clickbus/tree/master/testes/backend-developer challange
#
#
----
###### link: https://athever-clickbus-challange.herokuapp.com/api/places
## Add place
```
curl -XPOST -H "Content-type: application/json" -d '{
    "name": "From curl",
    "slug": "test",
    "city": "test",
    "state": "test"
}' 'https://athever-clickbus-challange.herokuapp.com/api/places'
```
## Update Place
```
curl -X PUT -H "Content-type: application/json" -d '{
    "name": "Changed cURL",
    "slug": "test",
    "city": "test",
    "state": "test"
}' 'https://athever-clickbus-challange.herokuapp.com/api/places/:id'
```
## Delete Place
```
curl -X DELETE 'https://athever-clickbus-challange.herokuapp.com/api/places/60f3ffc6306e2e1bdb52b5c4'
```