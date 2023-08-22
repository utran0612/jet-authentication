package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt-project/database"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshCalims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(200)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic("Error creating token:", err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshCalims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic("Error creating refersh token", err)
		return
	}

	return token, refreshToken, err
}

func UpdateAllTokens(signedToken, signedRefreshToken, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateObj := bson.D{{"token", signedToken}, {"refreshToken", signedRefreshToken}, {"updated_at", time.Now().Format(time.RFC3339)}}

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	update := bson.D{{"$set", updateObj}}

	_, err := userCollection.UpdateOne(ctx, filter, update, &opt)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
