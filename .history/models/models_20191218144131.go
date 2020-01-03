package models

// Gamer is an struct defining an individual user.
type Gamer struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gamelist []Game `json:"gamelist"`
}

// Game is a struct defining a single game.
type Game struct {
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Platform string `json:"platform"`
}

// Opts are projections options for specifying fields from a record.
type Opts struct {
	oOtions []string `json:"opts"`
}
