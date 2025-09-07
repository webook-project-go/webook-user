package domain

type User struct {
	Id       int64
	Email    string
	Password string
	Phone    string
	Name     string
	Birthday string
	AboutMe  string
}

const (
	ProviderUnknown = "UNKNOWN"
	ProviderWechat  = "WECHAT"
	ProviderGithub  = "GITHUB"
)

type Oauth2Binding struct {
	UserID int64
	ID     int64

	Provider   string
	ExternalID string

	AccessToken string
}
