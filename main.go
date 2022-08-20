package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Movies struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
	ImdbId string `json:"imdbId"`
	Type   string `json:"Type"`
}

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	// Ratings
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbId     string `json:"imdbId"`
	Type       string `json:"Type"`
	BoxOffice  string `json:"boxOffice"`
	Production string `json:"production"`
	Website    string `json:"website"`
	Response   string `json:"Response"`
	Error      string `json:"Error"`
}

type Search struct {
	Search []Movies `json:"search"`
}

// func init() {

// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	router := gin.Default()
	router.GET("/movies", getMovies)
	router.GET("/movie", getMovie)
	router.Run(":" + port)

}
func getMovies(c *gin.Context) {
	search, ok := c.GetQuery("search")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing search query parameter."})
		return
	}

	dotenv := os.Getenv("API_KEY")

	response, err := http.Get(fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&s=%s", dotenv, search))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Search
	json.Unmarshal(responseData, &responseObject)

	c.JSON(http.StatusOK, responseObject.Search)
}

func getMovie(c *gin.Context) {
	search, ok := c.GetQuery("search")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing search query parameter."})
		return
	}

	dotenv := os.Getenv("API_KEY")

	response, err := http.Get(fmt.Sprintf("https://www.omdbapi.com/?apikey=%s=%s", dotenv, search))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Movie
	json.Unmarshal(responseData, &responseObject)

	c.JSON(http.StatusOK, responseObject)
}
