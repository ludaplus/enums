package basic

import "github.com/ludaplus/enums"

type PostTypeElement struct {
	enums.Element
	CommentEnabled bool
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
