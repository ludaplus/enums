package basic

import (
	"encoding/json"
	"github.com/ludaplus/enums"
)

type PostTypeElement struct {
	enums.Element[*PostTypeElement]
	CommentEnabled bool
}

func (e *PostTypeElement) UnmarshalJSON(data []byte) error {
	var name string
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	*e = **(e.Element.UnmarshalHelper(name))
	return nil
}

var PostType = enums.Of(&struct {
	enums.Enum[*PostTypeElement]
	Unknown,
	Post,
	Page,
	Note *PostTypeElement
}{
	enums.Enum[*PostTypeElement]{},
	&PostTypeElement{
		CommentEnabled: false,
	},
	&PostTypeElement{
		CommentEnabled: true,
	},
	&PostTypeElement{
		CommentEnabled: false,
	},
	&PostTypeElement{
		CommentEnabled: false,
	},
})
