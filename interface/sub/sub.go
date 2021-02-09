package sub

type IInfo interface {
	Text() string
	SetText(text string)
}

type impl struct {
	name string
}

func (instance *impl)Text() string {
	return instance.name
}

func (instance *impl)SetText(text string) {
	instance.name = text
}

func CreateInfo(name string) IInfo {
	instance := impl{name}
	return &instance
}