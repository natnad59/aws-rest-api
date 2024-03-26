package models

type User struct {
	Id             string
	Name           string
	Catchprase     string
	FavoriteNumber int
	Retired        bool
}

var SampleUser = User{
	Id:         "1",
	Name:       "Sheldon",
	Catchprase: "Bazinga",
	Retired:    false,
}

var SampleUsers = []User{
	{
		Id:         "1",
		Name:       "Sheldon",
		Catchprase: "Bazinga",
		Retired:    false,
	},
	{
		Id:         "2",
		Name:       "Bob",
		Catchprase: "zooweemama",
		Retired:    true,
	},
}
