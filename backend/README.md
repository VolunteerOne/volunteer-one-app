## Environment Variables
For the application to work correctly you will need a .env file at the root of your ```/backend``` project.
If it doesn't already exist create it and use the following as a template:

```
PORT=8000
MYSQL_ROOT_PASSWORD=root
DB_USER=testuser
DB_PASSWORD=Password@123
DB_HOST=db              
DB_PORT=3306            
DB_NAME=volunteerone
DB_MIGRATION=true
ENVIRONMENT=local           
```

**WARNING:**
DO NOT ALTER ANY VARIABLES FROM THIS LIST
- PORT
- DB_HOST
- DB_PORT
- DB_NAME

## Run Script
To use the automation script to start the application, complete the following from your terminal:
- Run```chmod 705 ./run.sh```
    - Add execute permissions to the local file
- Run ``` ./run.sh ```