package content

import "html/template"

type Template struct {
	*template.Template
}

func NewTemplate(raw ...[]byte) (Template, error) {
	var t = new(Template)
	t.Template = template.New("index")

	var err error
	for i := range raw {
		t.Template, err = t.Parse(string(raw[i]))
		if err != nil {
			return Template{}, err
		}
	}

	return *t, nil
}

type Templates struct {
	Dev Template
	Rel Template
}

type Content struct {
	Root struct {
		Index Templates
		Gorum Templates
	}
	Static struct {
		Ico []byte
		Css []byte
	}
	Gorum struct {
		Index Templates
		User  struct {
			SignUp        Templates
			LogIn         Templates
			Profile       Templates
			Notifications Templates
			Settings      Templates
		}
	}
}
