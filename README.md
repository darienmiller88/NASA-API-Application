   # NASA-API-Application
  ![](https://img.shields.io/badge/made%20by-DarienMiller-blue)
  ![](https://img.shields.io/badge/Golang-46%25-blue)
  ![](https://img.shields.io/badge/Golang-1.14-yellow)
  ![](https://img.shields.io/badge/HTML%2B%20CSS-48%25-red)
  ![](https://img.shields.io/badge/test-passing-green)
   <p align = "center">
  <img width="600" alt="currentdate" src="https://user-images.githubusercontent.com/32966645/97977040-91e77900-1d99-11eb-8b6d-3b26804de828.PNG">
   </p>
  
  ## Description

Web Application to extract data from the NASA API and display them on dynamically on an HTML file! Written using the "Gin" web framework. View it at https://nasa-apod-application.herokuapp.com/. Click [here](https://github.com/gin-gonic/gin) to read more on the web framework I used for this project. So far, it allows the user to type in a specific date, and the app will re route the user to a new page with the NASA API picture and information for that date given, or an error page if invalid information is given. Here's a demo!

![nasa-apod-api](https://user-images.githubusercontent.com/32966645/97978044-0969d800-1d9b-11eb-9afb-2ccdba179385.gif)

   ### Possible error pages:
   
   <img width="600" alt="400error" src="https://user-images.githubusercontent.com/32966645/97978590-c3614400-1d9b-11eb-9922-0fdd8009ab76.PNG">
   
   <img width="600" alt="404error" src="https://user-images.githubusercontent.com/32966645/97978596-c65c3480-1d9b-11eb-89da-968e636a1bd0.PNG">


   ## Installation

```
   go build 
   go mod vender //if you desire to have a folder including the dependencies, otherwise ignore
   .\fresh //to restart server on change or
   go run main.go //to run the server normally
```

  ## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

  ## License
[MIT](https://choosealicense.com/licenses/mit/)
