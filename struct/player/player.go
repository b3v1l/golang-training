package player

type Avatar struct {
	Url string
}

type Player struct {
	Name   string
	Age    int
	Avatar Avatar
	//private value password
	password string
}

func New(name string) Player {
	return Player{
		Name:     name,
		password: "defaultpassword",
		Avatar:   Avatar{Url: "http://polo.polo"},
	}
}
