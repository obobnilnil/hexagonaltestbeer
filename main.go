package main

import (
	"app/database"
	"app/handler"
	"app/repository"
	"app/service"
	"app/transaction"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()
	conn := database.MongoDB()
	defer conn.Client().Disconnect(context.Background())

	r := repository.NewRepository(db)
	t := transaction.NewTransaction(conn)
	s := service.NewService(r, t)
	h := handler.NewHandler(s)

	router := gin.Default()

	router.GET("/get", h.Get)
	router.POST("/add", h.Add)
	router.PUT("/update/:id", h.Update)
	router.DELETE("/delete/:id", h.Delete)

	if err := router.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
