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
<<<<<<< HEAD
- Run ``` ./run.sh ```

## Project Layout
In order to support testability, the project now supports the 
Controller -> Service -> Repository schema.

`controllers`
This is where all the public facing routes are implemented. Controllers should
NOT have any direct interactions with the database. In order to interact with
the database, the controllers will interact with the service layer.

`service`
This is the layer that handles "business logic". Tasks used by controllers
should be implemented in the services layer. However, the services layer will
still NOT touch the database directly.

`repository`
The repository layer is the lowest level and interacts with the database. 
Functions should be written to interact with the database here.

### Why?
With three different layers, we can appropriately mock the underlying layer.
For example, in the `controllers` layer, we will mock the `service` layer. This
allows separation of the layers so that not all of the layers are strictly
dependent on each other. 

Controllers -> mock service
Service -> mock repository
Repository -> mock sql

### How To Implement
In order to test properly, we will have to use Go's design pattern of Interfaces.

#### Controllers

Check out `controllers/loginController.go`

You first define an interface, in this case `LoginController`. This interface
will have the function signatures of all the methods that will need to be
implemented. Any functions that you are planning to implement for that
controller should have a function signature in there.

Next, we have a struct, in this case `loginController`. This struct holds
one value -> the variable of the layer beneath, in this case `service`.

Next, we have to initialize the `LoginController` by taking in the corresponding
service and returning a `LoginController`. This function should initialize
the variable that holds the reference to the lower layer.

The `NewLoginController` function is called in server/router.go. This is
also where you place all of your created routes.

You then can implement your normal handlers, but making sure that if you
have to interact with the DB, you go through a `service` method. Other helper
methods can also be placed in `service` to try to make code easier to test.

#### Service

Check out `service/loginService.go`

We will take the same exact approach as controllers. All methods to implemented
should be placed in an interface.

There should be a struct that holds a variable to the lower layer -> the repo layer

There should be a method to create a new service object. This method should
be called in `server/router.go`

If a database connection is required, you will use the repo layer to get it.

#### Repo

Check out `repository/loginRepository.go`

The same interface/struct setup will be used, however since there is no lower
layer we will be using a reference to the database directly. For the methods
implemented here, make the actual calls to the database using GORM.

Again an initialization method should be placed in `server/router.go`

#### Mocks

The only way to Mock in Go is through interfaces. Since we implemented the interfaces,
we will just have to reimplement those interfaces for mocks. Luckily, Go has
libraries for automatically generating mock code based on interfaces. To generate
the mocks based on your implemented interfaces, you will need to install 
[`mockery`](https://github.com/vektra/mockery). It is recommended to install directly
on your machine using a package manager (like brew). Once installed, you can
run `mockery --all` to generate the mock interfaces for all interfaces in the project,
or you can specify exactly which files to generate from (which is safer). All mocks
are placed in the `mocks` directory and can be directly used in the tests.

#### Writing Tests

Remember, we are focusing on Unit test coverage. Unit tests should not require
external connections such as an http server or database to run. You can definitely
write tests that do make actual http calls and database connections, but these
are no longer Unit tests and are Integration tests.

To write tests, make a file in the package that ends in `_test.go`, for example
`loginController_test.go`. Look at the example Login file for each layer. 
For controllers, we will need to mock the service layer and http requests. For
the service layer, we will need to mock the repo layer. For the repo layer,
we will need to mock the actual sql queries. 

It should be possible to get 100% coverage of the methods with this design pattern,
so please ask if you are stuck or unsure with how the mocking works. The mock interfaces
generated by mockery should be used whenever possible.

To run the tests: `go test ./... -cover -v`
=======
- Run ``` ./run.sh ```
>>>>>>> 4e9ea3a (Update (#5))
