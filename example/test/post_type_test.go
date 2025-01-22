package test

import (
	"fmt"
	"github.com/ludaplus/enums/example/basic"
	"testing"
)

func TestPostType(t *testing.T) {
	values := basic.PostType.Values()
	size := 4
	if len(values) != size {
		t.Errorf("values length is not %d but %d", size, len(values))
	}
	if !basic.PostType.Post.CommentEnabled {
		t.Errorf("post should enable comment")
	}

	if *basic.PostType.ValueOf("Post") != basic.PostType.Post {
		t.Errorf("post should be the same")
	}

	fmt.Printf("PostType test passed: %v\n", basic.PostType)
}
