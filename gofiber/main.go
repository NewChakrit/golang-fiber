package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main () {

	app:=fiber.New()

	//GET
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("GET: Hello world")
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

	app.Listen(":8000")
}

type Person struct{
	Id int `json:"id"`
	Name string `json:"name"`
}


