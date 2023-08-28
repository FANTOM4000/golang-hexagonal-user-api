# golang-hexagonal-user-api
Template or Example for golang on hexagonal architecture with mongodb

## Run API

### On vscode 
#### 1. [install golang](https://go.dev/dl)
#### 2. Set go module `go env -w GO111MODULE="on"`
#### 3. Ctrl+Shift+D AND F5

###  On docker
#### 1. First time build and run
```
docker build -t golang-hexagonal-user-api .
docker run -d -p 8080:8080 --env-file ./.env.example --name golang-hexagonal-user-api golang-hexagonal-user-api
```
#### 2. delete container if update env 
```
docker stop golang-hexagonal-user-api
docker rm golang-hexagonal-user-api
```
#### 3. delete image if update api
```
docker image rm golang-hexagonal-user-api
```
#### 4. [goto 1. to build and run](#1-first-time-build-and-run)
### On command
#### 1. [install golang](https://go.dev/dl)
#### 2. Set go module `go env -w GO111MODULE="on"`
#### 3. run below command
``` 
go run cmd/main.go 
```
## API Document
import postman collection from postman.json
set global variable
| variable        | value       |
| ------------- |:-------------:|
| base_url      | http://localhost:8080 |
| token      | xxxx      |

## Additional
Create local redis
```
docker pull redis
docker run -p 6379:6379 -d redis
docker run --name redis -d -p 6379:6379 redis redis-server --requirepass "SUPER_SECRET_PASSWORD"
```
Create local mongodb
mongodb://docker:mongopw@localhost:2701
```
docker pull mongo
docker run -p 2701:2701 -d -e MONGO_INITDB_ROOT_USERNAME=docker -e MONGO_INITDB_ROOT_PASSWORD=mongopw mongo
```
