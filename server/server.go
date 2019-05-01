package server


import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo@~1.0.0"
)


func RunServer(port string,dburl string){
	e := echo.New()
	e.Debug = true
	// Create Mongodb Connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	
	
	// Server header
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ServerHeader)
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "app-manager-server")
	})
	/* Post */	
	e.GET("/posts",getPosts(client))
	
	
	e.Logger.Fatal(e.Start(port))
	
}

