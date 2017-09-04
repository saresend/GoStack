package GoStack

import "net/http"
import "io/ioutil"
import "encoding/json"

//CommentStruct is a structure containing the basic data for a given
type CommentStruct struct {
	baseURL string
	site    string
}

//GetComments makes a call to get comments on a given website
func GetComments(site string) ([]Comment, error) {
	url := "https://api.stackexchange.com/2.2/comments" + "?filter=withbody&site=" + site

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var comments CommentResponse
	json.Unmarshal(body, &comments)
	return comments.Items, nil
}
