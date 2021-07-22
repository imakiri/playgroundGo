package web

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/web/content"
	"github.com/imakiri/gorum/internal/web/transport"
	"github.com/imakiri/gorum/pkg/utils"
	"io/ioutil"
	"net/http"
	"time"
)

const path = "internal/web/content/"

type contentService struct {
	transport.UnimplementedContentServer
}

func (s *contentService) Get(context.Context, *transport.Request) (*transport.Content, error) {
	var err error

	var c = new(transport.Content)
	if c.Ico, err = ioutil.ReadFile(path + "ico.png"); err != nil {
		return nil, err
	}
	if c.Css, err = ioutil.ReadFile(path + "style.css"); err != nil {
		return nil, err
	}

	c.Html = new(transport.ContentHtml)
	c.Html.Layout = new(transport.Template)
	if c.Html.Layout.Rel, err = ioutil.ReadFile(path + "html/layout.html"); err != nil {
		return nil, err
	}
	if c.Html.Layout.Dev, err = ioutil.ReadFile(path + "html/layout_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body = new(transport.ContentHtmlBody)
	c.Html.Body.Pages = new(transport.ContentHtmlBodyPages)
	if c.Html.Body.Pages.Gorum, err = ioutil.ReadFile(path + "html/body/pages/gorum.html"); err != nil {
		return nil, err
	}
	if c.Html.Body.Pages.Index, err = ioutil.ReadFile(path + "html/body/pages/index.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum = new(transport.ContentHtmlBodyGorum)
	c.Html.Body.Gorum.Layout = new(transport.Template)
	if c.Html.Body.Gorum.Layout.Rel, err = ioutil.ReadFile(path + "html/body/gorum/layout.html"); err != nil {
		return nil, err
	}
	if c.Html.Body.Gorum.Layout.Dev, err = ioutil.ReadFile(path + "html/body/gorum/layout_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum.Content = new(transport.ContentHtmlBodyGorumContent)
	c.Html.Body.Gorum.Content.Pages = new(transport.ContentHtmlBodyGorumContentPages)
	c.Html.Body.Gorum.Content.Pages.Index = new(transport.Template)
	if c.Html.Body.Gorum.Content.Pages.Index.Dev, err = ioutil.ReadFile(path + "html/body/gorum/content/pages/index_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum.Content.Pages.Notifications = new(transport.Template)
	if c.Html.Body.Gorum.Content.Pages.Notifications.Dev, err = ioutil.ReadFile(path + "html/body/gorum/content/pages/notifications_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum.Content.Pages.LogIn = new(transport.Template)
	if c.Html.Body.Gorum.Content.Pages.LogIn.Dev, err = ioutil.ReadFile(path + "html/body/gorum/content/pages/login_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum.Content.Pages.Settings = new(transport.Template)
	if c.Html.Body.Gorum.Content.Pages.Settings.Dev, err = ioutil.ReadFile(path + "html/body/gorum/content/pages/settings_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum.Content.Pages.SignUp = new(transport.Template)
	if c.Html.Body.Gorum.Content.Pages.SignUp.Dev, err = ioutil.ReadFile(path + "html/body/gorum/content/pages/signup_dev.html"); err != nil {
		return nil, err
	}

	c.Html.Body.Gorum.Content.Pages.Profile = new(transport.Template)
	if c.Html.Body.Gorum.Content.Pages.Profile.Dev, err = ioutil.ReadFile(path + "html/body/gorum/content/pages/profile_dev.html"); err != nil {
		return nil, err
	}

	fmt.Printf("[%s] content/get\n", time.Now().Format("2006-01-02 15:04:05"))
	return c, nil
}

func NewContentService() (*contentService, error) {
	var s = new(contentService)

	var _, err = s.Get(context.Background(), new(transport.Request))
	if err != nil {
		return nil, err
	}

	return s, nil
}

type webService struct {
	dev      bool
	reload   bool
	https    bool
	services struct {
		content transport.ContentClient
	}
	content *content.Content
	router  *mux.Router
}

func (s *webService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if s.reload {
		var err = s.load()
		if err != nil {
			fmt.Println(err)
		}
	}
	s.router.ServeHTTP(writer, request)
}

func (s *webService) load() error {
	var raw, err = s.services.content.Get(context.Background(), &transport.Request{})
	if err != nil {
		return err
	}

	s.content = new(content.Content)
	s.content.Static.Css = raw.Css
	s.content.Static.Ico = raw.Ico

	fmt.Printf("%s\n", raw.Html.Body.Gorum.Content.Pages.LogIn.Dev)

	// s.content.Root.Index
	s.content.Root.Index.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Pages.Index)
	if err != nil {
		return err
	}

	s.content.Root.Index.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Pages.Index)
	if err != nil {
		return err
	}

	// s.content.Root.Gorum
	s.content.Root.Gorum.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Pages.Gorum)
	if err != nil {
		return err
	}

	s.content.Root.Gorum.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Pages.Gorum)
	if err != nil {
		return err
	}

	// s.content.Gorum.Index
	s.content.Gorum.Index.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Gorum.Layout.Dev, raw.Html.Body.Gorum.Content.Pages.Index.Dev)
	if err != nil {
		return err
	}

	s.content.Gorum.Index.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Gorum.Layout.Rel, raw.Html.Body.Gorum.Content.Pages.Index.Rel)
	if err != nil {
		return err
	}

	// s.content.Gorum.User.LogIn
	s.content.Gorum.User.LogIn.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Gorum.Layout.Dev, raw.Html.Body.Gorum.Content.Pages.LogIn.Dev)
	if err != nil {
		return err
	}

	s.content.Gorum.User.LogIn.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Gorum.Layout.Rel, raw.Html.Body.Gorum.Content.Pages.LogIn.Rel)
	if err != nil {
		return err
	}

	// s.content.Gorum.User.Notifications
	s.content.Gorum.User.Notifications.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Gorum.Layout.Dev, raw.Html.Body.Gorum.Content.Pages.Notifications.Dev)
	if err != nil {
		return err
	}

	s.content.Gorum.User.Notifications.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Gorum.Layout.Rel, raw.Html.Body.Gorum.Content.Pages.Notifications.Rel)

	// s.content.Gorum.User.Profile
	s.content.Gorum.User.Profile.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Gorum.Layout.Dev, raw.Html.Body.Gorum.Content.Pages.Profile.Dev)
	if err != nil {
		return err
	}

	s.content.Gorum.User.Profile.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Gorum.Layout.Rel, raw.Html.Body.Gorum.Content.Pages.Profile.Rel)
	if err != nil {
		return err
	}

	// s.content.Gorum.User.Settings
	s.content.Gorum.User.Settings.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Gorum.Layout.Dev, raw.Html.Body.Gorum.Content.Pages.Settings.Dev)
	if err != nil {
		return err
	}

	s.content.Gorum.User.Settings.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Gorum.Layout.Rel, raw.Html.Body.Gorum.Content.Pages.Settings.Rel)
	if err != nil {
		return err
	}

	// s.content.Gorum.User.SignUp
	s.content.Gorum.User.SignUp.Dev, err = content.NewTemplate(raw.Html.Layout.Dev, raw.Html.Body.Gorum.Layout.Dev, raw.Html.Body.Gorum.Content.Pages.SignUp.Dev)
	if err != nil {
		return err
	}

	s.content.Gorum.User.SignUp.Rel, err = content.NewTemplate(raw.Html.Layout.Rel, raw.Html.Body.Gorum.Layout.Rel, raw.Html.Body.Gorum.Content.Pages.SignUp.Rel)
	if err != nil {
		return err
	}

	return nil
}

