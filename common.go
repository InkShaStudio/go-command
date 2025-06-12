package command

type SCommonAttributes[T any] struct {
	Name        string `json:"name"`
	Value       T      `json:"value"`
	Description string `json:"description"`
}

type ICommonInterface interface {
	GetValue() any
	GetName() string
	GetDescription() string
}

type ICommonAttributeChange[T any, V any] interface {
	ChangeName(name string) T
	ChangeValue(value V) T
	ChangeDescription(description string) T
}

func newCommanAttributes[T any](name string) SCommonAttributes[T] {
	var val T

	return SCommonAttributes[T]{
		Name:        name,
		Value:       val,
		Description: "",
	}
}

func (common *SCommonAttributes[T]) GetValue() any {
	return &common.Value
}

func (common *SCommonAttributes[T]) GetDescription() string {
	return common.Description
}

func (common *SCommonAttributes[T]) GetName() string {
	return common.Name
}
