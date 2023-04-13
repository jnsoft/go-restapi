# go-restapi

```
go build -o ./bin/ ./gin-webservice/.
go build -o ./bin/ ./mux-webservice/.
```

```
openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt
(the common name of the certificate should be localhost for local testing)

sudo cp localhost.crt /usr/local/share/ca-certificates/
sudo update-ca-certificates
```

## Gin Example

### Start
```
mkdir gin-webservice
cd gin-webservice
go mod init go-restapi/gin-webservice
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

## Mux Example

### Start
```
mkdir mux-webservice
cd mux-webservice
go mod init go-restapi/mux-webservice
```

### Test Mux Web Service
```
curl http://localhost:8080/items

curl http://localhost:8080/items \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "new item"}'

curl http://localhost:8080/items/40b01ad4-bba4-4e18-8ac5-3fe201e9144c

curl http://localhost:8080/items/40b01ad4-bba4-4e18-8ac5-3fe201e9144c \
    --request "DELETE"

```
