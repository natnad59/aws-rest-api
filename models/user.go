package models

type User struct {
	Id             string `json:"id" dynamodbav:"id"`
	Name           string `json:"name" dynamodbav:"name"`
	Catchphrase    string `json:"catchphrase" dynamodbav:"catchphrase"`
	FavoriteNumber int    `json:"favoriteNumber" dynamodbav:"favoriteNumber"`
	Retired        bool   `json:"retired" dynamodbav:"retired"`
}

var SampleUser = User{
	Id:             "1",
	Name:           "Sheldon",
	Catchphrase:    "Bazinga",
	FavoriteNumber: 2,
	Retired:        false,
}

var SampleUsers = []User{
	{
		Id:             "1",
		Name:           "Sheldon",
		Catchphrase:    "Bazinga",
		FavoriteNumber: 2,
		Retired:        false,
	},
	{
		Id:             "2",
		Name:           "Bob",
		Catchphrase:    "zooweemama",
		FavoriteNumber: 45,
		Retired:        true,
	},
}
