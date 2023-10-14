package mansion

import (
	"fmt"
	"github.com/itchio/butler/comm"
	"github.com/itchio/httpkit/timeout"
	"github.com/itchio/wharf/pwr"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

type DoCommand func(ctx *Context)

type Context struct {
	App      *kingpin.Application
	Commands map[string]DoCommand

	// Identity is the path to the credentials file
	Identity string

	// String to include in our user-agent
	UserAgentAddition string

	// Quiet silences all output
	Quiet bool

	// Verbose enables chatty output
	Verbose bool

	// Verbose enables JSON output
	JSON bool

	// Path to the local sqlite database
	DBPath string

	CompressionAlgorithm string
	CompressionQuality   int

	ContextTimeout int64

	HTTPClient    *http.Client
	HTTPTransport *http.Transport

	// url of the itch.io API server we're talking to
	apiAddress string
	// url of the itch.io web instance we're talking to
	webAddress string
}

func NewContext(app *kingpin.Application) *Context {
	client := timeout.NewDefaultClient()
	originalTransport := client.Transport.(*http.Transport)

	ctx := &Context{
		App:           app,
		Commands:      make(map[string]DoCommand),
		HTTPClient:    client,
		HTTPTransport: originalTransport,
	}

	client.Transport = originalTransport
	return ctx
}

func (ctx *Context) Register(clause *kingpin.CmdClause, do DoCommand) {
	ctx.Commands[clause.FullCommand()] = do
}

func (ctx *Context) Must(err error) {
	if err != nil {
		if ctx.Verbose || ctx.JSON {
			comm.Dief("%+v", err)
		} else {
			comm.Dief("%s", err)
		}
	}
}

func (ctx *Context) CompressionSettings() pwr.CompressionSettings {
	var algo pwr.CompressionAlgorithm

	switch ctx.CompressionAlgorithm {
	case "none":
		algo = pwr.CompressionAlgorithm_NONE
	case "brotli":
		algo = pwr.CompressionAlgorithm_BROTLI
	default:
		panic(fmt.Errorf("Unknown compression algorithm: %s", algo))
	}

	return pwr.CompressionSettings{
		Algorithm: algo,
		Quality:   int32(ctx.CompressionQuality),
	}
}
