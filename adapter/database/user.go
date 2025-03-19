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
	var model model.User
	err := u.db.Model(&model).Where("user_id = ?", userID).First(&model).Error
	if err != nil {
		return entity.User{}, err
	}

	return model.ToEntity(), nil
}

func (u *UserRepository) Create(args output_port.CreateUserArgs) error {
	model := model.User{
		UserID: args.UserID,
		UserName: args.UserName,
		Email: args.Email,
		HashedPassword: args.HashedPassword,
		// 新規登録時はDisplayNameはUserNameと同じ
		DisplayName: args.UserName,
		IconUrl: "",
	}

	err := u.db.Create(&model).Error
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
	model := model.User{
		UserID: args.UserID,
		DisplayName: args.DisplayName,
		IconUrl: args.IconUrl,
	}

	err := u.db.Model(&model).Where("user_id = ?", args.UserID).Updates(&model).Error
	if err != nil {
		return err
	}

	return nil
}