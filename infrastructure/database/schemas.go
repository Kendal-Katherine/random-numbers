package database

type RandomNumber struct {
	ID        string `json:"primaryKey"`
	Number    int    `json:"number"`
	Seed      string `json:"seed"`
	CreatedAt string `json:"createdAt"`
}

