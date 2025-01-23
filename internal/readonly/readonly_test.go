package readonly

import (
	"fmt"
	"testing"
)

func TestReadonly(t *testing.T) {
	fmt.Println(ReadOnly.Unknown().Post())
}
