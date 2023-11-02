package plugin

import (
	"errors"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Kong/go-pdk"
)

var (
	ErrPathFailed = errors.New("could not get value for path")
	ErrPathEmpty  = errors.New("path can not be empty, check the plugin configuration")
)

type (
	StaticFile struct {
		ContentType string `json:"contentType"`
		Content     string `json:"content" schema:"{'required':true}"`
	}

	Config struct {
		Paths map[string]StaticFile `json:"paths" schema:"{'elements':{'type':'string','starts_with':'/','match_none':[{'pattern':'//','err':'must not have empty segments'}]}}"`
	}
)

func init() {
	// make sure txt is available
	if err := mime.AddExtensionType(".txt", "text/plain"); err != nil {
		panic(err)
	}
}

func New() interface{} {
	return &Config{}
}

func (sf StaticFile) ResolveContentType() string {
	if sf.ContentType != "" {
		return sf.ContentType
	}

	ext := filepath.Ext(sf.ContentType)
	if ct := mime.TypeByExtension(ext); ct != "" {
		return ct
	}

	return "application/octet-stream"
}

func (c *Config) Access(kong *pdk.PDK) {
	path, err := kong.Request.GetPath()
	if err != nil {
		_ = kong.Log.Err(ErrPathFailed)
		kong.Response.ExitStatus(http.StatusInternalServerError)
		return
	}

	if path != "" {
		path = strings.ToLower(path)
	}

	var (
		sf StaticFile
		ok bool
	)

	if sf, ok = c.Paths[path]; !ok {
		kong.Response.ExitStatus(http.StatusNotFound)
	}

	kong.Response.Exit(http.StatusOK, []byte(sf.Content), map[string][]string{"Content-Type": {sf.ResolveContentType()}})
}
