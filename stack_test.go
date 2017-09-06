package GoStack

import "testing"

func TestGetComments(t *testing.T) {
	comments, err := GetComments("stackoverflow")
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(len(comments))
		if len(comments) == 0 {
			t.Fatal("No Http error, but failed to read comments")
		}
	}
}

func TestGetAnswers(t *testing.T) {
	questions, err := GetQuestions("stackoverflow")
	if err != nil {
		t.Log(err.Error())
	}
	if len(questions) > 0 {
		answers, err := questions[0].GetAnswers()
		if err != nil {
			t.Log(err.Error())
		}
		println(answers)
	}

}

func TestGetPosts(t *testing.T) {
	posts, err := GetPosts("stackoverflow")
	if err != nil {
		t.Log(err.Error())
	} else {
		if len(posts) == 0 {
			t.Fatal("No HTTP error, but no posts")
		}
	}
}

func TestAuthentication(t *testing.T) {
	token, err := GetAccessToken("10737", "40SqxdGUyX01sxjH2UVwJw((")
	if err == nil {
		if token == "" {
			t.Log(token)
			t.Fatal("Couldn't find token")
		}
	}
}
