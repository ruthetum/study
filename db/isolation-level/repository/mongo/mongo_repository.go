package mongo

import (
	"context"
	"isolevel/model"
	"log"

	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Client         *mongo.Client
	Ctx            context.Context
	DatabaseName   string
	CollectionName string
	Node           *snowflake.Node
}

func New(client *mongo.Client, ctx context.Context, database, collection string) Repository {
	node, _ := snowflake.NewNode(2)
	return Repository{
		Client:         client,
		Ctx:            ctx,
		DatabaseName:   database,
		CollectionName: collection,
		Node:           node,
	}
}

func (r Repository) CreateUniqueID() string {
	return r.Node.Generate().String()
}

func (r Repository) Save(message model.MessageTwo) (err error) {
	_, err = r.collection().InsertOne(r.Ctx, message)
	return
}

func (r Repository) FindBySubjectLimit100(subject string) (messages []model.MessageTwo, err error) {
	const DescendingOptions = -1
	opts := options.Find().SetSort(bson.D{{"created_at", DescendingOptions}}).SetLimit(int64(100))

	filter := bson.M{"subject": subject}

	cursor, err := r.collection().Find(r.Ctx, filter, opts)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer cursor.Close(r.Ctx)

	for cursor.Next(r.Ctx) {
		var m model.MessageTwo
		if err = cursor.Decode(&m); err != nil {
			log.Fatalln(err)
			return
		}
		messages = append(messages, m)
	}
	return

}

func (r Repository) SaveAll(messages []model.MessageTwo) error {
	m := make([]interface{}, 0)
	for _, message := range messages {
		var i interface{}
		i = message
		m = append(m, i)
	}
	_, err := r.collection().InsertMany(r.Ctx, m)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (r Repository) DeleteAll() error {
	_, err := r.collection().DeleteMany(r.Ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (r Repository) collection() *mongo.Collection {
	return r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
}
