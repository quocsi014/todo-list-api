package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"todolist/modules/item/business"
	"todolist/modules/item/handler"
	"todolist/modules/item/repository"
)

func main() {

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	r := gin.Default()

	r.GET("/hello_world", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world")
	})

	itemRepo := repository.NewItemRepository(db)
	itemBus := business.NewItemBusiness(itemRepo)
	itemHandler := handler.NewItemHandler(itemBus)
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			itemHandler.SetupRoute(items)
		}
	}

	r.Run()
}
