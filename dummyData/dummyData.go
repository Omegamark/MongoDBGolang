package data

import "MongoDBGolang/models"

var games1 = []models.Game{{"Megaman", "Platformer", "NES"}, {"Daggerfall", "CRPG", "PC"}, {"Zelda64", "Adventure", "N64"}}
var gamer1 = models.Gamer{
	"Mark",
	38,
	games1,
}
var games2 = []models.Game{{"Pocky & Rocky", "Action", "SNES"}, {"Final Fantasy III", "JRPG", "SNES"}, {"Contra Hard Corps", "Action", "Sega Genesis"}}
var gamer2 = models.Gamer{
	"Brian",
	37,
	games2,
}
var games3 = []models.Game{{"Shining Force", "TRPG", "Sega Genesis"}, {"Shinobi", "Action", "Arcade"}, {"Phantasy Star", "JRPG", "Sega Master System"}}
var gamer3 = models.Gamer{
	"Sean",
	28,
	games3,
}
var games4 = []models.Game{{"Chess", "Board", "Commodore"}, {"Bakarat", "Ancient", "Wood"}, {"Cyberpunk 2077", "WRPG", "PC"}}
var gamer4 = models.Gamer{
	"Bean",
	65,
	games4,
}
