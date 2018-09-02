package main

import (
	"fmt"
	"github.com/arsidada/go-onemax/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting app...")

	// Defining routes for a gin webserver
	router := gin.Default()

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
