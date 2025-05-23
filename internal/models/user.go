package models

import "os"

type User struct {
	Username       string
	Session_Cookie string
	Database_URL   string
	Database_Key   string
}

func (u *User) GetUser() string {
	return u.Username
}

func (u *User) GetSession() string {
	return u.Session_Cookie
}

func (u *User) GetDbURL() string {
	return u.Database_URL
}
func (u *User) GetDbKey() string {
	return u.Database_Key
}

func CreateUser() User {
	return User{Username: os.Getenv("LEETCODE_USERNAME"),
		Session_Cookie: os.Getenv("LEETCODE_SESSION"),
		Database_URL:   os.Getenv("SUPABASE_URL"),
		Database_Key:   os.Getenv("SUPABASE_ANON_KEY")}
}
