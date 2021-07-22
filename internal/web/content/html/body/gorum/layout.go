package gorum

const (
	BodyPage_Main          = "main"
	BodyPage_LogIn         = "login"
	BodyPage_SignUp        = "signup"
	BodyPage_Notifications = "notifications"
	BodyPage_Settings      = "settings"
	BodyPage_Profile       = "profile"
)

type Body struct {
	Page         string
	IsAuthorized bool
	Content      interface{}
}
