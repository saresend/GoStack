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
