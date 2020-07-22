# Online-Shopping-microservices



Golang APIs for online shopping microservices.

this project contains backend APIs and logic for online shopping microservices powered by Go,Gin,Redis and MongoDB
> Application runs on Go 1.14 And redis and MongoDB.
  
## Table of Contents  
  
- [Project Structure](#Project Structure)  
- [Install application](#Install application)  
- [Configuration](#configuration)  
- [Features](#Features)  
- [Services and Database dependencies](#Services and Database dependencies)  

  


## Project Structure 
```
.
├── config                      # all config files is here
    ├── file                    # develop yml config file is here
    └── test_conf               # unit test yml config file is here
├── console                     # console commands dicretory
    └── cmd                     # all console commands is here 
├── constants                   # some static values is here
├── controller                  # controller layer directories. all server controllers logics is here 
    └── http                    # a directory for http controllers
├── docs                        # any documentation is here swagger,flowchat,gif , ....
├── domain                      # we save project struct/models in here
├── logic                       # business logic layer directories. all server business logics is here
├── repository                  # repository layer directories. all server repository logics is here
    ├── rediskeys               # redis keys names is here
├── services                    # request and connections to therd party servers is here
├── go.mod                      # go modules file
├── go.sum                      # go modules file
├── main.go                     # project started in this file

```

## Install application
###  Get application to run locally for development or deployment

```bash
cd micro-service
```
### installation dependencies
To installation dependencies use ```go mod vendor``` command.

## Configuration  
by default app server in running read config file from ./file/cofigs.yml if you want to change some service or database address or ports must edit config file
### Setup your own config
to setup config you can set your settings into yml file.

### Run Application
To run application on development config use
```go run main.go``` 

## Features
### consoles
some time we need some worker to do some jobs periodically we user `cobra` to write console commands to do some jobs.

#### worker
to run worker to get real time users go to `console` and run 
`go run main GetUsers`
### API's

#### http
we have multiple http route :

- `v1/discounts/discounts` `POST`: api for create discount ( enter the new discount code)
- `v1/wallet/info/{phone_number}` `GET`: api for get a wallet information
- `v1/wallet/win/user` `GET`: api for get a list of real time winning users



## Services and Database dependencies
his server need `redis`and`MongoDB` servers to start running.



