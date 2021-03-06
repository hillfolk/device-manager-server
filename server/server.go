package server


import (
	"context"
	_"io"
	"net/http"
	_"regexp"
	_"os"
	_"log"
	"time"
	"github.com/hillfolk/device-manager-server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)


func RunServer(port,db_url string){
	e := echo.New()
	e.Debug = false
	// Create Mongodb Connection
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://"+db_url))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)
	
	db := client.Database("dms")

	defer client.Disconnect(ctx)
	
	// Server header
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/"{
				return true
			}
			if c.Path() == "/login" || c.Path() == "/signup" {

				return true
			}
			if c.Path() == "/devices/reg" || c.Path() == "/devices/update" {	
				return true
			}

			return false
		},
	}))
	e.Use(middleware.Recover())
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "device-manager-server")
	})

	h := &handler.Handler{DB: db}


	/* User */
	e.POST("/signup",h.Signup)
	e.POST("/login",h.Login)
	
	/* Post */
	e.POST("/posts", h.CreatePost)
	e.PUT("/posts/:id",h.UpdatePost)
	e.GET("/posts", h.ReadPosts)
	e.GET("/posts/:id",h.ReadPost)
	
	/* Device */
	e.POST("/devices",h.CreateDevice)
	
	e.POST("/devices/reg",h.RegClientDevice)
	e.PUT("/devices/update",h.UpdateClientDevice)
	
	e.PUT("/devices/:id",h.UpdateDevice)
	e.GET("/devices",h.ReadDevices)
	e.GET("/devices/:id",h.ReadDevice)
	e.DELETE("/devices/:id",h.DeleteDevice)

	e.Logger.Fatal(e.Start(port))
}

