package model

type UserProfile struct {
	*User
	*Profile
}

type ProfileApiEditAvatarReq struct {
	Nickname string `v:"required-without-all:Motto"`
	Motto    string `v:"required-without-all:Nickname"`
}

type ProfileServiceEditAvatarReq struct {
	Nickname string `v:"required-without-all:Motto"`
	Motto    string `v:"required-without-all:Nickname"`
}
