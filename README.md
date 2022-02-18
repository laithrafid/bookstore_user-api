#BookStore User 
### it follows MVC patter 


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
