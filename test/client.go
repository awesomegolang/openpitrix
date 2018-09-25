package test

import (
	"context"
	"log"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	flag "github.com/spf13/pflag"

	apiclient "openpitrix.io/openpitrix/test/client"
	"openpitrix.io/openpitrix/test/config"
)

const UserSystem = "system"

type IgnoreLogger struct{}

func (IgnoreLogger) Printf(format string, args ...interface{}) {
}

func (IgnoreLogger) Debugf(format string, args ...interface{}) {
}

type ClientConfig struct {
	Debug      bool
	ConfigPath string
}

func (conf *ClientConfig) GetEndpoint() *url.URL {
	configFile, err := config.ReadConfigFile(context.Background(), conf.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	endpoint, err := url.Parse(configFile.EndpointURL)
	if err != nil {
		log.Fatal(err)
	}
	return endpoint
}

func GetClient(conf *ClientConfig) *apiclient.Openpitrix {
	endpoint := conf.GetEndpoint()
	client, err := config.GetClient(context.Background(), conf.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	transport := httptransport.NewWithClient(endpoint.Host, "/", []string{endpoint.Scheme}, client)
	transport.SetDebug(conf.Debug)
	//transport.SetLogger(IgnoreLogger{})
	Client := apiclient.New(transport, strfmt.Default)
	return Client
}

func GetClientConfig() *ClientConfig {
	var conf = ClientConfig{}
	config.AddFlag(flag.CommandLine, &conf.ConfigPath)
	flag.Parse()
	return &conf
}
