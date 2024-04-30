package routers

import (
	"store-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/item/add", &controllers.ItemController{}, "post:AddItem")
	beego.Router("/item/:id/remove", &controllers.ItemController{}, "delete:RemoveItem")
	beego.Router("/items", &controllers.ItemController{}, "get:ListItems")
	beego.Router("/item/:id", &controllers.ItemController{}, "get:ShowItem")
	beego.Router("/item/:id/update", &controllers.ItemController{}, "put:UpdateItem")
}
