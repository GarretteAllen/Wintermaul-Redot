package models

type Lobby struct {
	ID      string   `bson:"_id,omitempty"`
	Players []string `bson:"players"`
	MaxSize int      `bson:"max_size"`
}
