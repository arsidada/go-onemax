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
	router.GET("/submitted_nominees", api.GetSubmittedNominees)
	router.GET("/nominee_comments", api.GetComments)
	router.POST("/add_comment/:NOMID", api.AddComment)
	router.POST("/approve_nominee/:ID", api.ApproveNominee)
	router.POST("/reject_nominee/:ID", api.ApproveNominee)
	router.Run()
}
