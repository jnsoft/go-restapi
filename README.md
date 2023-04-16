# go-restapi

## MS SQL
```
# Import the public repository GPG keys
curl https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -

# Register the Microsoft Ubuntu repository
sudo apt-add-repository https://packages.microsoft.com/ubuntu/20.04/prod

# Update the list of products
sudo apt-get update

# Install mssql-cli
sudo apt-get install mssql-cli

# Install missing dependencies
sudo apt-get install -f

docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=123abcAbc" -e "MSSQL_PID=Express" -p 1433:1433 -d mcr.microsoft.com/mssql/server:2019-latest

pip install mssql-cli
pip install cli-helpers --upgrade --force
mssql-cli -S localhost -U sa -P 123abcAbc
```

## Init
```
sudo apt-get update
sudo apt-get install -y postgresql-client
docker-compose .devcontainer/ build (docker-compose -f .devcontainer/docker-compose.yml build)
docker-compose .devcontainer/ up (docker-compose -f .devcontainer/docker-compose.yml up)



psql -h localhost -p 5432 -d albums -U postgres -W  
postgres>select * from albums \x\g\x

postgres>\l (list databases)
postgres>\c db_name (change databse)
postgres>\dt (list tables)
postgres>\dt+ (list tables)
```
## Build
```
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

curl https://localhost:9001/items
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

 cat /etc/os-release
 	
pip install mssql-cli
```