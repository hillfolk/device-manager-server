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

func (h *Handler) CreateDevice(c echo.Context) (err error){

	d := &model.Device{}
	
	
	if err = c.Bind(d); err!= nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err} 
	}

	d.Id = xid.New().String()	
	d.Created  = time.Now()
	d.Updated = d.Created

	if d.Name == "" || d.Id == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid field"}		
	}

	resultOne,err := h.DB.Collection("device").InsertOne(context.Background(),d)
	
	err = h.DB.Collection("device").FindOne(
			context.Background(),
			bson.D{{"id",resultOne.InsertedID}},
	).Decode(&d)
	
	
	return c.JSON(http.StatusCreated,d)
}


func (h *Handler) ReadDevices(c echo.Context) (err error){
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
	
	devices := []*model.Device{}

	cur, err := h.DB.Collection("device").Find(
		context.Background(),bson.D{},findOptions)

	if err != nil {
			log.Fatal(err)
		}
	
	for cur.Next(context.TODO()) {

		var elem model.Device
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		devices = append(devices,&elem)
	}
	return c.JSON(http.StatusOK,devices)
	
}

func (h *Handler) ReadDevice(c echo.Context) (err error) {
	id := c.Param("id")
	
	d := &model.Device{}

	filter := bson.D{{"id", id}}

	err = h.DB.Collection("device").FindOne(
		context.Background(),
		filter,).Decode(&d)
	if err != nil {
		log.Fatal(err)
		// TODO:Error response by case
	}
	
	return c.JSON(http.StatusOK,d)
}

func (h *Handler) UpdateDevice(c echo.Context) (err error){

	
	id := c.Param("id")
	bd := &model.Device{}
	ud := &model.Device{}
	if err = c.Bind(ud); err!= nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err} 
	}

	filter := bson.D{{"id", id}}

	err = h.DB.Collection("device").FindOne(
		context.Background(),
		filter,).Decode(&bd)
	if err != nil {
		log.Fatal(err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Don't Have Resource"} 
	}

	ud.Created = bd.Created
	ud.Updated = time.Now()
	update := bson.D{
		{"$set",ud},
	}

	_, err = h.DB.Collection("device").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	
	return c.JSON(http.StatusOK,ud)
}


func (h *Handler) DeleteDevice(c echo.Context) (err error){
	id := c.Param("id")
	// TODO:Resource Role Check

	
	_, err = h.DB.Collection("device").DeleteOne(context.TODO(), bson.D{{"id",id}})
	if err != nil {
		return err
	}
 	return c.JSON(http.StatusOK,id)
}




