package models

type MessageProvider struct {
	ID             string `json:"id"`
	User_id        string `json:"user_id" form:"user_id"`
	Provider_type  string `json:"provider_type" form:"provider_type"`
	Provider_label string `json:"provider_label" form:"provider_label"`
	Webhook        string `json:"webhook" form:"webhook"`
}

type UpsertMessageProvider struct {
	User_id        string `json:"user_id" form:"user_id"`
	Provider_type  string `json:"provider_type" form:"provider_type"`
	Provider_label string `json:"provider_label" form:"provider_label"`
	Webhook        string `json:"webhook" form:"webhook"`
}

func (p MessageProvider) UpsertMessageProvider() UpsertMessageProvider {
	return UpsertMessageProvider{
		User_id:        p.User_id,
		Provider_type:  p.Provider_type,
		Provider_label: p.Provider_label,
		Webhook:        p.Webhook,
	}
}

type GetMessageProviderRequest struct {
	ID string `json:"id"`
}

type GetMessageProviderResponse struct {
	ID             string `json:"id"`
	User_id        string `json:"user_id" form:"user_id"`
	Provider_type  string `json:"provider_type" form:"provider_type"`
	Provider_label string `json:"provider_label" form:"provider_label"`
	Webhook        string `json:"webhook" form:"webhook"`
}

type PostMessageProviderResponse struct {
	ID             string `json:"id"`
	User_id        string `json:"user_id"`
	Provider_type  string `json:"provider_type"`
	Provider_label string `json:"provider_label"`
	Webhook        string `json:"webhook"`
}
