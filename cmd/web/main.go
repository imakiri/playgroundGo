package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/web"
	"github.com/imakiri/gorum/internal/web/transport"
	pkgHttp "github.com/imakiri/gorum/pkg/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net/http"
)

const path_ca = "secrets/web/ca.crt"

func connect(domain string, port string) (*grpc.ClientConn, error) {
	var err error
	var ca []byte

	if ca, err = ioutil.ReadFile(path_ca); err != nil {
		return nil, err
	}

	var cp = x509.NewCertPool()
	if !cp.AppendCertsFromPEM(ca) {
		return nil, erres.Error("certificate error").Extend(0)
	}

	var conf = &tls.Config{
		RootCAs:            cp,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(domain+":"+port, grpc.WithTransportCredentials(credentials.NewTLS(conf)))
	if err != nil {
		return nil, err
	}

	return conn, err
}

func NewLauncher(dev, reload, https bool, domain string, port string) *Launcher {
	var l = new(Launcher)
	l.dev = dev
	l.reload = reload
	l.https = https
	l.statusWeb = make(chan error)
	l.statusRedirector = make(chan error)

	return l
}

type Launcher struct {
	dev              bool
	reload           bool
	https            bool
	web              *pkgHttp.Server
	redirector       *pkgHttp.Redirector
	statusWeb        chan error
	statusRedirector chan error
}

func (l *Launcher) Prepare() error {
	var conn, err = connect(domain, port)
	if err != nil {
		return err
	}

	var cc = transport.NewContentClient(conn)

	var ws http.Handler
	ws, err = web.NewWebService(l.dev, l.reload, l.https, cc)
	if err != nil {
		return err
	}

	if l.https {
		l.redirector, err = pkgHttp.NewRedirector(l.statusRedirector)
		if err != nil {
			return err
		}

		l.web, err = pkgHttp.NewServer(ws, l.statusWeb, true)
		if err != nil {
			return err
		}
	} else {
		l.web, err = pkgHttp.NewServer(ws, l.statusWeb, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Launcher) Launch() error {
	if l.dev {
		l.web.Launch()
		return <-l.statusWeb
	} else {
		var err error
		l.web.Launch()
		l.redirector.Launch()

		select {
		case err = <-l.statusRedirector:
			l.web.Stop()
		case err = <-l.statusWeb:
			l.redirector.Stop()
		}

		return err
	}
}

const (
	domain = "imakiri-ips.ddns.net"
	port   = "25565"
)

func main() {
	var dev = flag.Bool("dev", true, "web dev environment")
	var reload = flag.Bool("reload", true, "reload html templates on request")
	var https = flag.Bool("https", false, "launch https")
	flag.Parse()

	var l = NewLauncher(*dev, *reload, *https, domain, port)
	var err = l.Prepare()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(l.Launch())
}
