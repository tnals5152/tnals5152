package model

type Git struct {
	Ref        sstring     `json:"ref"`
	Pusher     Pusher     `json:"pusher"`
	Repository Repository `json:"repository"`
}

type Repository struct {
	Name string `json:"name"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
