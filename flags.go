package command

type SCommandFlag[T any] struct {
	SCommonAttributes[T]
	Short string `json:"short"`
}

type ICommandFlag[T any] interface {
	ICommonAttributeChange[*SCommandFlag[T], T]
	ChangeShort(short string) *SCommandFlag[T]
}

type ICommandFlagValue interface {
	ICommonInterface
	GetShort() string
}

func (flag *SCommandFlag[T]) ChangeName(name string) *SCommandFlag[T] {
	short := string([]rune(name)[0])
	flag.Name = name
	flag.Short = short
	return flag
}

func (flag *SCommandFlag[T]) ChangeValue(value T) *SCommandFlag[T] {
	flag.Value = value
	return flag
}

func (flag *SCommandFlag[T]) ChangeShort(short string) *SCommandFlag[T] {
	flag.Short = short
	return flag
}

func (flag *SCommandFlag[T]) ChangeDescription(description string) *SCommandFlag[T] {
	flag.Description = description
	return flag
}

func (flag *SCommandFlag[T]) GetShort() string {
	return flag.Short
}

func NewCommandFlag[T any](name string) *SCommandFlag[T] {
	short := string([]rune(name)[0])

	return &SCommandFlag[T]{
		Short:             short,
		SCommonAttributes: newCommanAttributes[T](name),
	}
}
