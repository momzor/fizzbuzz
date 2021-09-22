# LBC test fizzbuzz

## The fizzbuzz server
This application is a web serverr exposing a [fizzbuzz entry point](http://127.0.0.1:8000/fizzbuzz?int1=3&int2=5&limit=30&str1=fizz&str2=buzz).
You can also check a [stats entry point](http://127.0.0.1:8000/stats) in order to get the most used request and its parameters


For mor informations, please install the project et check  [swagger](http://localhost:8000/swagger/fizzbuzz/index.html#)


## Requirements
- Docker-Compose
- Docker
- [GNU Make](https://www.gnu.org/software/make/)
- Golang 1.17 


## Install

 `make deps` Will install utils dependencies on your host  (swagger generation & unit test mock generation)

 `make up` Will use docker-compose to build and run your app 


## Use
This application provide an OpenApi documentation you can check here : [swagger](http://localhost:8000/swagger/fizzbuzz/index.html#), once you have lauched the app


## tests

Run `make unit-test` for unit testing the application

## V2. enhancements & comming features :
- [] Hot reload for dev env [with air](https://github.com/cosmtrek/air)
- [] functional tests
- [] inject dedecated logger
- [] add pre commit hooks for lint
- [] add a fixtures file for fizzbuzz builder test for more table cases
