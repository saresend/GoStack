package GoStack

import "net/http"
import "io/ioutil"
import "encoding/json"

//GetQuestions gets all the questions from a given site
func GetQuestions(site string) []Question {
	resp, err := http.Get("https://api.stackexchange.com/2.2/questions?site=" + site)
	if err != nil {
		return []Question{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	var questionResponse QuestionResponse
	json.Unmarshal(body, &questionResponse)

	return questionResponse.Items
}
