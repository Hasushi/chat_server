package database

import (
	"chat_server/adapter/database/model"
	"chat_server/domain/entity"
	"chat_server/usecase/output_port"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) output_port.User {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindByID(userID string) (entity.User, error) {
	return findByID(u.db, userID)
}

func (u *UserRepository) FindByIDWithTx(tx interface{}, userID string) (entity.User, error) {
	db, ok := tx.(*gorm.DB)
	if !ok {
		return entity.User{}, output_port.ErrInvalidTransaction
	}

	return findByID(db, userID)
}

func findByID(db *gorm.DB, userID string) (entity.User, error) {
	var model model.User
	err := db.Model(&model).Where("user_id = ?", userID).First(&model).Error
	if err != nil {
		return entity.User{}, err
	}

	return model.ToEntity(), nil
}

func (u *UserRepository) Create(args output_port.CreateUserArgs) error {
	return createUser(u.db, args)
}

func (u *UserRepository) CreateWithTx(tx interface{}, args output_port.CreateUserArgs) error {
	db, ok := tx.(*gorm.DB)
	if !ok {
		return output_port.ErrInvalidTransaction
	}

	return createUser(db, args)
}

func createUser(db *gorm.DB, args output_port.CreateUserArgs) error {
	model := model.User{
		UserID: args.UserID,
		UserName: args.UserName,
		Email: args.Email,
		HashedPassword: args.HashedPassword,
		Bio: "",
		IconUrl: "",
	}

	err := db.Create(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByEmail(email string) (entity.User, error) {
	var model model.User
	err := u.db.Model(&model).Where("email = ?", email).First(&model).Error
	if err != nil {
		return entity.User{}, err
	}

	return model.ToEntity(), nil
}

func (u *UserRepository) Update(args output_port.UpdateUserArgs) error {
	return updateUser(u.db, args)
}

func (u *UserRepository) UpdateWithTx(tx interface{}, args output_port.UpdateUserArgs) error {
	db, ok := tx.(*gorm.DB)
	if !ok {
		return output_port.ErrInvalidTransaction
	}

	return updateUser(db, args)
}


func updateUser(db *gorm.DB, args output_port.UpdateUserArgs) error {
	model := model.User{
		UserID: args.UserID,
		Bio: args.Bio,
		IconUrl: args.IconUrl,
	}

	err := db.Model(&model).Where("user_id = ?", args.UserID).Updates(&model).Error
	if err != nil {
		return err
	}

	return nil
}