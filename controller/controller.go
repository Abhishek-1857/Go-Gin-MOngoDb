package controllers

import (
	"context"
	"net/http"
	"oracleservice/config"
	"oracleservice/models"
	"oracleservice/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var KycInfoCollection *mongo.Collection = config.GetCollection(config.DB, "ipfsinfo")
var IpfsHashCollection *mongo.Collection = config.GetCollection(config.DB, "ipfshashinfo")

var validate = validator.New()

func AddInfo() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var info models.KycInfo
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&info); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&info); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        newUser := models.KycInfo{
            Id:       primitive.NewObjectID(),
            Name:     info.Name,
            Country:  info.Country,
        }
      
        result, err := KycInfoCollection.InsertOne(ctx, newUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}

func GetInfo() gin.HandlerFunc {
    return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var KycInfo []models.KycInfo
        defer cancel()

		cur,err:=KycInfoCollection.Find(ctx,bson.D{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,responses.Response{Status: http.StatusInternalServerError,Message: "error",Data:gin.H{"data":err.Error()}})
		}
		defer cur.Close(ctx)
        for cur.Next(ctx) {
            var info models.KycInfo
            if err = cur.Decode(&info); err != nil {
                c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
          
            KycInfo = append(KycInfo, info)
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": KycInfo}},
        )

}}


func GetIpfsHashes()  gin.HandlerFunc {
    return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var IpfsHash []models.IpfsHash
        defer cancel()

		cur,err:=IpfsHashCollection.Find(ctx,bson.D{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,responses.Response{Status: http.StatusInternalServerError,Message: "error",Data:gin.H{"data":err.Error()}})
		}
		defer cur.Close(ctx)
        for cur.Next(ctx) {
            var hashinfo models.IpfsHash
            if err = cur.Decode(&hashinfo); err != nil {
                c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
          
            IpfsHash = append(IpfsHash, hashinfo)
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": IpfsHash}},
        )		

	}}