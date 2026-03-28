package main

import (
	"context"
	"log"
	"time"

	"hexagonal2/adapter/handler"
	"hexagonal2/adapter/repository"
	"hexagonal2/core/ports"
	"hexagonal2/core/service"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {


	userRepo, dogRepo, err := connectDatabase(true) // pass false to attempt MongoDB connection
	if err != nil {
		log.Fatal("failed to connect to any database:", err)
	}

	// services
	userSrv := service.NewUserService(userRepo)
	dogSrv := service.NewDogService(dogRepo)

	// handlers
	dh := handler.NewDogHandler(dogSrv)
	hh := handler.NewUserHandler(userSrv)


	app := fiber.New()

	app.Get("/dogs", dh.GetAllDogs)
	app.Get("/dogs/:id", dh.GetADogs)
	app.Post("/dogs", dh.AddDog)

	app.Get("/users", hh.GetAllUsers)
	app.Get("/users/:id", hh.GetAUser)
	app.Post("/users", hh.AddUser)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}


func connectDatabase(flag bool) (userRepo ports.UserRepository, dogRepo ports.DogsRepository,err error) {
	if flag {
		db, err := gorm.Open(sqlite.Open("hgo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&repository.UserDB{}, &repository.DogsModel{}); err != nil {
		log.Fatal(err)
	}

	// repositories (default to GORM sqlite implementations)
	userRepo = repository.NewUserRepositoryDB(db)
	dogRepo = repository.NewDogsRepositoryDB(db)

	return userRepo, dogRepo, nil


	}else{
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err == nil {
		if err = mongoClient.Ping(ctx, nil); err == nil {
			log.Println("connected to mongodb, switching repositories")
			userRepo = repository.NewUserRepositoryMongo(mongoClient, "hgo")
			dogRepo = repository.NewDogsRepositoryMongo(mongoClient, "hgo")
			return userRepo, dogRepo, nil
		} else {
			log.Println("mongo ping failed:", err)
		}
	} else {
		log.Println("mongo connect failed:", err)
	}

	}
	
	// try connecting to MongoDB and swap repositories if available

	return nil, nil, err
}