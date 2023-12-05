package database

import (
	"fmt"

	"github.com/sharin-sushi/0016go_next_relation/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) CreateUser(user domain.Listener) (domain.Listener, error) {
	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *UserRepository) LogicalDeleteUser(user domain.Listener) error {
	err := db.Delete(&user, &user.ListenerId).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *UserRepository) FindUserByEmail(email string) (domain.Listener, error) {
	var user domain.Listener
	query := "email = '" + email + "'"
	err := db.Where(query).First(&user).Error
	fmt.Printf("user=%v \n", user)
	if err != nil {
		return user, err
	}
	return user, err
}
func (db *UserRepository) LogIn(user domain.Listener) (domain.Listener, error) {
	if err := db.First(&user, user.ListenerId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *UserRepository) FindUserByListenerId(ListenerId domain.ListenerId) (domain.Listener, error) {
	var user domain.Listener
	query := fmt.Sprintf("listener_id = %v", ListenerId)
	err := db.Where(query).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

// FindUserByListenerIdと同じ処理
// func (db *UserRepository) GetListenerProfile(user domain.Listener) (foundUser domain.Listener, err error) {
// 	var user domain.Listener
// 	err := db.First
// }

// func (db *UserRepository) GuestLogIn() (u domain.Listener, err error) {
// 	var user domain.Listener
// 	if err := db.First(&user, user.ListenerName).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }
