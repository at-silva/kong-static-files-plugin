package plugin_test

import (
	_ "embed"
	"net/http"
	"testing"

	"github.com/Kong/go-pdk/test"
	"github.com/at-silva/kong-plugin-static-files/plugin"
	. "github.com/onsi/gomega"
)

const (
	contentTypeTextPlain        = "text/plain"
	contentTypeApplicationJson  = "application/json"
	contentTypeApplicationOctet = "application/octet-stream"
)

var (
	//go:embed samples/assetlinks.json
	assetLinksJson string

	//go:embed samples/apple-app-site-association
	appleAppSiteAssociation string

	//go:embed samples/ads.txt
	adsTxt string

	//go:embed samples/robots.txt
	robotsTxt string
)

func TestPlugin(t *testing.T) {
	g := NewWithT(t)
	testCases := []struct {
		desc        string
		url         string
		path        string
		content     string
		contentType string
	}{
		{
			desc:        "should return assetlinks.json",
			url:         "https://example.com/.well-known/assetlinks.json",
			path:        "/.well-known/assetlinks.json",
			content:     assetLinksJson,
			contentType: contentTypeApplicationJson,
		},
		{
			desc:        "should return aasa",
			url:         "https://example.com/apple-app-site-association",
			path:        "/apple-app-site-association",
			content:     appleAppSiteAssociation,
			contentType: contentTypeApplicationJson,
		},
		{
			desc:        "should return aasa",
			url:         "https://example.com/.well-known/apple-app-site-association",
			path:        "/apple-app-site-association",
			content:     appleAppSiteAssociation,
			contentType: contentTypeApplicationJson,
		},
		{
			desc:        "should return ads.txt",
			url:         "https://example.com/ads.txt",
			path:        "/ads.txt",
			content:     adsTxt,
			contentType: contentTypeTextPlain,
		},
		{
			desc:        "should return robots.txt",
			url:         "https://example.com/robots.txt",
			path:        "/robots.txt",
			content:     robotsTxt,
			contentType: contentTypeTextPlain,
		},
	}

	config := &plugin.Config{
		Paths: map[string]plugin.StaticFile{
			"/ads.txt":                                {Content: adsTxt, ContentType: contentTypeTextPlain},
			"/apple-app-site-association":             {Content: appleAppSiteAssociation, ContentType: contentTypeApplicationJson},
			"/.well-known/apple-app-site-association": {Content: appleAppSiteAssociation, ContentType: contentTypeApplicationJson},
			"/.well-known/assetlinks.json":            {Content: assetLinksJson, ContentType: contentTypeApplicationJson},
			"/robots.txt":                             {Content: robotsTxt, ContentType: contentTypeTextPlain},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			env, err := test.New(t, test.Request{
				Method:  http.MethodGet,
				Url:     tC.url,
				Headers: map[string][]string{"Accept": {"*/*"}},
			})
			g.Expect(err).ShouldNot(HaveOccurred())
			env.DoHttps(config)
			g.Expect(http.StatusOK).Should(Equal(env.ClientRes.Status))
			g.Expect(string(tC.content)).Should(Equal(string(env.ClientRes.Body)))
		})
	}
}

func TestContentTypeResolution(t *testing.T) {
	g := NewWithT(t)
	testCases := []struct {
		desc        string
		url         string
		path        string
		content     string
		contentType string
	}{
		{
			desc:        "should return assetlinks.json",
			url:         "https://example.com/.well-known/assetlinks.json",
			path:        "/.well-known/assetlinks.json",
			content:     assetLinksJson,
			contentType: contentTypeApplicationJson,
		},
		{
			desc:        "should return aasa",
			url:         "https://example.com/apple-app-site-association",
			path:        "/apple-app-site-association",
			content:     appleAppSiteAssociation,
			contentType: contentTypeApplicationJson,
		},
		{
			desc:        "should return ads.txt",
			url:         "https://example.com/ads.txt",
			path:        "/ads.txt",
			content:     adsTxt,
			contentType: contentTypeTextPlain,
		},
		{
			desc:        "should return file.custom",
			url:         "https://example.com/file.custom",
			path:        "/file.custom",
			content:     adsTxt,
			contentType: contentTypeApplicationOctet,
		},
	}

	config := &plugin.Config{
		Paths: map[string]plugin.StaticFile{
			"/ads.txt":                     {Content: adsTxt, ContentType: ""},
			"/apple-app-site-association":  {Content: appleAppSiteAssociation, ContentType: contentTypeApplicationJson},
			"/.well-known/assetlinks.json": {Content: assetLinksJson, ContentType: ""},
			"/file.custom":                 {Content: adsTxt, ContentType: ""},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			env, err := test.New(t, test.Request{
				Method:  http.MethodGet,
				Url:     tC.url,
				Headers: map[string][]string{"Accept": {"*/*"}},
			})
			g.Expect(err).ShouldNot(HaveOccurred())
			env.DoHttps(config)
			g.Expect(http.StatusOK).Should(Equal(env.ClientRes.Status))
			g.Expect(string(tC.content)).Should(Equal(string(env.ClientRes.Body)))
		})
	}
}

func TestPathValidation(t *testing.T) {
	g := NewWithT(t)
	testCases := []struct {
		desc string
		url  string
	}{
		{
			desc: "path is empty",
			url:  "https://example.com/",
		},
		{
			desc: "path is unsupported",
			url:  "https://example.com/unsupported.json",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			config := &plugin.Config{}
			env, err := test.New(t, test.Request{
				Method:  http.MethodGet,
				Url:     tC.url,
				Headers: map[string][]string{"Accept": {contentTypeApplicationJson}},
			})
			g.Expect(err).ShouldNot(HaveOccurred())
			env.DoHttps(config)
			g.Expect(http.StatusNotFound).Should(Equal(env.ClientRes.Status))
		})
	}
}
