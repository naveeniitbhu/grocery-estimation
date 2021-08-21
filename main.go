package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/lib/pq/hstore"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pspasswd"
	dbname   = "restaurantdb"
)

type App struct {
	R  *gin.Engine
	Db *sql.DB
}

type Id struct {
	ID int64 `json:"id,omitempty" db:"id"`
}

type Dish struct {
	Id                 int64             `json:"id" db:"id"`
	Name               string            `json:"name" db:"name"`
	Preparation        string            `json:"preparation" db:"preparation"`
	Noofingredients    int64             `json:"noofingredients" db:"noofingredients"`
	Ingredientsdetails map[string]string `json:"ingredientsdetails" db:"ingredientsdetails"`
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("database error: %v", err)
		panic(err)
	} else {
		log.Println("connected to database")
	}

	defer db.Close()

	createtablestmt := `CREATE TABLE IF NOT EXISTS dishes (
		id serial primary key,
		name VARCHAR(255),
		preparation VARCHAR(255),
		noofingredients Int,
		ingredientsdetails hstore
	)`
	_, err = db.Exec(createtablestmt)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	} else {
		log.Println("INFO: Dishes table is created.")
	}

	app := App{
		R:  gin.Default(),
		Db: db,
	}

	app.R.Static("/static", "./static/static")
	app.R.LoadHTMLGlob("./static/index.html")
	app.R.GET("/", displayHtml)

	app.R.GET("/recipe/view/:dishname", app.ViewRecipe)
	app.R.POST("/recipe/create/", app.CreateRecipe)
	app.R.PUT("/recipe/create/", app.UpdateRecipe)
	app.R.DELETE("/recipe/delete/:dishname", app.DeleteRecipe)

	app.R.Run(":8070")
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func displayHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (app *App) CreateRecipe(c *gin.Context) {

	var (
		dish Dish
		id   int64
	)

	db := app.Db

	if err := c.ShouldBindJSON(&dish); err == nil {
		dish.Name = strings.Title(dish.Name)
		log.Println("INFO: json binding is successful")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"json binding error": err.Error()})
		panic(err)
	}

	dataIngred := hstore.Hstore{}
	dataIngred.Map = make(map[string]sql.NullString)

	for k, v := range dish.Ingredientsdetails {
		dataIngred.Map[k] = ToNullString(v)
	}
	err := db.QueryRow(`INSERT INTO dishes(
		name, 
		preparation, 
		noofingredients, 
		ingredientsdetails) 
		VALUES($1, $2, $3, $4) RETURNING id`,
		dish.Name, dish.Preparation, dish.Noofingredients, dataIngred).Scan(&id)

	if err == nil {
		fmt.Printf("INFO: Dish details inserted with id:%d & name:%s\n", id, dish.Name)
		c.JSON(http.StatusOK, gin.H{
			"id":                 id,
			"name":               dish.Name,
			"preparation":        dish.Preparation,
			"noofingredients":    dish.Noofingredients,
			"ingredientsdetails": dish.Ingredientsdetails,
		})
	} else {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Dish details not inserted properly"})
		panic(err)
	}
}

func (app *App) ViewRecipe(c *gin.Context) {

	// dishname capitalize below
	dishname := strings.Title(c.Query("dishname"))

	var data []uint8
	// hstore data from db is stored in data

	dish := Dish{Ingredientsdetails: make(map[string]string)}
	// dish initialized using to avoid nil map assignment Error

	db := app.Db
	dataIngred := hstore.Hstore{}

	err := db.QueryRow(`Select id,name,preparation,noofingredients,ingredientsdetails From dishes WHERE name=$1`, dishname).Scan(&dish.Id, &dish.Name, &dish.Preparation, &dish.Noofingredients, &data)

	if err != nil {
		log.Println("Error: error during getting data from database")
		panic(err.Error())
	}

	err = dataIngred.Scan(data)
	if err != nil {
		log.Println("Error: error during scanning data in hstore")
		panic(err.Error())
	}

	for k, v := range dataIngred.Map {
		dish.Ingredientsdetails[k] = v.String
	}

	if err != nil {
		log.Println(err.Error())
		panic(err)
	} else {
		log.Println("INFO: successful-get request for dish")
		c.JSON(200, gin.H{
			"id":                 dish.Id,
			"name":               dish.Name,
			"preparation":        dish.Preparation,
			"noofingredients":    dish.Noofingredients,
			"ingredientsdetails": dish.Ingredientsdetails,
		})
	}

}

func (app *App) UpdateRecipe(c *gin.Context) {
	var (
		dish Dish
		id   int64
	)

	db := app.Db

	if err := c.ShouldBindJSON(&dish); err == nil {
		dish.Name = strings.Title(dish.Name)
		log.Println("INFO: json binding is successful")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"json binding error": err.Error()})
		panic(err)
	}

	dataIngred := hstore.Hstore{}
	dataIngred.Map = make(map[string]sql.NullString)

	for k, v := range dish.Ingredientsdetails {
		dataIngred.Map[k] = ToNullString(v)
	}

	err := db.QueryRow(`Update dishes SET 
						name=$1,
						preparation=$2, 
						noofingredients=$3, 
						ingredientsdetails=$4
						WHERE name=$1
						AND (length($1)>0) 
						AND (length($2)>0) 
						AND ($3>0) 
						AND (array_length(akeys($4::hstore),1) >0) 
						RETURNING id`,
		&dish.Name, &dish.Preparation, &dish.Noofingredients, &dataIngred).Scan(&id)

	if err == nil {
		fmt.Printf("INFO: name:%s dish details updated with\n", dish.Name)
		c.JSON(http.StatusOK, gin.H{
			"id":                 id,
			"name":               dish.Name,
			"preparation":        dish.Preparation,
			"noofingredients":    dish.Noofingredients,
			"ingredientsdetails": dish.Ingredientsdetails,
		})
	} else {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Dish details not updated properly"})
		panic(err)
	}
}
func (app *App) DeleteRecipe(c *gin.Context) {

	// dishname capitalize below
	dishname := strings.Title(c.Query("dishname"))
	db := app.Db

	_, err := db.Exec(`DELETE FROM dishes WHERE name=$1`, dishname)

	if err != nil {
		log.Println("Error: error during deleting data from database")
		panic(err)
	} else {
		log.Printf("INFO: successful deleted dish %s", dishname)
		c.JSON(200, gin.H{"name": dishname})
	}

}
