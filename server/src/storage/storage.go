package storage

import "github.com/jinzhu/gorm"

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
		Title:          title,
		UserID:         userID,
		BaseModel:      *NewBaseModel(),
		NextStopNumber: 1,
	}

	return &learningPath, store.db.Create(&learningPath).Error
}

func (store *MySqlStore) GetLearningPathByID(userID uint, learningPathID uint) (*LearningPath, error) {
	learningPath := LearningPath{}

	return &learningPath, store.db.Where("userID = ? AND learningPathID = ?", userID, learningPathID).First(&learningPath).Error
}

func (store *MySqlStore) GetLearningPathStopByID(userID uint, learningPathStopID uint) (*LearningPathStop, error) {
	stop := LearningPathStop{}

	return &stop, store.db.Where("userID = ? AND learningPathStopID = ?", userID, learningPathStopID).First(&stop).Error
}

func (store *MySqlStore) AddStopToLearningPath(path *LearningPath, stop *LearningPathStop) error {
	var assoc *gorm.Association = store.db.Model(&path).Association("Stops").Append(&stop)
	return assoc.Error
}

func (store *MySqlStore) IncrementStopCount(path *LearningPath) error {
	return store.db.Model(path).Update("nextStopNumber", path.NextStopNumber+1).Error
}

func (store *MySqlStore) SetLearningPathTitle(path *LearningPath, newTitle string) error {
	return store.db.Model(path).Update("Title", newTitle).Error
}

func (store *MySqlStore) SetLearningPathStopTitle(stop *LearningPathStop, newTitle string) error {
	return store.db.Model(stop).Update("Title", newTitle).Error
}

func (store *MySqlStore) SetLearningPathStopBody(stop *LearningPathStop, body string) error {
	return store.db.Model(stop).Update("MarkdownBody", body).Error
}
