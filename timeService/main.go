package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var idCounter = 0

// https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html
// give me date must implement "handlerFunc" interface
func giveMeTime(context *gin.Context) {
	result := fmt.Sprintf("%s", time.Now().Format("15:04:05"))
	fmt.Printf("[debug] %s\n", result)
	idCounter++
	context.JSON(http.StatusOK, gin.H{"id" : idCounter, "time" : result })
}

// gin is a framework for web servicesc
// https://github.com/gin-gonic/gin
func main(){
	fmt.Println("hello from time service")
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	// mount router 
	router.GET("/getTime", giveMeTime)
	

	// By default it serves on :8015 unless a
	// PORT environment variable was defined.
	router.Run(":8015")
}
