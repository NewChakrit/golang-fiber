package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func main () {

	var err error
	db,err = sqlx.Open("mysql","root:P@ssw0rd@tcp(13.73.163.73:3306)/techcoach")
	if err!=nil{
		panic(err)
	}

	app:= fiber.New()

	app.Post("/signup",Signup)
	app.Post("/login",Login)
	app.Get("/heloo",Hello)
	
	app.Listen(":8000")
}

func Signup (c *fiber.Ctx)error {

	request := SignupRequest{}
	err:=c.BodyParser(&request)
	if err!=nil{
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	query:= "insert user (username, password) values (?, ?)"
	result,err:= db.Exec(query, request.Username, request.Password)
	if err!=nil{
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err:= result.LastInsertId()
	if err!=nil{
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := User{
		Id:int(id),
		Username: request.Username,
		Password: request.Password,
	}

	return c.JSON(user)
}

func Login (c *fiber.Ctx)error {
	return nil
}

func Hello (c *fiber.Ctx)error {
	return nil
}

type User struct{
	Id int `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type SignupRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func Fiber () {

	app:=fiber.New(fiber.Config{
		Prefork: true,
		// CaseSensitive: false,
		// StrictRouting: false,
	})

	//Middleware
	app.Use("/hello",func (c *fiber.Ctx) error  {
		c.Locals("name", "New")
		fmt.Println("before")
		err:=c.Next()
		fmt.Println("after")
		return err
	})

	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	// app.Use(logger.New(logger.Config{
	// 	TimeZone: "Asia/Bangkok",
	// }))

	//GET
	app.Get("/hello", func(c *fiber.Ctx) error {
		name:= c.Locals("name")
		fmt.Println("Hello")
		return c.SendString(fmt.Sprintf("GET: Hello world %v", name))
	})

	//POST
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("POST: Hello world")
	})

	//Parameters
	// app.Get("/hello/:name", func(c *fiber.Ctx) error {
	// 	name := c.Params("name")
	// 	return c.SendString("name: " + name)
	// })

	//Parameters Optional
	app.Get("/hello/:name/:surname", func(c *fiber.Ctx) error {      // surname? มีหรือไม่มีก็ได้
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("name: " + name + "surname: " + surname)
	})

	//ParamsInt
	app.Get("/hello/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil{
			return fiber.ErrBadRequest
		}

		return c.SendString(fmt.Sprintf("ID = %v", id))
	})

	//Query
	app.Get("query", func(c *fiber.Ctx) error {  //เวลาส่ง query string ต้องใส่? => curl "localhost:8000/query?name=<input>"
		name := c.Query("name")  
		surname:=c.Query("surname")   // curl "localhost:8000/query?name=New&surname=Chakrit"
		return c.SendString("name: "+ name +", surname: " + surname)
	})

	//Query2
	app.Get("query2", func(c *fiber.Ctx) error {  // curl "localhost:8000/query2?id=1&name=New"        
		person := Person{}  
		c.QueryParser(&person)
		return c.JSON(person)
	})


	//Wildcards
	app.Get("/wildcards/*", func(c *fiber.Ctx) error { //gofiber % curl localhost:8000/wildcards/hello/world
		wildcard := c.Params("*")
		return c.SendString(wildcard)   // result => hello/world
	})

	//Static file 
	app.Static("/", "./wwwroot", fiber.Static{
		Index: "index.html",
		CacheDuration: time.Second * 10,
	})

	//New Error
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	//Group
	v1:= app.Group("/v1",func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v1")
	})

	v2:= app.Group("/v2",func(c *fiber.Ctx) error {
		c.Set("Version", "v2")
		return c.Next()
	})

	v2.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v2")
	})

	//Mount
	userApp:=fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})

	app.Mount("/user",userApp)

	//Server
	app.Server().MaxConnsPerIP = 1
	app.Get("/server",func(c *fiber.Ctx) error {
		time.Sleep(time.Second*30)
		return c.SendString("server")
	})

	app.Get("/env",func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseURL": c.BaseURL(),
			"Hostname": c.Hostname(),
			"IP": c.IP(),
			"IPs": c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path": c.Path(),
			"Protocal":c.Protocol(),
			"Subdomains":c.Subdomains(),
		})
	})

	//Body
	app.Post("/body",func(c *fiber.Ctx) error {
		fmt.Printf("Is JSON: %v\n", c.Is("json"))
		// fmt.Println(string(c.Body()))

		person:=Person{}
		err:=c.BodyParser(&person)
		if err!=nil{
			return err
		}
		fmt.Println(person)
		return nil
	})
	
	app.Post("/body2",func(c *fiber.Ctx) error {
		fmt.Printf("Is JSON: %v\n", c.Is("json"))
		// fmt.Println(string(c.Body()))

		data := map[string] interface{}{}
		err:=c.BodyParser(&data)
		if err!=nil{
			return err
		}
		fmt.Println(data)
		return nil
	})
	app.Listen(":8000")
}

type Person struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

// enum WorkStatus {
// 	task=1,
// 	Working,
// 	Done
// }

// const.log(WorkStatus.task)

