package mysql

import (
	"isolevel/model"
	"log"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type Repository struct {
	DB   *gorm.DB
	Node *snowflake.Node
}

func New(db *gorm.DB) Repository {
	node, _ := snowflake.NewNode(1)
	return Repository{
		DB:   db,
		Node: node,
	}
}

func (r Repository) CreateUniqueID() string {
	return r.Node.Generate().String()
}

func (r Repository) Save(message model.MessageOne) (err error) {
	err = r.DB.Model(&model.MessageOne{}).Create(&message).Error
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (r Repository) FindBySubjectLimit100(subject string) (messages []model.MessageOne, err error) {
	err = r.DB.Limit(100).Order("created_at desc").Find(&messages, model.MessageOne{Subject: subject}).Error
	return
}

func (r Repository) SaveAll(messages []model.MessageOne) error {
	err := r.DB.Model(&model.MessageOne{}).CreateInBatches(&messages, 100).Error
	return err
}

func (r Repository) DeleteAll() error {
	err := r.DB.Unscoped().Where("1=1").Delete(&model.MessageOne{}).Error
	return err
}
