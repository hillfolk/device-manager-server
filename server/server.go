package server


import (
	"context"
	_"fmt"
	_"io"
	"net/http"
	_"os"
	"time"
	"github.com/hillfolk/app-manager-server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)


func RunServer(port string){
	e := echo.New()
	e.Debug = true
	// Create Mongodb Connection
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)
	
	db := client.Database("ams")
	
	defer client.Disconnect(ctx)
	
	// Server header
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "device-manager-server")
	})

	h := &handler.Handler{DB: db}


	/* Post */
	e.POST("/signup",h.Signup)
	e.POST("/login",h.Login)
	
	
	
	
	e.Logger.Fatal(e.Start(port))
	
}

