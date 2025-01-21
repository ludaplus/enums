package enums

import "testing"

type PostType struct {
	CommentEnabled bool
}

var PostTypes = Of(&struct {
	*Enum[PostType]
	Unknown *PostType
	Page    *PostType
	Post    *PostType
}{
	Unknown: &PostType{
		CommentEnabled: false,
	},
	Page: &PostType{
		CommentEnabled: false,
	},
	Post: &PostType{
		CommentEnabled: true,
	},
})

func TestPostType(t *testing.T) {
	values := PostTypes.Values()
	size := 3
	if len(values) != size {
		t.Errorf("values length is not %d but %d", size, len(values))
	}
	if !PostTypes.Post.CommentEnabled {
		t.Errorf("post should enable comment")
	}
}
