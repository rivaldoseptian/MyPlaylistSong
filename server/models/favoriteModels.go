package models

import "time"

type Favorite struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	SongID    uint      `json:"song_id"`
	Song      Song      `gorm:"foreignKey:SongID" json:"song"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FavoriteResponse struct {
	ID     uint `json:"id"`
	UserID uint `json:"-"`
	User   User `json:"-"`
	SongID uint `json:"-"`
	Song   Song `gorm:"foreignKey:SongID" json:"song"`
}
