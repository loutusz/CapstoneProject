package models

type Connection struct {
	ID                  string `json:"id"`
	Project_id          string `json:"project_id" form:"project_id"`
	Message_provider_id string `json:"message_provider_id" form:"message_provider_id"`
}

type UpsertConnection struct {
	Project_id          string `json:"project_id" form:"project_id"`
	Message_provider_id string `json:"message_provider_id" form:"message_provider_id"`
}

func (p Connection) UpsertConnection() UpsertConnection {
	return UpsertConnection{
		Message_provider_id: p.Message_provider_id,
		Project_id:          p.Project_id,
	}
}

type GetConnectionRequest struct {
	ID string `json:"id"`
}

type GetConnectionResponse struct {
	ID                  string `json:"id"`
	Project_id          string `json:"project_id"`
	Message_provider_id string `json:"message_provider_id"`
}

type PostConnectionResponse struct {
	ID                  string `json:"id"`
	Project_id          string `json:"project_id"`
	Message_provider_id string `json:"message_provider_id"`
}
