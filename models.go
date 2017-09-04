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

//QuestionResponse is the outer response provided by stack exchange
type QuestionResponse struct {
	Items []Question `json:"items"`
}

//Question represents the response from stackexchange api v.2.2
type Question struct {
	Tags         []string `json:"tags"`
	Owner        Owner    `json:"owner"`
	QuestionID   int      `json:"question_id"`
	IsAnswered   bool     `json:"is_answered"`
	ViewCount    int      `json:"view_count"`
	Score        int      `json:"score"`
	LastActivity int      `json:"last_activity_date"`
	CreationDate int      `json:"creation_date"`
	Link         string   `json:"link"`
	Title        string   `json:"title"`
	Body         string   `json:"body"`
}

type AnswerResponse struct {
	Items []Answer `json:"items"`
}

type Answer struct {
	Owner            Owner  `json:"owner"`
	IsAccepted       bool   `json:"is_accepted"`
	Score            int    `json:"score"`
	LastActivityDate int    `json:"last_activity_date"`
	CreationDate     int    `json:"creation_date"`
	AnswerID         int    `json:"answer_id"`
	QuestionID       int    `json:"question_id"`
	Body             string `json:"body"`
}
