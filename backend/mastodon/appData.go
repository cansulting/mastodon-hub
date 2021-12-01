package mastodon

type AppData struct {
	ClientId  string `json:"client_id"`
	ClientKey string `json:"client_secret"`
	VapIdKey  string `json:"vapid_key"`
}
