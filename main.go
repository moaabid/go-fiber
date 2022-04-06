package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/moaabid/go-fiber-crm-basic/database"
	"github.com/moaabid/go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/getlead/:id", lead.GetLead)
	app.Get("/api/v1/getleads", lead.GetLeads)
	app.Post("/api/v1/addnewlead", lead.NewLead)
	app.Delete("/api/v1/deletelead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error

	database.DBconn, err = gorm.Open("sqlite3", "lead.db")

	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Database connection opened")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated Successfully")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	defer database.DBconn.Close()
}
