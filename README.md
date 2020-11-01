# nextjs-golang-mysql
next.js + Golang + MySQL on Docker

## environment
* node.js 14
* next.js 9.5
* tailwindcss
* golang 1.13
* mysql 5.7

## How to Use
first time
```bash
docker-compose build
docker-compose up -d
```
next time or restart
```bash
docker-compose up -d frontend
docker-compose restart frontend
```
finish time
```bash
docker-compose stop
```
Go to it  
frontend
```bash
http://localhost:8080
```
backend
```bash
http://localhost:8888
```
db
```bash
docker exec -it [CONTAINER ID] sh
```
