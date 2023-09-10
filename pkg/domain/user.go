package domain

type Users struct {
	Uname    string `bson:"uname" json:"uname"`
	Fname    string `bson:"fname" json:"fname" `
	Lname    string `bson:"lname" json:"lname"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required,min=6"`
}

type UsersResponse struct {
	Uname    string `bson:"uname" json:"uname"`
	Fname    string `bson:"fname" json:"fname" `
	Lname    string `bson:"lname" json:"lname"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required,min=6"`
}



