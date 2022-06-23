package main

import (
	"github.com/gofiber/fiber/v2"
)

func main () {
	// app:=mux.NewRouter()
	
	// app.HandleFunc("/hello/{id}", Hello).Methods(http.MethodGet)
	
	// http.ListenAndServe(":8000", app)

	app:=fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Listen(":8000")
}

// func Hello  (w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	fmt.Println(id)
// }

