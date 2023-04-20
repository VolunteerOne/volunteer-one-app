# User Creation
Sign Up (POST):
	Pass the new user’s Email, FirstName, LastName, and Password through a JSON payload or a JSON body. 
    The call will then return either one or two json key values depending on the success or failure of the function. 
    The two json key values are “error” and “error message”.
http://www.localhost:8000/user

Login (GET):
	Pass the user’s email and password through the call. The call will then return two json key values 
    labeled “message” and “success”. The message will explain either what went wrong or right. 
    Success will return “true” or “false” depending on if the user is able to be found in the database or not. 
	Example call: http://www.localhost:8000/login/useremail@gmail.com/userpassword

Send Reset Code to Email (POST):
	Pass the user’s email through the call. The call will then return two json key values 
    labeled “message” and “success”. The message will explain either what went wrong or right. 
    Success will return “true” or “false” depending on if the function was successfully able to complete its job or not.
	Example call: http://www.localhost:8000/login/useremail@gmail.com

Reset User’s Password(PUT):
	Pass the user’s email, new password, and reset code (get from call above). The call will then return two 
    json key values labeled “message” and “success”. The message will explain either what went wrong or right. 
    Success will return “true” or “false” depending on if the password was changed or not.
	Example call: http://www.localhost:8000/login/useremail@gmail.com/resetcode/newpass

# Organizations

## Create Organization (POST)

Endpoint: `/organization`

Example Request Body
```
{
    "name": string,
    "description": string,
    "verified": bool,
    "interests" string,
}
```

Success: Status Code 200, JSON object

Fail: Status Code 400, JSON error message

## Get All Organizations (GET)

Endpoint: `/organization`

Success: Status Code 200, Objects In JSON

Fail: Status Code 400, JSON error message

## Get One Organization By ID (GET)

Endpoint: `/organization/:id`

Success: Status Code 200, JSON object

Fail: Status Code 400, JSON error message

## Update An Organization's Details (PUT)

Endpoint: `/organization/:id`

Example Request Body
```
{
    "name": string,
    "description": string,
    "verified": bool,
    "interests" string,
}
```

Success: Status Code 200, JSON object

Fail: Status Code 400, JSON error message

## Delete An Organization (DELETE)

Endpoint: `/organization/:id`

Success: Status Code 200, JSON success message

Fail: Status Code 400, JSON error message

# Organization User Roles

## Create An Organization User (POST)

> Note: The user and organization must exist first.

Endpoint: `/orgUsers`

Example Request Body
```
{
    "userID": uint,
    "orgID": uint,
    "role": uint,
}
```
Success: Status Code 200, JSON object

Fail: Status Code 400, JSON error message

## Get All Organization Users (GET)

Endpoint: `/orgUsers`

Success: Status Code 200, Objects In JSON

Fail: Status Code 400, JSON error message

## Get One Organization User By ID (GET)

Endpoint: `/orgUsers/:id`

Success: Status Code 202, JSON object

Fail: Status Code 400, JSON error message

## Update An Organization Users Details (PUT)

> Useful for updating a role

Endpoint: `/orgUsers/:id`

Example Request Body
```
{
    "userID": uint,
    "orgID": uint,
    "role": uint,
}
```

Success: Status Code 200, JSON object

Fail: Status Code 400, JSON error message

## Delete An Organization User (DELETE)

Endpoint: `/orgUsers/:id`

Success: Status Code 200, JSON success message

Fail: Status Code 400, JSON error message
