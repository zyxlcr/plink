package router

// HTTP Header keys
const (
	Age                           string = "Age"
	AltSCV                        string = "Alt-Svc"
	Accept                        string = "Accept"
	AcceptCharset                 string = "Accept-Charset"
	AcceptPatch                   string = "Accept-Patch"
	AcceptRanges                  string = "Accept-Ranges"
	AcceptedLanguage              string = "Accept-Language"
	AcceptEncoding                string = "Accept-Encoding"
	Authorization                 string = "Authorization"
	CrossOriginResourcePolicy     string = "Cross-Origin-Resource-Policy"
	CacheControl                  string = "Cache-Control"
	Connection2                   string = "Connection"
	ContentDisposition            string = "Content-Disposition"
	ContentEncoding               string = "Content-Encoding"
	ContentLength                 string = "Content-Length"
	ContentType                   string = "Content-Type"
	ContentLanguage               string = "Content-Language"
	ContentLocation               string = "Content-Location"
	ContentRange                  string = "Content-Range"
	Date                          string = "Date"
	DeltaBase                     string = "Delta-Base"
	ETag                          string = "ETag"
	Expires                       string = "Expires"
	Host                          string = "Host"
	IM                            string = "IM"
	IfMatch                       string = "If-Match"
	IfModifiedSince               string = "If-Modified-Since"
	IfNoneMatch                   string = "If-None-Match"
	IfRange                       string = "If-Range"
	IfUnmodifiedSince             string = "If-Unmodified-Since"
	KeepAlive                     string = "Keep-Alive"
	LastModified                  string = "Last-Modified"
	Link                          string = "Link"
	Pragma                        string = "Pragma"
	ProxyAuthenticate             string = "Proxy-Authenticate"
	ProxyAuthorization            string = "Proxy-Authorization"
	PublicKeyPins                 string = "Public-Key-Pins"
	RetryAfter                    string = "Retry-After"
	Referer                       string = "Referer"
	Server                        string = "Server"
	SetCookie                     string = "Set-Cookie"
	StrictTransportSecurity       string = "Strict-Transport-Security"
	Trailer                       string = "Trailer"
	TK                            string = "Tk"
	TransferEncoding              string = "Transfer-Encoding"
	Location                      string = "Location"
	Upgrade                       string = "Upgrade"
	Vary                          string = "Vary"
	Via                           string = "Via"
	Warning                       string = "Warning"
	WWWAuthenticate               string = "WWW-Authenticate"
	XForwardedFor                 string = "X-Forwarded-For"
	XForwardedHost                string = "X-Forwarded-Host"
	XForwardedProto               string = "X-Forwarded-Proto"
	XRealIP                       string = "X-Real-Ip"
	XContentTypeOptions           string = "X-Content-Type-Options"
	XFrameOptions                 string = "X-Frame-Options"
	XXSSProtection                string = "X-XSS-Protection"
	XDNSPrefetchControl           string = "X-DNS-Prefetch-Control"
	Allow                         string = "Allow"
	Origin                        string = "Origin"
	AccessControlAllowOrigin      string = "Access-Control-Allow-Origin"
	AccessControlAllowCredentials string = "Access-Control-Allow-Credentials"
	AccessControlAllowHeaders     string = "Access-Control-Allow-Headers"
	AccessControlAllowMethods     string = "Access-Control-Allow-Methods"
	AccessControlExposeHeaders    string = "Access-Control-Expose-Headers"
	AccessControlMaxAge           string = "Access-Control-Max-Age"
	AccessControlRequestHeaders   string = "Access-Control-Request-Headers"
	AccessControlRequestMethod    string = "Access-Control-Request-Method"
	TimingAllowOrigin             string = "Timing-Allow-Origin"
	UserAgent                     string = "User-Agent"

	Gzip     string = "gzip"
	Compress string = "compress"
	Deflate  string = "deflate"
	Br       string = "br"
	Identity string = "identity"
	Any      string = "*"
)

// HTTP Constant Terms and Variables
const (
	// use constants from github.com/go-playground/net/http

	WildcardParam = "*wildcard"
	BasePath      = "/"
	Blank         = ""
	SlashByte     = '/'
	ParamByte     = ':'
	WildByte      = '*'
)

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

var (
	defaultContextIdentifier = &struct {
		name string
	}{
		name: "pure",
	}
)
