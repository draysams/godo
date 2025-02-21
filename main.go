package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {

	fmt.Println("Hello, Worlds!")

	err := godotenv.Load(".env.development.local") // env file for local development
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	MONGO_CONNECTION_STRING := os.Getenv("MONGO_CONNECTION_STRING")

	clientOptions := options.Client().ApplyURI(MONGO_CONNECTION_STRING)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(nil)
	}

	fmt.Println("Connected to MongoDB successfully!")

	collection = client.Database("godo_db").Collection("todos")

	app := fiber.New()

	app.Use((cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
	})))
	app.Get("/api/todos", getTodos)
	app.Get("/api/todos/:id", getTodo)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen(":" + port))

}

func getTodo(c *fiber.Ctx) error {
	fmt.Println("Check read one!")
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var todo Todo

	filter := bson.M{"_id": objectID}

	err = collection.FindOne(context.Background(), filter).Decode(&todo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func getTodos(c *fiber.Ctx) error {
	fmt.Println("Check read all!")

	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		todos = append(todos, todo)
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	fmt.Println("Check create!")
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if todo.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body is required"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	fmt.Println("Check patch!")
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo updated successfully"})

}

func deleteTodo(c *fiber.Ctx) error {
	fmt.Println("Check delete!")
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Todo deleted successfully"})
}
