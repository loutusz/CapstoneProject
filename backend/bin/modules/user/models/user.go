package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
}

type UpsertUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func (u User) UpsertUser() UpsertUser {
	return UpsertUser{
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
	}
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type GetUserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
