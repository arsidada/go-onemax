package api

import (
	"github.com/arsidada/go-onemax/psql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
  "fmt"
)

// GetSubmittedNominees is the function handler used to handle the GET request
// for the route /submitted_nominees. This calls the psql package's GetSubmittedNomineesFromDB
func GetSubmittedNominees(c *gin.Context) {
	user := c.GetHeader("user")
	status := checkAuthorization(user)
	if status == http.StatusUnauthorized {
		c.String(status, "User is unauthorized to perform this action!")
		return
	}

	result, err := psql.GetSubmittedNomineesFromDB()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, result)
}

func GetComments(c *gin.Context) {
	NomineeIDString := c.GetHeader("NomineeID")
	NomineeID, err := strconv.Atoi(NomineeIDString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	result, err := psql.GetCommentsFromDB(NomineeID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, result)
}
func AddComment(c *gin.Context) {
  fmt.Println("1. starting add comment")
	useridString := c.GetHeader("user")
	userid , err := strconv.Atoi(useridString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
  fmt.Println("2. ")
	NomineeIDString := c.Param("NOMID") // generate ID
	NomineeID , err := strconv.Atoi(NomineeIDString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	content := c.GetHeader("content")
  fmt.Println("3. ")

  _, err2 := psql.AddCommentDB(userid, NomineeID, content)
	if err2 != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
  fmt.Println("4. ")
	c.JSON(http.StatusOK, "the comment has been added")
}

// ApproveNominee is the function hanlder user to handle the POST request
// for the route /approve_nominee/:ID. This calls the psql package's ApproveNomineeDB
func ApproveNominee(c *gin.Context) {
	user := c.GetHeader("user")
	status := checkAuthorization(user)
	if status == http.StatusUnauthorized {
		c.String(status, "User is unauthorized to perform this action!")
		return
	}

	IDString := c.Param("ID")
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	_, err = psql.ApproveNomineeDB(ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "Nominee approved successfully!")
}

// RejectNominee is the function hanlder user to handle the POST request
// for the route /rejecte_nominee/:ID. This calls the psql package's RejectNomineeDB
func RejectNominee(c *gin.Context) {
	user := c.GetHeader("user")
	status := checkAuthorization(user)
	if status == http.StatusUnauthorized {
		c.String(status, "User is unauthorized to perform this action!")
		return
	}

	IDString := c.Param("ID")
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	_, err = psql.RejectNomineeDB(ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "Nominee rejected successfully!")
}

// checkAuthorization is a helper function to check if the provided user param matches
// what we have declared as authorized users
func checkAuthorization(user string) int {
	if user == "" && user != "105364027055888" && user != "111223425403387795098" {
		return http.StatusUnauthorized
	}
	return http.StatusOK
}
