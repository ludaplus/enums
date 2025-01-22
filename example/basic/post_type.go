package basic

import "github.com/ludaplus/enums"

type innerPostType struct {
	enums.Element
	CommentEnabled bool
}

var PostType = enums.Of(&struct {
	enums.Enum[*innerPostType]
	Post,
	Page,
	Unknown *innerPostType
}{
	Unknown: &innerPostType{
		CommentEnabled: false,
	},
	Page: &innerPostType{
		CommentEnabled: false,
	},
	Post: &innerPostType{
		CommentEnabled: true,
	},
})
