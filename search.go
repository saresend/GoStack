package GoStack

import "net/http"
import "io/ioutil"
import "encoding/json"

//SearchStackExchange calls the search Functionality of stackexchange
func SearchStackExchange(keywords []string, site string) ([]SearchItem, error) {
	url := "https://api.stackexchange.com/2.2/search?order=desc&sort=activity&site=" + site + "&tagged=" + addParams(keywords)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response SearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}

func addParams(keywords []string) string {
	str := ""
	for i := 0; i < len(keywords); i++ {
		str += keywords[i]
		if i+1 < len(keywords) {
			str += "%3B"
		}
	}
	return str
}
