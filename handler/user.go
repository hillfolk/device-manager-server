package handler

import (
	"net/http"
	"time"
	"strconv"
	"context"
	"log"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model "github.com/hillfolk/app-manager-server/model"
)

func (h *Handler) creatUser(c echo.Context) error {
	
	return nil
}

func (h *Handler) readUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	filter := bson.D{{"id", id}}
	coll := h.DB.Collection("user")
	var u model.User
	err := coll.FindOne(
		context.Background(),
		filter,).Decode(&u)
	
	if err != nil {
		log.Fatal(err)
	}

	
	
	return c.JSON(http.StatusOK, u)
}


func (h *Handler) Signup(c echo.Context) (err error) {
	
	u := &model.User{}

	if err = c.Bind(u); err!= nil {
		return
	}

	//Validate
	if u.Email == "" || u.Password == "" {

		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}		
	}
	
	u.Id = xid.New().String()
	u.Created = time.Now()
	_,err = h.DB.Collection("user").InsertOne(context.Background(),u)
	
	return c.JSON(http.StatusCreated,u)
	
}


func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid value"}
	}

	// Find user


	if err = h.DB.Collection("user").FindOne(context.Background(),bson.M{"email": u.Email, "password": u.Password},).Decode(&u); err != nil {
		
		if err == mongo.ErrNilDocument {
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
	claims["id"] = u.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password

	return c.JSON(http.StatusOK, u.Token)
}


func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
