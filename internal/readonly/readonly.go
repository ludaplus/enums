package readonly

type PostType string

type (
	internalPostTypeInterface interface {
		Unknown() PostType
		Post() PostType
	}
)

func (t PostType) Unknown() PostType {
	return "Unknown"
}

func (t PostType) Post() PostType {
	return "Post"
}

const ReadOnly = PostType("")
