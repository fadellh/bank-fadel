package model

type InsertUserRequest struct {
	Fullname      string `json:"fullname"`
	Address       string `json:"address"`
	BirthdayPlace string `json:"birthday_place"`
	Birthday      string `json:"birthday"`
	KTPnumber     string `json:"no_ktp"`
	PhoneNumber   string `json:"phone_number"`
}

type UpdatetUserRequest struct {
	Fullname      string `json:"fullname"`
	Address       string `json:"address"`
	BirthdayPlace string `json:"birthday_place"`
	Birthday      string `json:"birthday"`
	PhoneNumber   string `json:"phone_number"`
}

type Users struct {
	Fullname      string
	Address       string
	BirthdayPlace string
	Birthday      string
	KTPnumber     string
	PhoneNumber   string
}
