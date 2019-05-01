package server

import (
	"net/http"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)


func createApp(c echo.Context, client *mongo.Client){
	

	
}

func readApp(c echo.Context, client *mongo.Client){
}

func updateApp(c echo.Context, client *mongo.Client){
}

func deleteApp(c echo.Context, client *mongo.Client){
}



//Version
func createVersion(c echo.Context, client *mongo.Client){
	
}


func updateVersion(c echo.Context, client *mongo.Client){
	
}


func delteVersion(c echo.Context, client *mongo.Client){

}


func readVersion(c echo.Context, client *mongo.Client){

}


//Post

func createPosts(c echo.Context, client *mongo.Client){
}

func readPosts(c echo.Context)error{
		return c.String(http.StatusOK, "readPosts")
}


func updatePosts(c echo.Context, client *mongo.Client){


}


func deletePosts(c echo.Context, client *mongo.Client){

}
