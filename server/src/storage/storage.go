package storage

import "github.com/jinzhu/gorm"

type Store interface {
	Connect() error
	GetUserByEmail(email string, user *User) *gorm.DB
	CreateUser(email string, passwordHash string) (*User, error)
	CreateLearningPath(title string, userID uint) (*LearningPath, error)
	GetLearningPathByID(userID uint, learningPathID uint) (*LearningPath, error)
	AddStopToLearningPath(path *LearningPath, stop *LearningPathStop) error
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

func (store *MySqlStore) CreateUser(email string, passwordHash string) (*User, error) {
	user := User{
		Email:     email,
		Hash:      passwordHash,
		BaseModel: *NewBaseModel(),
	}

	return &user, store.db.Create(&user).Error
}

func (store *MySqlStore) CreateLearningPath(title string, userID uint) (*LearningPath, error) {

	learningPath := LearningPath{
		Title:     title,
		UserID:    userID,
		BaseModel: *NewBaseModel(),
	}

	return &learningPath, store.db.Create(&learningPath).Error
}

func (store *MySqlStore) GetLearningPathByID(userID uint, learningPathID uint) (*LearningPath, error) {
	learningPath := LearningPath{}

	return &learningPath, store.db.Where("userID = ? AND learningPathID = ?", userID, learningPathID).First(&learningPath).Error
}

func (store *MySqlStore) AddStopToLearningPath(path *LearningPath, stop *LearningPathStop) error {
	var assoc *gorm.Association = store.db.Model(&path).Association("Stops").Append(&stop)
	return assoc.Error
}
