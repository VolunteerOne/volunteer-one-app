package repository

import (
	"errors"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"gorm.io/gorm"
)

type FriendRepository interface {
	CreateFriend(friend models.Friend) (models.Friend, error)
	AcceptFriend(friend models.Friend) (models.Friend, error)
	RejectFriend(friend models.Friend) error
	OneFriend(id string) (models.Friend, error)
	GetFriends() ([]models.Friend, error)
}

type friendRepository struct {
	DB *gorm.DB
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return friendRepository{
		DB: db,
	}
}

func (f friendRepository) CreateFriend(friend models.Friend) (models.Friend, error) {

	err := f.DB.Create(&friend).Error

	return friend, err
}

func (f friendRepository) AcceptFriend(friend models.Friend) (models.Friend, error) {
	result := f.DB.Save(&friend)

	if result.Error != nil {
		return models.Friend{}, errors.New("could not update friend")
	}

	return friend, nil
}

func (f friendRepository) RejectFriend(friend models.Friend) error {
	// more explicit for testing -> f.DB.Delete(&friend) would do the same but it trickier to test
	result := f.DB.Where("ID = ?", friend.ID).Delete(&friend)
	if result.Error != nil {
		return errors.New("could not delete friend")
	}

	return nil
}

func (f friendRepository) OneFriend(id string) (models.Friend, error) {
	var friend models.Friend
	result := f.DB.First(&friend, id)

	if result.Error != nil {
		return models.Friend{}, errors.New("could not retrieve friend")
	}

	return friend, nil
}

func (f friendRepository) GetFriends() ([]models.Friend, error) {
	var friends []models.Friend
	result := f.DB.Find(&friends)

	if result.Error != nil {
		return []models.Friend{}, errors.New("could not retrieve friends")
	}

	return friends, nil
}
