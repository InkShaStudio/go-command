package command

type SCommandArg[T any] struct {
	SCommonAttributes[T]
	Index int `json:"index"`
}

type ICommandArg[T any] interface {
	ICommonAttributeChange[*ICommandArg[T], T]
	ChangeIndex(index int) *SCommandArg[T]
}

type ICommandArgValue interface {
	ICommonInterface
	GetIndex() int
	SetIndex(index int)
}

func (arg *SCommandArg[T]) ChangeIndex(index int) *SCommandArg[T] {
	arg.Index = index
	return arg
}

func (arg *SCommandArg[T]) ChangeName(name string) *SCommandArg[T] {
	arg.Name = name
	return arg
}

func (arg *SCommandArg[T]) ChangeValue(value T) *SCommandArg[T] {
	arg.Value = value
	return arg
}

func (arg *SCommandArg[T]) ChangeDescription(description string) *SCommandArg[T] {
	arg.Description = description
	return arg
}

func (arg *SCommandArg[T]) GetIndex() int {
	return arg.Index
}

func (arg *SCommandArg[T]) SetIndex(index int) {
	arg.ChangeIndex(index)
}

func NewCommandArg[T any](name string) *SCommandArg[T] {
	return &SCommandArg[T]{
		Index:             0,
		SCommonAttributes: newCommanAttributes[T](name),
	}
}
