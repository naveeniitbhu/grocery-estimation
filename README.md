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

- POST: CREATE RECIPE
    localhost:8070/recipe/create
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
        
- GET: VIEW RECIPE(Using querystring)
    localhost:8070/recipe/view?dishname=paneer
        Output:
            {
                "id": 1,
                "ingredientsdetails": {
                    "onion": "10",
                    "paneer": "30",
                    "tomato": "20"
                },
                "name": "Paneer",
                "noofingredients": 5,
                "preparation": "Make gravy and add paneer"
            }

- PUT: UPDATE RECIPE (First Click view to retrieve details and then update it)
    localhost:8070/recipe/update
        Input:
            {
                "name": "juice",
                "preparation":"cut fd dsdadsadand mic",
                "noofingredients": 4,
                "ingredientsdetails": {"a":"103330","b":"200","c":"300","d":"500"}
            }

        Output:
            {
                "id": 12,
                "ingredientsdetails": {
                    "a": "103330",
                    "b": "200",
                    "c": "300",
                    "d": "500"
                },
                "name": "Juice",
                "noofingredients": 4,
                "preparation": "cut fd dsdadsadand mic"
            }

- DELETE: DELETE RECIPE by dishname
    localhost:8070/recipe/delete?dishname=paneer
