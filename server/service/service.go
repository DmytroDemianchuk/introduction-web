package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Easy-Job-Developer/catalog_plus/domain"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// ------------------------------------------------- DATABASE  -------------------------------------------------

const connectionString = "mongodb+srv://Dmytro:moqi7e123@mongodb.nuim2dr.mongodb.net/?retryWrites=true&w=majority&appName=MongoDB"
const dbName = "people"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

// -------------------------------------------------

func InsertOneName(user domain.User) {
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 name in db with id:", inserted.InsertedID)
}

func DeleteAllName() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func GetAllNames() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var names []primitive.M
	for cur.Next(context.Background()) {
		var name bson.M
		err := cur.Decode(&name)
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	defer cur.Close(context.Background())
	return names
}

func CreateUser(user domain.User) error {
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

const jwtSecret = "your_jwt_secret"

func SignIn(email, password string) (string, error) {
	// Find the user by email
	var user domain.User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", domain.ErrUserNotFound
		}
		log.Println("Error finding user:", err)
		return "", err
	}

	// Compare the provided password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", domain.ErrUserNotFound // Passwords don't match
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with a secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error generating JWT token:", err)
		return "", err
	}

	return tokenString, nil
}
