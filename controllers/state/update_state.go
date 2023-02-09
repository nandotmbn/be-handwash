package controller_state

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"tutorial/configs"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

var handwashCollection = configs.GetCollection(configs.DB, "handwash")

func UpdateState() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var statePayload views.PayloadState
		defer cancel()
		var handwashId = c.Param("handwash_id")
		c.BindJSON(&statePayload)

		if validationErr := validate.Struct(&statePayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		handwashIdObj, err := primitive.ObjectIDFromHex(handwashId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		count, err_ := handwashCollection.CountDocuments(ctx, bson.M{"_id": handwashIdObj})
		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count == 0 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Handwash by given Id is not found"})
			return
		}

		update := bson.M{
			"state":      statePayload.State,
			"battery":    statePayload.Battery,
			"updated_at": time.Now(),
		}
		result, err := handwashCollection.UpdateOne(ctx, bson.M{"_id": handwashIdObj}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": err.Error()})
			return
		}

		fmt.Println(result)

		finalView := views.FinalState{
			State:     statePayload.State,
			Battery:   statePayload.Battery,
			UpdatedAt: time.Now(),
		}

		json_data, err__ := json.Marshal(finalView)
		if err__ != nil {
			log.Fatal("JSON Marshalling Error: Update State")
		}

		http.Post("https://gdsc-pens-iot-listener-lxz6xwlfka-et.a.run.app/handwash/"+handwashId, "application/json", bytes.NewBuffer(json_data))

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})

	}
}
