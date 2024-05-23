package models

type User struct {
	UserID   string `gorm:"column:user_id;primaryKey"`
	Password string
}

type Level struct {
	UserID string `gorm:"column:user_id;primaryKey"`
	Level  int
}

func (db *Database) CreateUser(userId string, password string) error {
	user := User{UserID: userId, Password: password}
	result := db.Gorm.Create(user)
	return result.Error
}

func (db *Database) GetUser(userId string) (User, error) {
	user := &User{}
	result := db.Gorm.First(user, "user_id = ?", userId)
	return *user, result.Error
}
