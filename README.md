## Setup Golang
```
curl https://dl.google.com/go/go1.23.1.linux-amd64.tar.gz -o ~/go.tar.gz
```
```
tar -xvzf ~/go.tar.gz -C ~
```
```
echo "export PATH=$PATH:~/go/bin" >> ~/.bashrc
```
```
rm ~/go.tar.gz
```

# TP middleware example

## Run

Tidy / download modules :
```
go mod tidy
```
Build & run :
```
go run cmd/main.go
```

---
Or build : 
```
go build -o middleware_users cmd/main.go
```
Then run : 
```
./middleware_users
```

## Documentation

Documentation is visible in **api** directory ([here](api/swagger.json)).
