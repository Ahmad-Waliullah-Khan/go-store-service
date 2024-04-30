package controllers

import (
	"encoding/json"
	"strconv"

	"store-service/models"

	"github.com/astaxie/beego"
)

type ItemController struct {
	beego.Controller
}

func (c *ItemController) AddItem() {
	var newItem models.Item
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &newItem); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body. Please provide a valid JSON object with 'name' and 'price' fields."}
		c.ServeJSON()
		return
	}

	if newItem.Name == "" || newItem.Price == 0 {
		c.Data["json"] = map[string]string{"error": "Missing or invalid fields in request body. 'name' and 'price' fields are required."}
		c.ServeJSON()
		return
	}

	if err := models.AddItem(newItem); err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to add item to the store. Please try again later."}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = newItem
	c.ServeJSON()
}

func (c *ItemController) RemoveItem() {
	itemID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err := models.RemoveItem(itemID); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"message": "Item removed successfully"}
	c.ServeJSON()
}

func (c *ItemController) ListItems() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	pageSize, _ := strconv.Atoi(c.GetString("pageSize", "10"))

	items, err := models.ListItems(page, pageSize)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = items
	c.ServeJSON()
}

func (c *ItemController) ShowItem() {
	itemID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	item, err := models.GetItem(itemID)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = item
	c.ServeJSON()
}

func (c *ItemController) UpdateItem() {
	itemID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var updatedItem models.Item
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &updatedItem); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	if err := models.UpdateItem(itemID, updatedItem); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = updatedItem
	c.ServeJSON()
}
