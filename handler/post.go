package handler

import (
	"net/http"
	"time"
	"strconv"
	"context"
	"log"
	"github.com/rs/xid"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	model "github.com/hillfolk/device-manager-server/model"
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

func (h *Handler) ReadPosts(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	skip := int64((page-1) * limit)
	findOptions.SetSkip(skip)
	
	posts := []*model.Post{}

	cur, err := h.DB.Collection("post").Find(
		context.Background(),bson.D{},findOptions)

	if err != nil {
			log.Fatal(err)
		}
	
	for cur.Next(context.TODO()) {

		var elem model.Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts,&elem)
	}

	
	
	
	
	return c.JSON(http.StatusOK,posts)
}


func (h *Handler) ReadPost(c echo.Context) (err error) {

	id := c.Param("id")
	
	p := &model.Post{}

	filter := bson.D{{"id", id}}

	err = h.DB.Collection("post").FindOne(
		context.Background(),
		filter,).Decode(&p)
	if err != nil {
		log.Fatal(err)
		// TODO:Error response by case
	}
	return c.JSON(http.StatusOK,p)
}
