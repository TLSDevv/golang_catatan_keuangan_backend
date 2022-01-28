package entity

type Auth struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token`
}

type ResponseToken struct{
	Token string `json:"token"`
}
