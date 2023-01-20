package main

import (
	"context"
	"fmt"
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
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 60)

	defer func() {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}()

	// Init Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
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

	factory := config.NewDummyFactory(mysqlRepository, mongoRepository)
	factory.DeleteMySQLDummy()
	factory.DeleteMongoDummy()

	sizes := []int{1000, 5000, 10000, 50000}
	dbTypes := []bool{false}

	for _, dbType := range dbTypes {
		for _, size := range sizes {
			coverage(factory, size, dbType)
			time.Sleep(time.Second * 1)
		}
	}
	return
}

func coverage(factory config.DummyFactory, size int, isMySQL bool) {
	if isMySQL {
		fmt.Printf("[MySQL] Size:%d\n", size)
		// size 만큼 쓰기
		factory.CreateMySQLDummy(size)

		// size 읽기
		factory.FindMySQLDummy()
		time.Sleep(time.Second * 2)

		// 전체 삭제
		factory.DeleteMySQLDummy()
		time.Sleep(time.Second * 2)

		// size 만큼 쓰기/읽기 반복
		factory.CreateAndFindMySQLDummy(size)
		factory.DeleteMySQLDummy()
		fmt.Println()
		time.Sleep(time.Second * 2)
		return
	}

	fmt.Printf("[Mongo] Size:%d\n", size)
	// size 만큼 쓰기
	factory.CreateMongoDummy(size)

	// size 읽기
	factory.FindMongoDummy()

	// 전체 삭제
	factory.DeleteMongoDummy()
	time.Sleep(time.Second * 2)

	// size 만큼 쓰기/읽기 반복
	factory.CreateAndFindMongoDummy(size)
	factory.DeleteMongoDummy()
	fmt.Println()
	time.Sleep(time.Second & 2)
	return
}
