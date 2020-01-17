const Post = {
    "name": "Mark",
    "age": 38,
    "gamelist": [
        {
            "name": "Zelda 64",
            "genre": "Adventure",
            "platform": "N64"
        },
        {
            "name": "Phantasy Star",
            "genre": "JRPG",
            "platform": "Genesis"
        },
        {
            "name": "Daggerfall",
            "genre": "CRPG",
            "platform": "PC"
        }
    ]
}

const Delete = {
    "name": "Mark"
}

const Get = {
    "name": "Sean",
    "opts": ["name", "gamelist", "age"]
}

const UpdateGamer = {
    "name": "Sean",
    "age": 55
}

const UpdateGamelist = {
    "name": "Mark",
    "game": {
        "name": "Bravely Default",
        "genre": "JRPG",
        "platform": "3ds"
    }
}
