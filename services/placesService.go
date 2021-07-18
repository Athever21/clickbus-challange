package services

import (
	"clickbus/models"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type PlaceBody struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Slug  string `json:"slug" xml:"slug" form:"slug"`
	City  string `json:"city" xml:"city" form:"city"`
	State string `json:"state" xml:"state" form:"state"`
}

func GetAllPlaces(c *fiber.Ctx) error {
	p := new(PlaceBody)
	var query map[string]string

	if err := c.QueryParser(p); err != nil {
		return c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	pJson, err := json.Marshal(p)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	err = json.Unmarshal(pJson, &query)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	for k, v := range query {
		if v == "" {
			delete(query, k)
		}
	}

	places, err := models.GetAllPlaces(query)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	if places == nil {
		return c.JSON([]string{})
	}

	return c.JSON(places)
}

func GetPlace(c *fiber.Ctx) error {
	id := c.Params("id")

	place, err := models.GetPlace(id)

	if err != nil {
		return c.Status(404).JSON(&fiber.Map{"error": "not found"})
	}

	return c.JSON(place)
}

func CreatePlace(c *fiber.Ctx) error {
	p := new(PlaceBody)

	if err := c.BodyParser(p); err != nil {
		return c.Status(400).JSON(&fiber.Map{"error": "invalid body"})
	}

	if p.Name == "" || p.Slug == "" || p.City == "" || p.State == "" {
		return c.Status(400).JSON(&fiber.Map{"error": "missing fields"})
	}

	newPlace, err := models.CreatePlace(p.Name, p.Slug, p.City, p.State)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"error": err})
	}

	created, nil := models.GetPlaceNew(newPlace.InsertedID)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"error": err})
	}

	return c.JSON(created)
}

func UpdatePlace(c *fiber.Ctx) error {
	id := c.Params("id")
	p := new(PlaceBody)
	var update map[string]string

	if err := c.BodyParser(p); err != nil {
		c.Status(400).JSON(&fiber.Map{"error": "invalid body"})
	}

	pJson, err := json.Marshal(p)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	err = json.Unmarshal(pJson, &update)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	for key, value := range update {
		if value == "" {
			delete(update, key)
		}
	}

	if len(update) == 0 {
		return c.Status(400).JSON(&fiber.Map{"error": "no fields to update"})
	}

	update["updatedAt"] = time.Now().String()
	_, err = models.UpdatePlace(update, id)

	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{"error": err})
	}

	changed, err := models.GetPlace(id)

	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	return c.JSON(changed)
}

func DeletePlace(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeletePlace(id)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"error": "something went wrong"})
	}

	return c.JSON(&fiber.Map{"success": true})
}
