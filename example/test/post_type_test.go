package test

import (
	"encoding/json"
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

	marshalled, _ := json.Marshal(basic.PostType.Post)

	unmarshalled := &basic.PostTypeElement{}
	_ = json.Unmarshal(marshalled, unmarshalled)

	if unmarshalled.CommentEnabled != basic.PostType.Post.CommentEnabled {
		t.Errorf("unmarshalled post field should be the same")
	}

	if !basic.PostType.Post.Equals(unmarshalled) {
		t.Errorf("unmarshalled post should be the same")
	}

}
