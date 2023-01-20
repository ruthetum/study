package config

import (
	"fmt"
	"isolevel/model"
	"isolevel/repository/mongo"
	"isolevel/repository/mysql"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const typeList = 4

type DummyFactory struct {
	mysqlRepository mysql.Repository
	mongoRepository mongo.Repository
}

func NewDummyFactory(mysqlRepository mysql.Repository, mongoRepository mongo.Repository) DummyFactory {
	return DummyFactory{
		mysqlRepository: mysqlRepository,
		mongoRepository: mongoRepository,
	}
}

func (f DummyFactory) CreateMySQLDummy(size int) {
	defer timer("CreateMySQLDummy")()
	messages := make([]model.MessageOne, 0)
	for i := 0; i < size; i++ {
		messageID := f.mysqlRepository.CreateUniqueID()
		message := model.MessageOne{}
		content := make(map[string]any)
		content["to"] = rand.Intn(100) + 1000
		content["from"] = 1
		message.Create(messageID, "test type", strconv.Itoa(size%typeList), content)
		messages = append(messages, message)
	}
	err := f.mysqlRepository.SaveAll(messages)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (f DummyFactory) DeleteMySQLDummy() {
	defer timer("DeleteMySQLDummy")()
	err := f.mysqlRepository.DeleteAll()
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) CreateMongoDummy(size int) {
	defer timer("CreateMongoDummy")()
	messages := make([]model.MessageTwo, 0)
	for i := 0; i < size; i++ {
		messageID := f.mongoRepository.CreateUniqueID()
		message := model.MessageTwo{}
		content := make(map[string]any)
		content["to"] = rand.Intn(100) + 1000
		content["from"] = 1
		message.Create(messageID, "test type", strconv.Itoa(size%typeList), content)
		messages = append(messages, message)
	}
	err := f.mongoRepository.SaveAll(messages)
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) DeleteMongoDummy() {
	defer timer("DeleteMongoDummy")()
	err := f.mongoRepository.DeleteAll()
	if err != nil {
		log.Fatalln(err)
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() { fmt.Println(name, ":", time.Since(start)) }
}
