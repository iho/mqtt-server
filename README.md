# mqtt-server

## How to run 
```
sudo snap install mosquitto
go get -v github.com/rubenv/sql-migrate/...
sql-migrate up
go run cmd/server/main.go
```
To check if everything works right send test message with this command: 
```
mosquitto_pub -t test/topic -m "{\"message\": \"test\", \"name\": \"mike\"}"
``` 
## How to run with Docker 
```
docker-compose up
```
To check if everything works right send test message with this command: 
```
mosquitto_pub -t test/topic -m "{\"message\": \"test\", \"name\": \"mike\"}" -p 1885
```
