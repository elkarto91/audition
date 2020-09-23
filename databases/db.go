package databases

import (
	"errors"
	"fmt"
	models "github.com/elkarto91/audition/common"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

var (
	ErrInvalidCredentials = errors.New("admin credentials dont match")
)

const (
	Database       = "auditionDb"
	UserCollection = "userCollection"
)

func getMongoSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return session, nil
}
func InitMongo() error {
	session, err := getMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(Database).C(UserCollection)

	// Index
	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return fmt.Errorf("error initialing indexes db %v", err)
	}
	return nil
}

func RegisterUser(user *models.User) error {

	if user.Username != "" && user.Password != "" {
		cost := bcrypt.DefaultCost
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
		if err != nil {
			return err
		}
		user.Password = string(hash)
		return AddUser(user)
	}
	return ErrInvalidCredentials
}
func AddUser(user *models.User) error {
	session, err := getMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(Database).C(UserCollection)
	err = c.Insert(user)
	if err != nil {
		return fmt.Errorf("error inserting user to db %v", err)
	}
	return nil
}

func UpdateUser(user *models.User) error {
	session, err := getMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(Database).C(UserCollection)
	colQueried := bson.M{"username": user.Username}
	err = c.Update(colQueried, user)
	if err != nil {
		return fmt.Errorf("error updating user to db %v", err)
	}
	return nil
}

func ListAllUsers() ([]*models.User, error) {
	session, err := getMongoSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	c := session.DB(Database).C(UserCollection)
	var users []*models.User
	err = c.Find(nil).Select(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	if exists, err := DoesUserExist(username); exists && err == nil {
		session, err := getMongoSession()
		if err != nil {
			return nil, err
		}
		defer session.Close()
		c := session.DB(Database).C(UserCollection)
		user := &models.User{}
		err = c.Find(bson.M{"username": username}).One(user)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else if err != nil {
		return nil, err
	} else {
		return nil, fmt.Errorf("no such user found with username = [%v] ", username)
	}
}

func DoesUserExist(username string) (bool, error) {
	session, err := getMongoSession()
	if err != nil {
		return false, err
	}
	defer session.Close()
	c := session.DB(Database).C(UserCollection)
	i, err := c.Find(bson.M{"username": username}).Count()
	if err != nil {
		return false, err
	}
	return i > 0, nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return user, nil
}
