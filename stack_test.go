package GoStack

import "testing"

func TestGetComments(t *testing.T) {

	comments := GetComments("stackoverflow")
	if len(comments) == 0 {
		t.Log(comments)
		t.Fatal("Couldn't load comments")
	} else {
		t.Log(len(comments))
	}
}

func TestGetAnswers(t *testing.T) {
	questions := GetQuestions("stackoverflow")
	answers := questions[0].GetAnswers()
	println(answers)
}
