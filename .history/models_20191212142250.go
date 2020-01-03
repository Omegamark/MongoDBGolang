package models

// User is an struct defining an individual user.
type Gamer struct {
	Name     string
	Age      int
	Gamelist []Game
}

// Game is a struct defining a single game.
type Game struct {
	Name     string
	Genre    string
	Platform string
}
