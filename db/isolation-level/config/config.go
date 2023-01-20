package config

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MongoDatabaseName   = "test"
	MongoCollectionName = "message"
)

func InitMySQL() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:13306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func InitMongo(ctx context.Context) *mongo.Client {
	// mongodb+srv://<username>:<password>@cluster0-zzart.mongodb.net/test?retryWrites=true&w=majority
	uri := "mongodb://root:1234@localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMinPoolSize(30)
	clientOptions.SetMaxPoolSize(30)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	indexModels := []mongo.IndexModel{
		{
			Keys: bson.M{
				"message_id": -1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"subject": 1,
			},
			Options: options.Index().SetUnique(false),
		},
	}
	_, err = client.Database(MongoDatabaseName).Collection(MongoCollectionName).Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		panic(err)
	}
	return client
}
