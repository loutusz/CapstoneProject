package models

type Project struct {
	ID      string `json:"id"`
	User_id string `json:"user_id" form:"user_id"`
	Name    string `json:"name" form:"name"`
}

type UpsertProject struct {
	User_id string `json:"user_id" form:"user_id"`
	Name    string `json:"name" form:"name"`
}

func (p Project) UpsertProject() UpsertProject {
	return UpsertProject{
		Name:    p.Name,
		User_id: p.User_id,
	}
}

type GetProjectRequest struct {
	ID string `json:"id"`
}

type GetProjectResponse struct {
	ID      string `json:"id"`
	User_id string `json:"user_id"`
	Name    string `json:"name"`
}

type PostProjectResponse struct {
	ID      string `json:"id"`
	User_id string `json:"user_id"`
	Name    string `json:"name"`
}
