package main

import (
	"log"
	"os"

	"github.com/dictuantran/TaskPad/api/controllers"
	"github.com/dictuantran/TaskPad/api/util"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db, err := util.GetDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	taskController := controllers.TaskController{}
	taskController.Init(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port)
}
