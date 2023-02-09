package controller_handwash

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetIdHandwash() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var handwashPayload views.PayloadRetriveId
		defer cancel()
		c.BindJSON(&handwashPayload)

		if validationErr := validate.Struct(&handwashPayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		var resultMeter models.Handwash
		var finalPayload views.FinalRetriveId
		result := handwashCollection.FindOne(ctx, bson.M{"handwash_name": handwashPayload.HandwashName})
		result.Decode(&resultMeter)
		result.Decode(&finalPayload)
		err := bcrypt.CompareHashAndPassword([]byte(resultMeter.Password), []byte(handwashPayload.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
				"data":    "Handwash name or password is not valid",
			})
			return
		}

		c.JSON(http.StatusOK,
			bson.M{
				"status":  http.StatusOK,
				"message": "Success",
				"data":    finalPayload,
			},
		)
	}
}
