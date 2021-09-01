- create extension hstore;
- create database restaurantdb;

- Post - Create/recipe done
    Input:
    {
        "name": "Paneer",
        "preparation":"dfndsfndfjndlkfdfdsfndsfndsfdnkl",
        "noofingredients": 5,
        "ingredientsdetails": {"a":"10", "b":"20", "c":"30"}
    }
    Output
    {
    "id": 26,
    "ingredientsdetails": {
        "a": "10",
        "b": "20",
        "c": "30"
    },
    "name": "Paneer",
    "noofingredients": 5,
    "preparation": "dfndsfndfjndlkfdfdsfndsfndsfdnkl"
    }

- Update Recipe
    - All fields are mandatory
    - First click view to retreive and then update from it
    - Check for no of ingredients to be provided
    - 

- Execute scripts/ui-build.sh to build and serve ui files


// React
    - Creaating create recipe ui
    - 