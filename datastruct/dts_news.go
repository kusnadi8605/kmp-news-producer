package datastruct

//News struct for news
type News struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

//NewsResponse ..
type NewsResponse struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
}
