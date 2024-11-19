package models

type User struct {
    ID            int       `gorm:"primaryKey" json:"id"`
    Username string    `gorm:"size:50;not null" json:"username"`
    Password    string    `gorm:"size:100;not null" json:"password"`
    Credentials  []Credential `gorm:"foreignKey:UserID" json:"credentials"`
}

type Credential struct {
    ID           int       `gorm:"primaryKey" json:"id"`
    Title       string    `gorm:"size:100;not null" json:"title"`
    Credential_User string   `gorm:"size:50;not null" json:"credential_user"`
    Credential_Password    string   `gorm:"size:100;not null" json:"credential_password"`
    Description   string   `gorm:"size:100;not null" json:"description"`
    UserID       int       `gorm:"not null" json:"user_id"`
}
