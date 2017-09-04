package GoStack

import "testing"

func TestGetComments(t *testing.T) {
	cmt := NewCommentStruct("stackoverflow")
	comments := cmt.GetComments()
	if len(comments) == 0 {
		t.Log(comments)
		t.Fatal("Couldn't load comments")
	} else {
		t.Log(len(comments))
	}
}
