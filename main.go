package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const creditScoreMin = 500
const creditScoreMax = 900

type credit_rating struct {
	CreditRating int `json:"credit_rating"`
}

type CreateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Api1Input struct {
	RunequestId string `json:"requestId"`
	Data        struct {
		Value1 int `json:"value1"`
		Value2 int `json:"value2"`
	} `json:"data"`
}

type Api1Output struct {
	RunequestId string `json:"requestId"`
	Data        struct {
		Sum int `json:"sum"`
	} `json:"data"`
}

func main() {
	fmt.Println("Hello word")

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/creditscore", getCreditScore) // create item
		v1.POST("/creditscore", postCreditScore)
		v1.POST("/api1", postBuoi1)
		// v1.GET("/items", getListOfItems(db))        // list items
		// v1.GET("/items/:id", readItemById(db))      // get an item by ID
		// v1.PUT("/items/:id", editItemById(db))      // edit an item by ID
		// v1.DELETE("/items/:id", deleteItemById(db)) // delete an item by ID
	}

	router.Run()
}

func getCreditScore(c *gin.Context) {
	var creditRating = credit_rating{
		CreditRating: (rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin),
	}

	c.IndentedJSON(http.StatusOK, creditRating)
}

func postCreditScore(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	// book := models.Book{Title: input.Title, Author: input.Author}
	// models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func postBuoi1(c *gin.Context) {
	// Validate input
	var b1 Api1Input
	if err := c.ShouldBindJSON(&b1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	// book := models.Book{Title: input.Title, Author: input.Author}
	// models.DB.Create(&book)
	sum := b1.Data.Value1 + b1.Data.Value2

	oput := Api1Output{RunequestId: b1.RunequestId}
	oput.Data.Sum = sum

	c.JSON(http.StatusOK, oput)
}
