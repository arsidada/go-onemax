package psql

import (
	"crypto/tls"
	"github.com/go-pg/pg"
	"os"
  "fmt"
	// "github.com/go-pg/pg/orm"
)

// Nomination is an object definition to store values from the Nomination table
type Nomination struct {
	ID          int
	Name        string
	Description string
	Country     string
	Province    string
	Status      string
	Image       string
	Duas        int
}

type Comment struct {
	ID          int
	Nomineeid   int
	Userid      int
	Content     string
	Createdat   string
}



// Global var to hold the database object
var db *pg.DB

func init() {
	// Import credentials from env vars
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	addr := os.Getenv("DB_ADDR")
	database := os.Getenv("DB_DB")

	// Connect to the database using the creds
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Addr:     addr,
		Database: database,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
}

// GetSubmittedNomineesFromDB queries the Nomination table and returns all Nominees
// with status == "submitted". If there is an error then we return the error
// and an empty Nomination slice
func GetSubmittedNomineesFromDB() ([]Nomination, error) {
	result := make([]Nomination, 0)
	err := db.Model(&result).
		Where("status = ?", "submitted").
		Select()
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetCommentsFromDB(NomineeID int) ([]Comment, error) {
	result := make([]Comment, 0)
	err := db.Model(&result).
		Where("nomineeid = ?", NomineeID).
		Select()
	if err != nil {
		return result, err
	}
	return result, nil
}
func AddCommentDB(userID int, nomineeID int, content string) (int, error){
  fmt.Println("going to add the comment ", content, " userid: ", userID, " nomineeID: ", nomineeID)
  err := db.Insert(&Comment{
    Userid : userID,
    Nomineeid : nomineeID,
    Content : content,
  })
  if(err != nil){
    return -1, err
  }

  return 0, nil
}

// ApproveNomineeDB uses the ID parameter to updates a record's status value
// from 'submitted' to 'approved'. We return a 200 if the update is successful.
// Otherwise we return a 500 and an error
func ApproveNomineeDB(ApprovalID int) (int, error) {
	nomination := &Nomination{ID: ApprovalID}
	err := db.Select(nomination)
	if err != nil {
		return 500, err
	}

	nomination.Status = "approved"
	err = db.Update(nomination)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

// RejectNomineeDB uses the ID parameter to updates a record's status value
// from 'submitted' to 'rejected'. We return a 200 if the update is successful.
// Otherwise we return a 500 and an error
func RejectNomineeDB(ApprovalID int) (int, error) {
	nomination := &Nomination{ID: ApprovalID}
	err := db.Select(nomination)
	if err != nil {
		return 500, err
	}

	nomination.Status = "rejected"
	err = db.Update(nomination)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
