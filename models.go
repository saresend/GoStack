package GoStack

//Owner represents a comment or post owner
type Owner struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}

//Comment represents a given comment
type Comment struct {
	Owner        Owner  `json:"owner"`
	Edited       string `json:"edited"`
	Score        int    `json:"score"`
	CreationDate int    `json:"creation_date"`
	PostID       int    `json:"post_id"`
	CommentID    int    `json:"comment_id"`
}

//Response represents the outer layer of json response
type Response struct {
	Items []Comment `json:"items"`
}
