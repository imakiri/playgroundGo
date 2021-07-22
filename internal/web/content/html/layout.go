package html

const (
	Page_Home  = "home"
	Page_Gorum = "gorum"
)

type Main struct {
	Page string
	Body interface{}
}
