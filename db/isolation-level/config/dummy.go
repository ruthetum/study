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

const (
	typeList          = 4
	maxGoroutineCount = 101
)

type DummyFactory struct {
	MysqlRepository mysql.Repository
	MongoRepository mongo.Repository
}

func NewDummyFactory(mysqlRepository mysql.Repository, mongoRepository mongo.Repository) DummyFactory {
	return DummyFactory{
		MysqlRepository: mysqlRepository,
		MongoRepository: mongoRepository,
	}
}

func (f DummyFactory) CreateMySQLDummy(size int) {
	defer timer("CreateMySQLDummy")()
	messages := make([]model.MessageOne, 0)
	for i := 0; i < size; i++ {
		messageID := f.MysqlRepository.CreateUniqueID()
		message := model.MessageOne{}
		content := make(map[string]any)
		content["to"] = rand.Intn(100) + 1000
		content["from"] = 1
		message.Create(messageID, "test type", strconv.Itoa(size%typeList), content)
		messages = append(messages, message)
	}
	err := f.MysqlRepository.SaveAll(messages)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (f DummyFactory) FindMySQLDummy() {
	defer timer("FindMySQLDummy")()
	_, err := f.MysqlRepository.FindBySubjectLimit100("0")
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) DeleteMySQLDummy() {
	err := f.MysqlRepository.DeleteAll()
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) CreateAndFindMySQLDummy(size int) {
	defer timer("CreateAndFindMySQLDummy")()

	g := make(chan struct{}, maxGoroutineCount)
	for i := 0; i < size*2; i++ {
		g <- struct{}{}
		if i%2 == 0 {
			messageID := f.MysqlRepository.CreateUniqueID()
			message := model.MessageOne{}
			content := make(map[string]any)
			content["to"] = rand.Intn(100) + 1000
			content["from"] = 1
			message.Create(messageID, "test type", strconv.Itoa(size%typeList), content)
			go func() {
				err := f.MysqlRepository.Save(message)
				if err != nil {
					log.Fatalln(err)
				}
				<-g
			}()
		} else {
			go func() {
				_, err := f.MysqlRepository.FindBySubjectLimit100("0")
				if err != nil {
					log.Fatalln(err)
				}
				<-g
			}()
		}
	}
	return
}

func (f DummyFactory) CreateMongoDummy(size int) {
	defer timer("CreateMongoDummy")()
	messages := make([]model.MessageTwo, 0)
	for i := 0; i < size; i++ {
		messageID := f.MongoRepository.CreateUniqueID()
		message := model.MessageTwo{}
		content := make(map[string]any)
		content["to"] = rand.Intn(100) + 1000
		content["from"] = 1
		message.Create(messageID, "test type", strconv.Itoa(size%typeList), content)
		messages = append(messages, message)
	}
	err := f.MongoRepository.SaveAll(messages)
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) FindMongoDummy() {
	defer timer("FindMongoDummy")()
	_, err := f.MongoRepository.FindBySubjectLimit100("0")
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) DeleteMongoDummy() {
	err := f.MongoRepository.DeleteAll()
	if err != nil {
		log.Fatalln(err)
	}
}

func (f DummyFactory) CreateAndFindMongoDummy(size int) {
	defer timer("CreateAndFindMongoDummy")()

	g := make(chan struct{}, maxGoroutineCount)
	for i := 0; i < size*2; i++ {
		g <- struct{}{}
		if i%2 == 0 {
			go func() {
				messageID := f.MongoRepository.CreateUniqueID()
				message := model.MessageTwo{}
				content := make(map[string]any)
				content["to"] = rand.Intn(100) + 1000
				content["from"] = 1
				message.Create(messageID, "test type", strconv.Itoa(size%typeList), content)
				err := f.MongoRepository.Save(message)
				if err != nil {
					log.Fatalln(err)
				}
				<-g
			}()
		} else {
			go func() {
				_, err := f.MongoRepository.FindBySubjectLimit100("0")
				if err != nil {
					log.Fatalln(err)
				}
				<-g
			}()
		}
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() { fmt.Println(name, ":", time.Since(start)) }
}
