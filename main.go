package main

import (
	"fmt"
	"github.com/arsidada/go-onemax/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	fmt.Println("Starting app...")

	// Defining routes for a gin webserver
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// router.OPTIONS("/", api.ReturnHeaders)
	// router.OPTIONS("/submitted_nominees", api.ReturnHeaders)
	// router.OPTIONS("/approve_nominee/:ID", api.ReturnHeaders)
	// router.OPTIONS("/reject_nominee/:ID", api.ReturnHeaders)
	// router.OPTIONS("/comments/:NomineeID", api.ReturnHeaders)

	// Routes for Approvals API
	router.GET("/submitted_nominees", api.GetSubmittedNominees)
	router.POST("/approve_nominee/:ID", api.ApproveNominee)
	router.POST("/reject_nominee/:ID", api.ApproveNominee)

	// Routes for Comments API
	router.GET("/comments/:NomineeID", api.GetComments)
	router.POST("/comments/:NomineeID", api.AddComment)

	// Starting up gin server
	router.Run()
}
