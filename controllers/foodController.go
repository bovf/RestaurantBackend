package controller

import (
	"github.com/gin-gonic/gin"
	"context"
	"fmt"
)

var foodCollection *= mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
		foodId := c.Param("food_id")
		var food model.Food

		err := FoodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occurred while fetching food id"})
			}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
		var menu model.Menu
		var food model.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := MenuCollection.FindOne(ctx, bson.M{"menu_id": food.MenuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu id %s not found", food.MenuId)
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		food.CreatedAt = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.UpdatedAt = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()

		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := FoodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)	
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
