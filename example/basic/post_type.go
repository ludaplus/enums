package basic

import "github.com/ludaplus/enums"

type innerPostType struct {
	enums.Element
	CommentEnabled bool
}

var PostType = enums.Of(&struct {
	enums.Enum[*innerPostType]
	Unknown *innerPostType
	Post,
	Page,
	Note *innerPostType
}{
	enums.Enum[*innerPostType]{},
	&innerPostType{
		CommentEnabled: false,
	},
	&innerPostType{
		CommentEnabled: true,
	},
	&innerPostType{
		CommentEnabled: false,
	},
	&innerPostType{
		CommentEnabled: false,
	},
})
