package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const url string -9*"https://api.nasa.gov/planetary/apod?"

func main() {
	//Get a default gin router
	app := gin.Default()

	//Set project to release mode for deployment
	gin.SetMode(gin.ReleaseMode)

	//Preload all the html files so accessing them will be much faster
	app.LoadHTMLGlob("templates/*")

	//Serve the static files in the public folder.
	app.Static("/public", "public")

	//Load the dependency to read in .env files.
	errHandle(godotenv.Load())

	app.GET("/", renderCurrentAPODPicture)
	app.GET("/api/:date", getAPODJSON)
	app.GET("/apod/:date", renderPictureByDate)
	app.POST("/changeDate", redirectToAPODPicture)
	app.NoRoute(render404Page) //Handle 404 errors

	fmt.Println("Server is up!")
	app.Run()
}

func render404Page(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"Errorcode":        http.StatusNotFound,
		"ErrorDescription": "Page not found!",
	})
}

func getAPODJSON(c *gin.Context) {
	webData, response := loadAPODInfo("date="+c.Param("date"), "&api_key="+os.Getenv("API_KEY"))

	if response.StatusCode != 200 {
		c.JSON(response.StatusCode, gin.H{
			"errmsg": "No data found for date: " + c.Param("date"),
		})
	} else {
		c.JSON(response.StatusCode, webData)
	}
}

func redirectToAPODPicture(c *gin.Context) {
	//Once this function is hit when the user posts a date, redirect them to the above route with the date
	//they entered.
	c.Redirect(http.StatusFound, "/apod/"+c.PostForm("date"))
}

func renderCurrentAPODPicture(c *gin.Context) {
	//Retrieve the current date, and format into the desired style using the below date as a placeholder.
	//This is needed because of possible time zone issues with the NASA api and EST, so retrieving the current
	//data late at night results in NASA thinking I am trying to receive TOMORROW'S information, which then
	//results in "webData" being empty.
	date := time.Now().Format("2006-01-02")

	//Access the NASA api and retrieve the most recent picture (today's date) using the date variable.
	webData, response := loadAPODInfo("date="+date, "&api_key="+os.Getenv("API_KEY"))

	renderHTML(c, response.StatusCode, webData, webData["media_type"] == "image")
}

func renderPictureByDate(c *gin.Context) {
	//Access the NASA api and retrieve the picture according to the date the user entered.
	webData, response := loadAPODInfo("date="+c.Param("date"), "&api_key="+os.Getenv("API_KEY"))

	fmt.Println("repsonse:", response.StatusCode)
	//If the response returned a status code other than 200, render the 404 html page.
	if response.StatusCode != 200 {
		c.HTML(response.StatusCode, "404.html", gin.H{
			"Errorcode":        response.StatusCode,
			"ErrorDescription": "Bad Request!",
		})
	} else {
		//Otherwise, render the index.html file with the necessary information.
		renderHTML(c, response.StatusCode, webData, webData["media_type"] == "image")
	}
}

func renderHTML(c *gin.Context, StatusCode int, webData map[string]interface{}, isAPicture bool) {
	c.HTML(
		StatusCode, "index.html", gin.H{
			"URL":         webData["url"],
			"Description": webData["explanation"],
			"Date":        webData["date"],
			"Title":       webData["title"],
			"Author":      webData["copyright"],
			"IsAPicture":  isAPicture,
		},
	)
}

func loadAPODInfo(date string, apiKey string) (map[string]interface{}, *http.Response) {
	response, err := http.Get(url + date + apiKey)
	errHandle(err)

	//Read the response from the endpoint into a byte array.
	data, err := ioutil.ReadAll(response.Body)
	errHandle(err)

	webData := make(map[string]interface{})

	//Unload the byte array information into a hashmap.
	errHandle(json.Unmarshal(data, &webData))

	return webData, response
}

func errHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
