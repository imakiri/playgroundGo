package web

import (
	"fmt"
	"net/http"
	"time"
)

func (s *webService) rootIndex(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] / hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Root.Index.Dev.Execute(w, nil)
	} else {
		err = s.content.Root.Index.Rel.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}

func (s *webService) rootGorum(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] / hit by ip:%s\n", t, r.RemoteAddr)

	s.header(w, r)
	s.push(w, r)

	var err error
	if s.dev {
		err = s.content.Root.Gorum.Dev.Execute(w, nil)
	} else {
		err = s.content.Root.Gorum.Rel.Execute(w, nil)
	}

	if err != nil {
		s.ise(w, err)
	}
}
