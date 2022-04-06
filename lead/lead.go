package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/moaabid/go-fiber-crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBconn

	var leads []Lead

	db.Find(&leads)

	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn

	var lead Lead

	db.Find(&lead, id)

	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with given ID")
	}
	return c.JSON(lead)

}

func NewLead(c *fiber.Ctx) error {
	db := database.DBconn

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&lead)
	return c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn

	var lead Lead

	db.First(&lead, id)

	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with given ID")
	}

	db.Delete(&lead)
	return c.SendString("Lead Successfully deleted!")
}
