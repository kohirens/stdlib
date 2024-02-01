package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kohirens/stdlib/log"
	"net/http"
	"path/filepath"
	"strings"
)

// GetPageType Get the content type via header.
func GetPageType(headers StringMap) string {
	ct := GetHeader(headers, "content-type")

	fct := strings.Split(ct, ",")
	if fct != nil {
		ct = fct[0]
	}

	return ct
}

// GetPageTypeByExt Get the content type by the extension of the file being
// requested.
func GetPageTypeByExt(pagePath string) string {
	var ct string

	ext := filepath.Ext(pagePath)

	switch ext {
	case ".css":
		ct = ContentTypeCSS
	case ".html":
		ct = ContentTypeHtml
	case ".js":
		ct = ContentTypeJS
	case ".json":
		ct = ContentTypeJson
	case ".jpg":
		ct = ContentTypeJpg
	case ".gif":
		ct = ContentTypeGif
	case ".png":
		ct = ContentTypePng
	case ".svg", ".svgz":
		ct = ContentTypeSvg
	default:
		ct = ""
	}

	return ct
}

// GetHeader Retrieve a header from a request.
func GetHeader(headers StringMap, name string) string {
	value := ""
	lcn := strings.ToLower(name)

	for h, v := range headers {
		lch := strings.ToLower(h)
		if lch == lcn {
			ov := v
			if lch == "authorization" {
				ov = "*************"
			}
			log.Infof("found header %v = %v", name, ov)
			value = v
			break
		}
	}

	return value
}

// GetMapItem Retrieve an item from a string map.
func GetMapItem(mapData StringMap, name string) string {
	value := ""
	ln := strings.ToLower(name)

	for k, v := range mapData {
		lk := strings.ToLower(k)
		if lk == ln {
			log.Infof("found item %q in string map", name)
			value = v
			break
		}
	}

	return value
}

// Respond200 Send a status OK HTTP response.
func Respond200(content []byte, contentType string) *Response {
	code := http.StatusOK
	res := &Response{
		Headers: StringMap{
			"Content-Type": contentType,
		},
		StatusCode: code,
		Cookies:    []string{},
	}

	switch contentType {
	case ContentTypeGif, ContentTypeJpg, ContentTypePng:
		res.Body = base64.StdEncoding.EncodeToString(content)
		res.IsBase64Encoded = true
	default:
		res.Body = string(content)
	}

	return res
}

// Respond301Or308 Send a 301 or 308 HTTP response redirect to another location.
// Deprecated see Respond301 or Respond308
func Respond301Or308(method, location string) *Response {
	code := http.StatusMovedPermanently
	content := Http301RedirectContent

	if method == "POST" {
		code = http.StatusPermanentRedirect
		content = Http308RedirectContent
	}

	if !strings.Contains(location, "https://") {
		location = "https://" + location
	}

	return &Response{
		Body: content,
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
			"Location":     location,
		},
		StatusCode: code,
	}
}

const Footer = "Acme"

// Respond301 Send a 301 HTTP response redirect to another location (full URL).
func Respond301(location string) *Response {
	code := http.StatusMovedPermanently
	return &Response{
		Body: fmt.Sprintf(HttpRedirectContent, code, "Moved Permanently", Footer),
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
			"Location":     location,
		},
		StatusCode: code,
	}
}

// Respond302 Send a 302 HTTP response redirect to another location (full URL).
func Respond302(location string) *Response {
	code := http.StatusFound
	return &Response{
		Body: fmt.Sprintf(HttpRedirectContent, code, "Found", Footer),
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
			"Location":     location,
		},
		StatusCode: code,
	}
}

// Respond308 Send a 308 HTTP response redirect to another location (full URL).
func Respond308(location string) *Response {
	code := http.StatusPermanentRedirect
	return &Response{
		Body: fmt.Sprintf(HttpRedirectContent, code, "Permanent Redirect", Footer),
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
			"Location":     location,
		},
		StatusCode: code,
	}
}

// Respond401 Send a 401 Unauthorized HTTP response.
func Respond401() *Response {
	code := http.StatusUnauthorized
	return &Response{
		Body: Http401UnauthorizedContent,
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
		},
		StatusCode: code,
	}
}

// Respond404 Send a 404 Not Found HTTP response.
func Respond404() *Response {
	code := http.StatusNotFound
	return &Response{
		Body: Http404NotFoundContent,
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
		},
		StatusCode: code,
	}
}

// Respond405 Send a 405 Method Not Allowed HTTP response.
//
//	allowedMethods is a comma delimited string of HTTP methods that are allowed.
//	Example:
//	  GET, HEAD, PUT
func Respond405(allowedMethods string) *Response {
	code := http.StatusMethodNotAllowed
	return &Response{
		Body: fmt.Sprintf(HttpStatusContent, code, http.StatusText(code), Footer),
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
			"Allow":        allowedMethods,
		},
		StatusCode: code,
	}
}

// Respond500 Send a 500 Internal Server Error HTTP response.
func Respond500() *Response {
	code := http.StatusInternalServerError
	return &Response{
		Body: Http500InternalErrorContent,
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
		},
		StatusCode: code,
	}
}

// Respond501 Send a 501 Not Implemented HTTP response.
//
//	501 is the appropriate response when the server does not recognize the
//	request method and is incapable of supporting it for any resource. The only
//	methods that servers are required to support (and therefore that must not
//	return 501) are GET and HEAD.
func Respond501() *Response {
	code := http.StatusNotImplemented
	return &Response{
		Body: Http501NotImplemented,
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
		},
		StatusCode: code,
	}
}

// RespondDebug Respond with a debug message and whatever code your like.
//
//	This was handy when testing AWS Lambda function or initial set up of the
//	Lambda URL feature.
func RespondDebug(code int, message, footer string) *Response {
	return &Response{
		Body: fmt.Sprintf(Http200Debug, code, message, footer),
		Headers: StringMap{
			"Content-Type": ContentTypeHtml,
		},
		StatusCode: code,
	}
}

// RespondJSON Send a JSON HTTP response.
func RespondJSON(content interface{}) (*Response, error) {
	jsonEncodedContent, e1 := json.Marshal(content)
	if e1 != nil {
		return nil, fmt.Errorf(Stderr.CannotEncodeToJson, e1.Error())
	}

	return Respond200(jsonEncodedContent, ContentTypeJson), nil
}

// ResponseOptions Respond with an HTTP Allow header listing all HTTP methods
// allowed for a request.
func ResponseOptions(options string) *Response {
	return &Response{
		Body: "",
		Headers: StringMap{
			"Allow": options,
		},
		StatusCode: 204,
	}
}
