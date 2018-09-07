package api

import (
	"bytes"
	"github.com/arsidada/go-onemax/psql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// GetComments is the function handler to handle the GET request
// for fetching all comments associated with a NomineeID
func GetComments(c *gin.Context) {
	NomineeIDString := c.Param("NomineeID")
	NomineeID, err := strconv.Atoi(NomineeIDString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	result, err := psql.GetCommentsFromDB(NomineeID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	newResult := make([]psql.Comment, len(result))
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		newResult[i], newResult[j] = result[j], result[i]
	}

	c.JSON(http.StatusOK, newResult)
}

// AddComment is the function handler to create a comment for the
// given NomineeID and associate it with the user who has supplied the comment
func AddComment(c *gin.Context) {
	username := c.GetHeader("username")

	NomineeIDString := c.Param("NomineeID") // generate ID
	NomineeID, err := strconv.Atoi(NomineeIDString)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	buf := new(bytes.Buffer)
	body := c.Request.Body
	buf.ReadFrom(body)
	content := buf.String()

	_, err2 := psql.AddCommentDB(username, NomineeID, content)
	if err2 != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "Comment added successfully!")
}

// checkAuthorization is a helper function to check if the provided user param matches
// what we have declared as authorized users
func checkAuthorization(user string) int {
	if user == "105364027055888" || user == "111223425403387795098" {
		return http.StatusOK
	}
	return http.StatusUnauthorized
}
