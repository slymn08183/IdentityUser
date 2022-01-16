package controller

import (
	"IdentityUser/constant"
	"IdentityUser/dal"
	"IdentityUser/entity"
	"IdentityUser/helper"
	"IdentityUser/model"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"time"
)

var userCollection = dal.GetUserCollection()
var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user entity.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, model.Error{Message: err.Error()}.GetAsEnvelope())
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			errStr := validationErr.Error()
			if strings.Contains(errStr, constant.MongoMailValidationErr) {
				c.JSON(http.StatusBadRequest, model.Error{Message: constant.EmailValidationError}.GetAsEnvelope())
				return
			}
			c.JSON(http.StatusBadRequest, model.Error{Message: errStr}.GetAsEnvelope())
			return
		}

		/*	_, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

			if err != nil {
				c.JSON(http.StatusInternalServerError, model.Error{Message: constant.InternalServerError}.GetAsEnvelope())
				log.Panic(err)
				return
			}*/

		password := helper.HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserId)
		user.Token = &token
		user.RefreshToken = &refreshToken

		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			errStr := insertErr.Error()

			// Duplicate Errors
			if strings.Contains(errStr, constant.MongoDuplicateKeyErrCodeStr) {
				if strings.Contains(errStr, constant.MongoDuplicateMailErr) {
					c.JSON(http.StatusBadRequest, model.Error{Message: constant.DuplicateMailErr}.GetAsEnvelope())
					return
				}
			}

			c.JSON(http.StatusInternalServerError, model.Error{Message: constant.UserDidNotCreatedErr}.GetAsEnvelope())
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, model.Success{}.True())
		return
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user entity.User
		var foundUser entity.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, model.Error{Message: err.Error()}.GetAsEnvelope())
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Error{Message: err.Error()}.GetAsEnvelope())
			return
		}

		passwordIsValid, _ := helper.VerifyPassword(*user.Password, *foundUser.Password)

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, model.Error{Message: constant.PasswordOrMailValidationError}.GetAsEnvelope())
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)

		helper.UpdateAllTokens(token, refreshToken, foundUser.UserId)

		c.JSON(http.StatusOK, model.LoginResponse{Token: foundUser.Token, RefreshToken: foundUser.RefreshToken})

	}
}
