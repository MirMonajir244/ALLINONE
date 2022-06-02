package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	BookName           string `json:"book_name"`
	AuthorName         string `json:"author_name"`
	Price              string `json:"price"`
	PublicationCompany string `json:"publication_company"`
	Code               string `json:"book_code"`
}

/*
func New() *List {
	return &List{}
}
*/

var Books []Item

func init() {
	Books = append(Books, Item{
		BookName:           "Golang Guide",
		AuthorName:         "Mir Monajir",
		Price:              "$300",
		PublicationCompany: "Hacktech",
		Code:               "CD001",
	})
	Books = append(Books, Item{
		BookName:           "Docker",
		AuthorName:         "Jhonathan",
		Price:              "$200",
		PublicationCompany: "MicroWorld",
		Code:               "CD002",
	})
	Books = append(Books, Item{
		BookName:           "Kubernetes",
		AuthorName:         "Jimmy",
		Price:              "$250",
		PublicationCompany: "MicroWorld",
		Code:               "CD003",
	})
}

//To add Items
func Add(c *gin.Context) {
	var newItem Item
	err := c.BindJSON(&newItem)
	if err != nil {
		return
	}
	Books = append(Books, newItem)
	c.IndentedJSON(http.StatusCreated, Books)
}

//To Fetch all the items
func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}

//To Fetch Items By Name
func GetItemByCode(code string) (*Item, error) {
	for i, t := range Books {
		if t.Code == code {
			return &Books[i], nil
		}
	}
	return nil, errors.New("Item Not Found")
}

//Handler of GetItemByCode()

func GetItem(c *gin.Context) {
	code := c.Param("code")
	it, err := GetItemByCode(code)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item Not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, it)
}

//To Delete The Item By code
func DelItemByCode(code string) error {
	for i, t := range Books {
		if t.Code == code {
			Books = append(Books[:i], Books[i+1:]...)
			return nil
		}
	}
	return errors.New("Item Not Found")
}

//Handler for DelItemByCode()
func Delitem(c *gin.Context) {
	code := c.Param("code")
	it := DelItemByCode(code)
	if it != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item Not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Item deleted"})
}

func main() {
	router := gin.Default()

	router.GET("/Book", GetAll)           //GET endpoint
	router.POST("/Book", Add)             //POST Endpoint
	router.GET("/Book/:code", GetItem)    //GET endpoint by code
	router.DELETE("/Book/:code", Delitem) //DELETE endpoint by code
	router.Run("localhost:8091")          //to run the API

}
