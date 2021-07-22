package web

import (
	"fmt"
	"net/http"
	"time"
)

func (s *webService) rootGorumIndex(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] /gorum hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Gorum.Index.Dev.Execute(w, nil)
	} else {
		err = s.content.Gorum.Index.Dev.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}

func (s *webService) rootGorumUserSignup(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] /gorum hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Gorum.User.SignUp.Dev.Execute(w, nil)
	} else {
		err = s.content.Gorum.User.SignUp.Rel.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}

func (s *webService) rootGorumUserSettings(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] /gorum hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Gorum.User.Settings.Dev.Execute(w, nil)
	} else {
		err = s.content.Gorum.User.Settings.Rel.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}

func (s *webService) rootGorumUserProfile(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] /gorum hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Gorum.User.Profile.Dev.Execute(w, nil)
	} else {
		err = s.content.Gorum.User.Profile.Rel.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}

func (s *webService) rootGorumUserNotifications(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] /gorum hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Gorum.User.Notifications.Dev.Execute(w, nil)
	} else {
		err = s.content.Gorum.User.Notifications.Dev.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}

func (s *webService) rootGorumUserLogIn(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] /gorum hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Gorum.User.LogIn.Dev.Execute(w, nil)
	} else {
		err = s.content.Gorum.User.LogIn.Rel.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}
