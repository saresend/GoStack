package GoStack

import "net/http"
import "io/ioutil"
import "encoding/json"

//CommentStruct is a structure containing the basic data for a given
type CommentStruct struct {
	baseURL string
	site    string
}

//NewCommentStruct instantiates an objects through which all functionality can be obtained
func NewCommentStruct(site string) CommentStruct {
	comment := CommentStruct{"https://api.stackexchange.com/2.2/comments", site}
	return comment
}

//GetComments makes a default call
func GetComments(site string) []Comment {
	url := "https://api.stackexchange.com/2.2/comments" + "?site=" + site

	resp, err := http.Get(url)
	if err != nil {
		println(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
	}
	var comments Response
	json.Unmarshal(body, &comments)

	return comments.Items
}
