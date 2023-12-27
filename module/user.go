package module

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/indrariksa/be_presensi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	return randomString, nil
}

func GetUserByUsername(username string, db *mongo.Database, col string) (user model.User, err error) {
	usersCollection := db.Collection(col)
	filter := bson.M{"username": username}

	err = usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("user not found with username: %s", username)
		}
		return user, fmt.Errorf("error retrieving user: %s", err.Error())
	}

	return user, nil
}

func VerifyPassword(user model.User, providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Login(username, password string, db *mongo.Database, col string) (loggedIn bool, token string, err error) {
	user, err := GetUserByUsername(username, db, col)
	if err != nil {
		return false, "", err
	}

	passwordMatched, err := VerifyPassword(user, password)
	if err != nil {
		return false, "", err
	}

	if !passwordMatched {
		return false, "", fmt.Errorf("invalid password")
	}

	token, err = GenerateRandomString(26)
	if err != nil {
		return false, "", err
	}

	return true, token, nil
}

func CreateUser(db *mongo.Database, col string, username, password string) (insertedID primitive.ObjectID, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to hash the password: %s", err.Error())
	}

	user := model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	result, err := db.Collection(col).InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert user: %s", err.Error())
	}

	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
