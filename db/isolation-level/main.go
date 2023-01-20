package main

import (
	"context"
	"isolevel/config"
	"isolevel/model"
	"isolevel/repository/mongo"
	"isolevel/repository/mysql"
	"time"
)

func main() {
	// Init MySQL
	db := config.InitMySQL()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}()

	// Init Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// ctx := context.Background()

	client := config.InitMongo(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// migration
	err = db.AutoMigrate(&model.MessageOne{})
	if err != nil {
		panic(err)
	}

	mysqlRepository := mysql.New(db)
	mongoRepository := mongo.New(client, ctx, config.MongoDatabaseName, config.MongoCollectionName)

	dummyFactory := config.NewDummyFactory(mysqlRepository, mongoRepository)
	//dummyFactory.CreateMySQLDummy(100000)
	dummyFactory.DeleteMySQLDummy()
	//dummyFactory.CreateMongoDummy(100000)
	//dummyFactory.DeleteMongoDummy()
	return
}
