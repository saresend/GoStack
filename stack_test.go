package GoStack

import "testing"

func TestGetComments(t *testing.T) {
	comments, err := GetComments("stackoverflow")
	if err != nil {
		t.Log(err.Error())
		//t.Fatal("Couldn't load comments")
	} else {
		t.Log(len(comments))
	}
}

func TestGetAnswers(t *testing.T) {
	questions, err := GetQuestions("stackoverflow")
	if err != nil {
		t.Log(err.Error())
	}
	if len(questions) > 0 {
		answers := questions[0].GetAnswers()
		println(answers)
	}

}
