package samples

import (
	"bytes"
	_ "embed"
	"io"
)

//go:embed tech.txt
var TechAcronyms []byte

//go:embed web.txt
var WebAcronyms []byte

//go:embed business.txt
var BusinessAcronyms []byte

// ReadTechAcronyms returns an io.Reader for the tech acronyms.
func ReadTechAcronyms() io.Reader {
	return bytes.NewReader(TechAcronyms)
}

// ReadWebAcronyms returns an io.Reader for the web acronyms.
func ReadWebAcronyms() io.Reader {
	return bytes.NewReader(WebAcronyms)
}

// ReadBusinessAcronyms returns an io.Reader for the business acronyms.
func ReadBusinessAcronyms() io.Reader {
	return bytes.NewReader(BusinessAcronyms)
}
