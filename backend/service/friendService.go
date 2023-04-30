package service

import (
	"github.com/gin-gonic/gin"
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/models"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
)

type FriendService interface {
	CreateFriend(friend models.Friend) (models.Friend, error)
	AcceptFriend(friend models.Friend) (models.Friend, error)
	RejectFriend(friend models.Friend) error
	OneFriend(id string) (models.Friend, error)
	GetFriends() ([]models.Friend, error)
	Bind(*gin.Context, any) error
}

type friendService struct {
	friendRepository repository.FriendRepository
}

func NewFriendService(r repository.FriendRepository) FriendService {
	return friendService{
		friendRepository: r,
	}
}

func (f friendService) CreateFriend(friend models.Friend) (models.Friend, error) {
	log.Println("[FriendService] Create friend request...")
	return f.friendRepository.CreateFriend(friend)
}

func (f friendService) AcceptFriend(friend models.Friend) (models.Friend, error) {
	log.Println("[FriendService] Accept friend...")
	return f.friendRepository.AcceptFriend(friend)
}

func (f friendService) RejectFriend(friend models.Friend) error {
	log.Println("[FriendService] Reject friend...")
	return f.friendRepository.RejectFriend(friend)
}

func (f friendService) OneFriend(id string) (models.Friend, error) {
	log.Println("[FriendService] Accept friend...")
	return f.friendRepository.OneFriend(id)
}

func (f friendService) GetFriends() ([]models.Friend, error) {
	log.Println("[FriendService] Get all friends...")
	return f.friendRepository.GetFriends()
}

func (f friendService) Bind(c *gin.Context, obj any) error {
	return c.Bind(obj)
}
