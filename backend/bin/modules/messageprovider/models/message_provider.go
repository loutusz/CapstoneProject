package models

type MessageProvider struct {
	ID             string `json:"id"`
	Project_id     string `json:"project_id" form:"project_id"`
	Provider_type  string `json:"provider_type" form:"provider_type"`
	Provider_label string `json:"provider_label" form:"provider_label"`
	Channel        string `json:"channel" form:"channel"`
}

type UpsertMessageProvider struct {
	Project_id     string `json:"project_id" form:"project_id"`
	Provider_type  string `json:"provider_type" form:"provider_type"`
	Provider_label string `json:"provider_label" form:"provider_label"`
	Channel        string `json:"channel" form:"channel"`
}

func (p MessageProvider) UpsertMessageProvider() UpsertMessageProvider {
	return UpsertMessageProvider{
		Project_id:     p.Project_id,
		Provider_type:  p.Provider_type,
		Provider_label: p.Provider_label,
		Channel:        p.Channel,
	}
}

type GetMessageProviderRequest struct {
	ID string `json:"id"`
}

type GetMessageProviderResponse struct {
	ID             string `json:"id"`
	Project_id     string `json:"project_id" form:"project_id"`
	Provider_type  string `json:"provider_type" form:"provider_type"`
	Provider_label string `json:"provider_label" form:"provider_label"`
	Channel        string `json:"channel" form:"channel"`
}

type PostMessageProviderResponse struct {
	ID             string `json:"id"`
	Project_id     string `json:"project_id"`
	Provider_type  string `json:"provider_type"`
	Provider_label string `json:"provider_label"`
	Channel        string `json:"channel"`
}
