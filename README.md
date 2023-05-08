
# SAM Monitoring BOT

Service and Website for Monitoring BOT

  
## API

  **Endpoint** : [localhost:3000/](localhost:3000/)
	
    
---	
### Validate User
* **Method**
	`GET`
* **URL**
`/monitoringbot/login`
	* **Parameter**
		`username`
		`password`
		
---
### Get Bots Activity
* **Method**
	`GET`
* **URL**
`/monitoringbot/getbots`
	* **Parameter**
		`username`
		`password`
		`monitor`
		
* **Response**
	```json 
	{
	 	"list"	: [ 
					{ 
						"lastupdate": "int",
						"monitor"	: "string",
						"name"		: "string",
						"status"	: "string",
						"world"		: "string", 
						"level"		: "string", 
						"captcha" 	: "string", 
						"x" 		: "int", 
						"y" 		: "int", 
						"profit" 	: [
										{
											"id" : "int", 
											"total": "int"
										},
									] 
					},
				]
	} 
	```
---
### Get Monitors
* **Method**
	`GET`
* **URL**
`/monitoringbot/findmonitors`
	* **Parameter**		
		`username`
		`password`

* **Response**
	```json 
	["monitor1", "monitor2"]
	```
	
---
### Get Bots by Status
* **Method**
	`GET`
* **URL**
`/monitoringbot/findbotsbystatus`
	* **Parameter**
		`username`
		`password`
		`monitor`
		`status`
		
* **Response**
	```json 
	{
		"list" 	: [ 
					{ 
						"lastupdate": "int",
						"monitor"	: "string",
						"name"		: "string",
						"status"	: "string",
						"world"		: "string", 
						"level"		: "string", 
						"captcha" 	: "string", 
						"x" 		: "int", 
						"y" 		: "int", 
						"profit" 	: [
										{
											"id" : "int", 
											"total": "int"
										},
									] 
					},
				]
	} 
	```

---
### Remove Monitor
* **Method**
	`POST`
* **URL**
`/monitoringbot/removemonitor`
	* **Parameter**
		`username`
		`password`
		`monitor`

---
### Insert Bot
* **Method**
	`POST`
* **URL**
`/monitoringbot/insertbot`
	* **Parameter**
		`username`
* **Body**
	```json
	{ 
		"password"	: "string" ,
		"monitor" 	: "string" ,
		"list" 		: [ 
						{ 
							"name"		: "string" ,
							"status"	: "string" ,
							"world"		: "string" ,
							"level"		: "string" ,
							"captcha" 	: "string" ,
							"x" 		: "int" ,
							"y" 		: "int" ,
							"profit" 	: [
											{
												"id" : "int", 
												"total": "int"
											},
										  ] 
						} ,
					 ] 
	}
	```

---
### Insert User
* **Method**
	`POST`
* **URL**
	`/monitoringbot/insertuser`
	* **Parameter**
		`username`
		`password`
* **Response**
	```json
	{
		"username" : "string",
		"password" : "string",
		"monitors" : []
	}
	```