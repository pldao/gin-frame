package form

type UserSocial struct {
	UserName   string `form:"username" json:"username"  binding:"required,min=5"`
	SocialName string `form:"social_name" json:"social_name"  binding:"required,min=1"`
	SocialLink string `form:"social_link" json:"social_link"  binding:"required,min=5"`
}

func NewUserSocialForm() *UserSocial {
	return &UserSocial{}
}

type UserName struct {
	UserName string `form:"username" json:"username"  binding:"required,min=5"`
}

func NewUserNameForm() *UserName {
	return &UserName{}
}
