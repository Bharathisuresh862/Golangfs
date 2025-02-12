package main

import (
	//"context"
	"fmt"
	//"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config

// Database variables

// Model Car for Collection "cars"
type Car struct {
	ID     string
	Number string
	Model  string
	Type   string
}

// POST /cars
func createCar(c *gin.Context) {
	var jbodyCar Car

	// Bind JSON body to jbodyCar
	if err := c.BindJSON(&jbodyCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdCar Car = Car{ID: jbodyCar.ID, Number: jbodyCar.Number, Model: jbodyCar.Model, Type: jbodyCar.Type}

	// Return created car
	c.JSON(http.StatusCreated, gin.H{
		"message": "Car created successfully",
		"car":     createdCar,
	})
}

// GET /cars
/*func readAllCars(c *gin.Context) {
	var cars []Car = []Car{
		{ID: "id001", Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"},
		{ID: "id002", Number: "KA40 Z2233", Model: "Toyoto Innova", Type: "SUV"},
	}

	c.JSON(http.StatusOK, cars)
}

// GET /cars/:id
func readCarById(c *gin.Context) {
	id := c.Param("id")
	var car Car = Car{ID: id, Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"}
	c.JSON(http.StatusOK, car)
}*/
func readCars(c *gin.Context) {
    id := c.Param("id")

    cars := []Car{
        {ID: "id001", Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"},
        {ID: "id002", Number: "KA40 Z2233", Model: "Toyota Innova", Type: "SUV"},
    }

    if id == "" {
        // If no ID is provided, return all cars
        c.JSON(http.StatusOK, cars)
        return
    }

    // If ID is provided, find the specific car
    for _, car := range cars {
        if car.ID == id {
            c.JSON(http.StatusOK, car)
            return
        }
    }

    // If no car is found with the given ID
    c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
}

// PUT /cars/:id
func updateCar(c *gin.Context) {
	id := c.Param("id")
	var jbodyCar Car
	if err := c.BindJSON(&jbodyCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var oldCar Car = Car{ID: id, Number: "KA03 A1010", Model: "Maruti Suzuki", Type: "CUV"}

	oldCar.Number = jbodyCar.Number
	oldCar.Model = jbodyCar.Model
	oldCar.Type = jbodyCar.Type

	// Return updated car
	c.JSON(http.StatusOK, gin.H{
		"message": "Car updated successfully",
		"car":     oldCar,
	})
}

// DELETE /cars/:id
func deleteCar(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}

func 

func main() {
	// Set up Gin router
	r := gin.Default()
	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	r.POST("/cars", createCar)
	r.GET("/cars", readAllCars)
	r.GET("/cars/:id", readCarById)
	r.PUT("/cars/:id", updateCar)
	r.DELETE("/cars/:id", deleteCar)

	// Start server
	r.Run(":8080")
}
