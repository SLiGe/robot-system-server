package plugin

type Plugin interface {
	Code() string
	Do()
}
