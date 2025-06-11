package database

type RandomNumber struct {
	ID        uint   `gorm:"primaryKey"`
	Number    int    `json:"number"`
	Seed      string `json:"seed"`
	CreatedAt string `json:"createdAt"`
}
