package storage

import "github.com/jinzhu/gorm"

type Store interface {
	Connect() error
	GetUserByEmail(email string, user *User) *gorm.DB
	CreateUser(email string, passwordHash string) (*gorm.DB, *User)
	CreateLearningPath(title string, userID uint) (*gorm.DB, *LearningPath)
}

type MySqlStore struct {
	db *gorm.DB
}

func NewMySqlStore() *MySqlStore {
	return &MySqlStore{}
}

func (store *MySqlStore) GetUserByEmail(email string, user *User) *gorm.DB {
	return store.db.Where("email = ?", email).First(user)
}

func (store *MySqlStore) CreateUser(email string, passwordHash string) (*gorm.DB, *User) {
	user := User{
		Email:     email,
		Hash:      passwordHash,
		BaseModel: *NewBaseModel(),
	}

	return store.db.Create(&user), &user
}

func (store *MySqlStore) CreateLearningPath(title string, userID uint) (*gorm.DB, *LearningPath) {

	learningPath := LearningPath{
		Title:     title,
		UserID:    userID,
		BaseModel: *NewBaseModel(),
	}

	return store.db.Create(&learningPath), &learningPath
}
