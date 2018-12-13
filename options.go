package kubemq

type Option interface {
	apply(*Options)
}

type Options struct {
	host                 string
	port                 int
	isSecured            bool
	certFile             string
	serverOverrideDomain string
	token                string
	clientId             string
	receiveBufferSize    int
}

type funcOptions struct {
	fn func(*Options)
}

func (fo *funcOptions) apply(o *Options) {
	fo.fn(o)
}

func newFuncOption(f func(*Options)) *funcOptions {
	return &funcOptions{
		fn: f,
	}
}

// WithAddress - set host and port address of KubeMQ server
func WithAddress(host string, port int) Option {
	return newFuncOption(func(o *Options) {
		o.host = host
		o.port = port
	})
}

// WithCredentials - set secured TLS credentials from the input certificate file for client.
// serverNameOverride is for testing only. If set to a non empty string,
// it will override the virtual host name of authority (e.g. :authority header field) in requests.
func WithCredentials(certFile, serverOverrideDomain string) Option {
	return newFuncOption(func(o *Options) {
		o.isSecured = true
		o.certFile = certFile
		o.serverOverrideDomain = serverOverrideDomain
	})
}

// WithToken - set KubeMQ token to be used for KubeMQ connection - not mandatory, only if enforced by the kubemq server
func WithToken(token string) Option {
	return newFuncOption(func(o *Options) {
		o.token = token
	})
}

// WithClientId - set client id to be used in all functions call with this client - mandatory
func WithClientId(id string) Option {
	return newFuncOption(func(o *Options) {
		o.clientId = id
	})
}

// WithReceiveBufferSize - set length of buffered channel to be set in all subscriptions
func WithReceiveBufferSize(size int) Option {
	return newFuncOption(func(o *Options) {
		o.receiveBufferSize = size
	})
}

func GetDefaultOptions() *Options {
	return &Options{
		host:                 "localhost",
		port:                 50000,
		isSecured:            false,
		certFile:             "",
		serverOverrideDomain: "",
		token:                "",
		clientId:             "clientId",
		receiveBufferSize:    10,
	}
}
