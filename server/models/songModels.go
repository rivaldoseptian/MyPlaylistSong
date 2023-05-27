package models

import "time"

type Song struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Duration  string    `gorm:"type:varchar(100)" json:"duration"`
	ArtisID   uint      `json:"artis_id"`
	Artis     Artis     `gorm:"foreignKey:ArtisID" json:"artis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SongResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
	ArtisID  uint   `json:"artis_id"`
	Artis    Artis  `gorm:"foreignKey:ArtisID" json:"artis"`
}
