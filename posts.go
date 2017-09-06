package GoStack

import "net/http"
import "io/ioutil"
import "encoding/json"

//GetPosts retrieves all posts in a given forum
func GetPosts(site string) ([]Post, error) {
	url := "https://api.stackexchange.com/2.2/posts?filter=withbody&site=" + site
	println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	var response PostResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}
