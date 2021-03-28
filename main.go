package main

import (
	"log"
	"os"

	"github.com/dictuantran/TaskPad/api/controllers"
	"github.com/dictuantran/TaskPad/api/util"
	"github.com/gin-contrib/cors"
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
	reportController := controllers.ReportController{}
	reportController.Init(db)
	/*
		router.Use(func(ctx *gin.Context) {
			if !util.Contains([]string{"POST", "PUT", "PATCH"}, ctx.Request.Method) {
				return
			}

			if ctx.Request.Header["Content-Length"][0] == "0" {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "Payload should not be empty"})
				ctx.AbortWithStatus(http.StatusBadRequest)
				return
			}

			if len(ctx.Request.Header["Content-Type"]) == 0 ||
				!util.Contains(ctx.Request.Header["Content-Type"], "application/json") {
				ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"message": "Content type should be application/json"})
				ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
				return
			}
		})
	*/
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0ClientID := os.Getenv("AUTH0_CLIENT_ID")
	auth0Audience := os.Getenv("AUTH0_AUDIENCE")
	auth0Callback := os.Getenv("AUTH0_CALLBACK")

	if auth0Domain == "" || auth0ClientID == "" || auth0Audience == "" || auth0Callback == "" {
		//log.Panic("AUTH0 details not found in environment variables")
	}

	/*
		dataToUIPage := gin.H{
			"AUTH0_DOMAIN":    auth0Domain,
			"AUTH0_CLIENT_ID": auth0ClientID,
			"AUTH0_AUDIENCE":  auth0Audience,
			"AUTH0_CALLBACK":  auth0Callback,
		}

		router.GET("/api/reports", middlewares.AuthMiddleware(), reportController.GetReports)

		router.NoRoute(func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", dataToUIPage)
		})
	*/

	//router.GET("/api/reports", middlewares.AuthMiddleware(), reportController.GetReports)

	router.GET("/api/reports", reportController.GetReports)

	router.Use(cors.Default())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port)
}
