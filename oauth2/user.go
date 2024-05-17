package oauth2

type User struct {
	ID            string                 `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Provider      string                 `json:"provider,omitempty"`
	Email         string                 `json:"email,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	FirstName     string                 `json:"first_name,omitempty"`
	LastName      string                 `json:"last_name,omitempty"`
	NickName      string                 `json:"nick_name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	AvatarURL     string                 `json:"avatar_url,omitempty"`
	Location      string                 `json:"location,omitempty"`
	RawData       map[string]interface{} `json:"raw_data,omitempty"`
}
