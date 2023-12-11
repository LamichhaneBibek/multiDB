package repo

import (
	"github.com/LamichhaneBibek/multiDB/database"
	"github.com/LamichhaneBibek/multiDB/model"
)

type UserRepo interface {
	Save(...model.User) error
	FindAll() ([]model.User, error)
}

type UserRepoImpl struct {
	db DBOperator
}

func NewUserRepo(dbConIdentificationName string) UserRepo {
	return UserRepoImpl{
		db: DBOperator{
			Con: database.GetDBConnection(dbConIdentificationName),
		},
	}
}

func (repo UserRepoImpl) Save(user ...model.User) error {
	err := repo.db.Save(user)
	if err != nil {
		// handle any special logic hdto
		return err
	}
	return nil
}

func (repo UserRepoImpl) FindAll() ([]model.User, error) {
	var users []model.User
	err := repo.db.Find(&users)
	if err != nil {
		// handle any special logic here
		return nil, err
	}

	return users, nil
}