func (s *webService) header(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Origin", "https://i.imgur.com")
}

func (webService) ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)

	_, _ = w.Write([]byte(err.Error()))
}

func (webService) push(w http.ResponseWriter, _ *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		var err error
		if err = pusher.Push("/static/css", nil); err != nil {
			fmt.Println(err)
		}
		if err = pusher.Push("/static/ico", nil); err != nil {
			fmt.Println(err)
		}
	}
}

func NewWebService(dev, reload, https bool, contentClient transport.ContentClient) (*webService, error) {
	if utils.IsNil(contentClient) {
		return nil, erres.NilArgument
	}

	var s = new(webService)
	s.dev = dev
	s.reload = reload
	s.https = https
	s.router = mux.NewRouter()
	s.services.content = contentClient

	var err = s.load()
	if err != nil {
		return nil, err
	}

	s.router.HandleFunc("/", s.rootIndex)
	s.router.HandleFunc("/gorum", s.rootGorum)
	s.router.HandleFunc("/static/css", s.rootStaticCss)
	s.router.HandleFunc("/static/ico", s.rootStaticIco)
	s.router.HandleFunc("/gorum/", s.rootGorumIndex)
	s.router.HandleFunc("/gorum/user/login", s.rootGorumUserLogIn)
	s.router.HandleFunc("/gorum/user/notifications", s.rootGorumUserNotifications)
	s.router.HandleFunc("/gorum/user/profile", s.rootGorumUserProfile)
	s.router.HandleFunc("/gorum/user/signup", s.rootGorumUserSignup)
	s.router.HandleFunc("/gorum/user/settings", s.rootGorumUserSettings)

	return s, nil
}
