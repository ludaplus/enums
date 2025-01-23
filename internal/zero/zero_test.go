package zero

import (
	"fmt"
	"testing"
)

func TestZero(t *testing.T) {
	if StringPostTypes.Unknown != "Unknown" {
		t.Errorf("unknown should be Unknown but `%s`", StringPostTypes.Unknown)
	}

	if NumberPostTypes.Post != 1 {
		t.Errorf("post should be 1 but `%d`", NumberPostTypes.Post)
	}

	fmt.Println(StringPostTypes.Unknown)
}
