Tech Stack Used:
    - Golang, gin-gonic web framework, Postgres, React

Setup Stesp:
- Execute scripts/ui-build.sh
    - This will setup frontend UI and copy the build files in the appropriate location to be served by gin server.
- Database Creation
    - CREATE DATABASE restaurantdb;
- Activate extension hstore in postgresql
    - CREATE EXTENSIO hstore;
- Go to root directory and execute : go build main.go
    - This build binary executable file
- Execute ./main
    - This will start our server on Port 8070.
    - go to localhost:8070 to see UI
 

APIs:
To test this apis, use Postman.

- POST: localhost:8070/recipe/create
    Input:
    {
        "name": "Paneer",
        "preparation":"Make gravy and add Paneer",
        "noofingredients": 3,
        "ingredientsdetails": {"onion":"10", "tomato":"20", "paneer":"30"}
    }

    Output
    {
        "id": 1,
        "ingredientsdetails": {
            "onion": "10",
            "paneer": "30",
            "tomato": "20"
        },
        "name": "Paneer",
        "noofingredients": 3,
        "preparation": "Make gravy and add Paneer"
    }

- Update Recipe
    - All fields are mandatory
    - First click view to retreive and then update from it
    - Check for no of ingredients to be provided
    - 
