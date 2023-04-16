# go-restapi

## Build
```
docker-compose .devcontainer/ build
docker-compose .devcontainer/ up

go build -o ./bin/ ./gin-webservice/.
go build -o ./bin/ ./mux-webservice/.

./bin/gin-webservice&
./bin/mux-webservice&
```

### Test Gin Web Service
```
curl http://localhost:8080/albums

curl http://localhost:8080/albums/2

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Test Mux Web Service
```
curl http://localhost:9000/items

curl http://localhost:9000/items \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "new item"}'

curl http://localhost:9000/items/40b01ad4-bba4-4e18-8ac5-3fe201e9144c

curl http://localhost:9000/items/40b01ad4-bba4-4e18-8ac5-3fe201e9144c \
    --request "DELETE"

```

## Enable SSL (with MUX)
```
openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt
(the common name of the certificate should be localhost for local testing)

sudo cp localhost.crt /usr/local/share/ca-certificates/
sudo update-ca-certificates

./bin/mux-webservice&

curl http://localhost:9001/items
```

### Ref
```
mkdir gin-webservice
cd gin-webservice
go mod init go-restapi/gin-webservice
```

```
mkdir mux-webservice
cd mux-webservice
go mod init go-restapi/mux-webservice
```