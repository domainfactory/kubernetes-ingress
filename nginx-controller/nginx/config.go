package nginx

// Config holds NGINX configuration parameters
type Config struct {
	ProxyConnectTimeout           string
	ProxyReadTimeout              string
	ClientMaxBodySize             string
	MainServerNamesHashBucketSize string
	MainServerNamesHashMaxSize    string
	MainLogFormat                 string
	MainProxyBuffers              string
	MainProxyBufferSize           string
}

// NewDefaultConfig creates a Config with default values
func NewDefaultConfig() *Config {
	return &Config{
		ProxyConnectTimeout:        "60s",
		ProxyReadTimeout:           "60s",
		ClientMaxBodySize:          "1m",
		MainServerNamesHashMaxSize: "512",
                MainLogFormat:              "",
                MainProxyBuffers:           "",
                MainProxyBufferSize:        "",
	}
}
