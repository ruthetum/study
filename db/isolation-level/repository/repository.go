package repository

import "isolevel/model"

type Repository interface {
	CreateUniqueID() string
	Save(message model.Message) (err error)
	FindBySubjectLimit100(subject string) (messages []model.Message, err error)
	SaveAll(messages []model.Message) error
	DeleteAll() error
}
