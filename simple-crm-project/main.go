package main

import (
	"fmt"
	"log"
	"simple-crm/database"
	"simple-crm/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	fmt.Println("hello")

	app := fiber.New()

	initDB()
	setRoutes(app)
	app.Listen("8080")

}

func setRoutes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDB(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		log.Panic("failed to connect Sqlite DB")
	}

	fmt.Println("Connection opened to database")

	// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
	database.DBConn.AutoMigrate(&lead.Lead{})

	fmt.Println("DB migated")


}