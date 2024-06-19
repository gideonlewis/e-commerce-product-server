package datatypes

type Boolean bool

const (
	TRUE  Boolean = true
	FALSE Boolean = false
)

func (b *Boolean) IsNull() bool {
	return b == nil
}

func (b *Boolean) IsTrue() bool {
	return *b == TRUE
}

func (b *Boolean) IsFalse() bool {
	return *b == FALSE
}
