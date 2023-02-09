package controller_handwash

import (
	"context"
	"net/http"
	"time"
	"tutorial/configs"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

var handwashCollection *mongo.Collection = configs.GetCollection(configs.DB, "handwash")

func RegisterHandwash() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var handwash views.PayloadRetriveId
		defer cancel()
		c.BindJSON(&handwash)

		if validationErr := validate.Struct(&handwash); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		count, err_ := handwashCollection.CountDocuments(ctx, bson.M{"handwash_name": handwash.HandwashName})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count >= 1 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Meter name has been taken"})
			return
		}

		bytes, errors := bcrypt.GenerateFromPassword([]byte(handwash.Password), 14)
		if errors != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Password tidak valid"})
		}

		newHandwash := models.Handwash{
			HandwashName: handwash.HandwashName,
			State:        false,
			Battery:      0,
			Password:     string(bytes),
			UpdatedAt:    time.Now(),
			CreatedAt:    time.Now(),
		}

		result, err := handwashCollection.InsertOne(ctx, newHandwash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		finalView := views.HandwashView{
			HandwashId:   result.InsertedID,
			HandwashName: handwash.HandwashName,
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})
	}
}
