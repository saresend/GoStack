package GoStack

import "net/http"
import "io/ioutil"
import "encoding/json"
import "strconv"

//GetQuestions gets all the questions from a given site
func GetQuestions(site string) ([]Question, error) {
	resp, err := http.Get("https://api.stackexchange.com/2.2/questions?filter=withbody&site=" + site)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var questionResponse QuestionResponse
	json.Unmarshal(body, &questionResponse)

	return questionResponse.Items, nil
}

//GetAnswers returns the relevant answers to a question
func (q *Question) GetAnswers() ([]Answer, error) {
	url := "https://api.stackexchange.com/2.2/questions/" + strconv.Itoa(q.QuestionID) + "/answers"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var aR AnswerResponse
	json.Unmarshal(body, &aR)
	return aR.Items, nil

}
