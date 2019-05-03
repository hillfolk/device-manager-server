package handler

import (
	"net/http"
	_ "time"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	_ "go.mongodb.org/mongo-driver/bson"
	model "github.com/hillfolk/app-manager-server/model"
)


func (h *Handler) Signup(c echo.Context) (err error) {
	
	u := &model.User{}
	if err = c.Bind(u); err!= nil {
		return
	}

	//Validate
	if u.Email == "" || u.Password == "" {

		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}		
	}

	_,err = h.DB.Collection("user").InsertOne(context.Background(),u)
	return c.JSON(http.StatusCreated,u)
	
}


func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	/*
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user

	defer DB.Close()
	if err = db.
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
*/
	return c.JSON(http.StatusOK, "login")
}


func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
