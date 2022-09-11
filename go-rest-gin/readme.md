## Developing a RESTful API with Go and Gin

$ curl http://localhost:8080/albums \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

curl http://localhost:8080/albums \
--header "Content-Type: application/json" \
--request "GET"

# Reference
pkg.go.dev
https://go.dev/src/

https://go.dev/doc/tutorial/getting-started
https://go.dev/doc/tutorial/web-service-gin
https://go.dev/doc/effective_go