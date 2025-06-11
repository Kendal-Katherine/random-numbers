package database

type RandomNumber struct {
	ID        string `json:"id"`
	Number    int    `json:"number"`
	Seed      string `json:"seed"`
	CreatedAt string `json:"createdAt"`
}
