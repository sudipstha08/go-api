package main
import (
	"go/mappings"
)

// var Router *gin.Engine

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}