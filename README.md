
# SAM Monitoring BOT

Service and Website for Monitoring BOT

  
## API

  **Endpoint** : [localhost:3000/](localhost:3000/)
	
    
---	
### Validate User
* **URL**
`/monitoringbot/login`
	* **Parameter**
		`username`
		`password`
		
---
### Get Bots Activity
* **URL**
`/monitoringbot/getbots`
	* **Parameter**
		`username`
		`password`
		`monitor`
		
* **Response**
	```json 
	{
		"list" 		: [ 
						{ 
							"lastupdate": int,
							"monitor"	: string,
							"name"		: string,
							"status"	: string,
							"world"		: string, 
							"level"		: string, 
							"captcha" 	: string, 
							"x" 		: int, 
							"y" 		: int, 
							"profit" 	: [
											{
												id : int, 
												total: int
											},
										  ] 
						},
					 ]
	} 
	```
---
### Get Monitors
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
		"list" 		: [ 
						{ 
							"lastupdate": int,
							"monitor"	: string,
							"name"		: string,
							"status"	: string,
							"world"		: string, 
							"level"		: string, 
							"captcha" 	: string, 
							"x" 		: int, 
							"y" 		: int, 
							"profit" 	: [
											{
												id : int, 
												total: int
											},
										  ] 
						},
					 ]
	} 
	```

---
### Remove Monitor
* **URL**
`/monitoringbot/removemonitor`
	* **Parameter**
		`username`
		`password`
		`monitor`

---
### Insert Bot
* **URL**
`/monitoringbot/insertbot`
	* **Parameter**
		`username`
* **Body**
	```json
	{ 
		"password"	: string 
		"monitor" 	: string 
		"list" 		: [ 
						{ 
							"name"		: string 
							"status"	: string 
							"world"		: string 
							"level"		: string 
							"captcha" 	: string 
							"x" 		: int 
							"y" 		: int 
							"profit" 	: [
											{
												id : int, 
												total: int
											}
										  ] 
						} 
					 ] 
	}
	```

---
### Insert User
* **URL**
	`/monitoringbot/insertuser`
	* **Parameter**
		`username`
		`password`
* **Response**
	```json
	{
		"username" : string,
		"password" : string
		"monitors" : []
	}
	```