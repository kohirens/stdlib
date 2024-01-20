package web

const (
	Http301RedirectContent      = `<!DOCTYPE html><html><head><title>301 Moved Permanently</title></head><body><center><h1>301 Moved Permanently</h1></center><hr><center>CloudFront</center></body></html>`
	Http308RedirectContent      = `<!DOCTYPE html><html><head><title>308 Permanent Redirect</title></head><body><center><h1>308 Permanent Redirect</h1></center><hr><center>CloudFront</center></body></html>`
	Http401UnauthorizedContent  = `<!DOCTYPE html><html><head><title>401 Unauthorized</title></head><body><center><h1>401 Unauthorized</h1></center><hr><center>CloudFront</center></body></html>`
	Http404NotFoundContent      = `<!DOCTYPE html><html><head><title>404 Not Found</title></head><body><center><h1>404 Not Found</h1></center><hr><center>CloudFront</center></body></html>`
	Http500InternalErrorContent = `<!DOCTYPE html><html><head><title>500 Internal Server Error</title></head><body><center><h1>500 Internal Server Error</h1></center><hr><center>CloudFront</center></body></html>`
	Http501NotImplemented       = `<!DOCTYPE html><html><head><title>501 Not Implemented</title></head><body><center><h1>501 Not Implemented</h1></center><hr><center>CloudFront</center></body></html>`

	// See [Media Types](https://www.iana.org/assignments/media-types/media-types.xhtml)
	// Also see [IETF Media Types](https://www.rfc-editor.org/rfc/rfc9110.html#media.type)

	ContentTypeCSS  = "text/css;charset=utf-8"
	ContentTypeGif  = "image/gif;charset=utf-8"
	ContentTypeHtml = "text/html;charset=utf-8"
	ContentTypeJpg  = "image/jpeg;charset=utf-8"
	ContentTypeJS   = "text/javascript;charset=utf-8"
	ContentTypeJson = "application/json;charset=utf-8"
	ContentTypePng  = "image/png;charset=utf-8"
	ContentTypeSvg  = "image/svg+xml;charset=utf-8"
)
