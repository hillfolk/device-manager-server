package handler

import (
	"net/http"
	"time"
	"context"
	_"log"
	"github.com/rs/xid"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	_"go.mongodb.org/mongo-driver/mongo"
	model "github.com/hillfolk/app-manager-server/model"
)



func (h *Handler) CreatePost(c echo.Context) (err error) {
	p := &model.Post{}

	if err = c.Bind(p); err!= nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid value"} 
	}

	p.Id = xid.New().String()	
	p.Created  = time.Now()
	p.Updated = p.Created

	if p.Title == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid title"}		
	}
	resultOne,err := h.DB.Collection("post").InsertOne(context.Background(),p)

	err = h.DB.Collection("post").FindOne(
			context.Background(),
			bson.D{{"id",resultOne.InsertedID}},
	).Decode(&p)

	
	return c.JSON(http.StatusCreated,p)
}
