package lead

import (
	"fmt"
	"simple-crm/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Reason of defining gorm.Model -> Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models

/*
type Model struct { // size=64 (0x40)
ID        uint `gorm:"primary_key"`
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt *time.Time `sql:"index"`
}
*/

type Lead struct {

	gorm.Model
	Name string `json:"name"`
	Company string `json:"company"`
	Email string `json:"email"`
	Phone int `json:"phone"`

}

func GetLeads(c *fiber.Ctx){
	fmt.Println("get leads")

	db := database.DBConn
	var leads []Lead
	db.Find(&leads)

	// JSON converts any interface or string to JSON. This method also sets the content header to application/json.
	c.JSON(leads)
}


func GetLead(c *fiber.Ctx){
	id := c.Params("id")

	db := database.DBConn

	var lead Lead

	db.Find(&lead, id)
	c.JSON(lead)
}


func NewLead(c *fiber.Ctx) {
	db := database.DBConn

	lead := new(Lead)

	// BodyParser binds the request body to a struct.
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn

	var lead Lead

	db.First(&lead, id)

	if lead.Name == ""{
		c.Status(500).Send("No lead found with ID")
		return
	}

	db.Delete(&lead)

	c.Send("Lead deleted.")


}