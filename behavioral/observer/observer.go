package observer

type Observer interface {
	Update(template string, value string)
}
