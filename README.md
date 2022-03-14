![UserAPI-docker-Image](https://github.com/laithrafid/user-api/actions/workflows/main.yml/badge.svg?branch=main)

# USERAPI service written in go

## Development Pattern
![alt text](https://github.com/laithrafid/infra-api/blob/main/Images/devpattern.png?raw=true)


## Archietecture 
![alt text](https://github.com/laithrafid/infra-api/blob/main/Images/infra-api-architecture.png?raw=true)


## User APi Endpoints and Calls
GET  /users/:user_id
PUT /users/:user_id
PATCH /users/:user_id
DELETE /users/:user_id
GET internal/users/search
POST /users
POST /users/login


#variables
adding new variable
1. add your variable to app.env 
2. add your variable to struct in utils/config_utils/config_utils.go with map reference to file.env
3. load variable function 
```
    config, err := config_utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config of users_db:", err)
	}
```
4. make sure path works , tip add special log for when to load into package

#Programing tip for new features/functions creation , always start in this order:
1. domain
2. service (services Integration with other Rest service providers)
3. Database
4. controller
5. App Server(Gin,Urlmappings) 


#How to Run 

1. db Layer

```
CREATE SCHEMA `users_db` DEFAULT CHARACTER SET utf8 ;
CREATE TABLE `users_db`.`users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NOT NULL,
  `date_created` DATETIME NULL,
  `status` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);
```
