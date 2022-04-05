package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type ALLINONE struct {
	ID    string `json:"id"`
	Item  string `json:"item"`
	Price int    `json:"price"`
}

var LISTS = []ALLINONE{
	{ID: "1", Item: "Go-Book", Price: 2000},
	{ID: "2", Item: "Python-Book", Price: 1000},
	{ID: "3", Item: "C-Book", Price: 1500},
}

func getLISTS(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, LISTS)
}
func addLISTS(context *gin.Context) {
	var newItem ALLINONE
	if err := context.BindJSON(&newItem); err != nil {
		return
	}
	LISTS = append(LISTS, newItem)
	context.IndentedJSON(http.StatusCreated, newItem)
}
func getLISTSByID(id string) (*ALLINONE, error) {
	for i, t := range LISTS {
		if t.ID == id {
			return &LISTS[i], nil
		}
	}
	return nil, errors.New("id not found")
}
func getLIST(context *gin.Context) {
	id := context.Param("id")
	LIST, err := getLISTSByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, LIST)
}
func main() {
	router := gin.Default()
	router.GET("/LISTS", getLISTS)
	router.GET("/LISTS/:id", getLIST)
	router.POST("/LISTS", addLISTS)
	router.Run("localhost:9001")

}
