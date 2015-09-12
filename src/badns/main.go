package main

import (
	"badns/glitch"

	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/miekg/dns"
)

const (
	applicationName        = "badns"
	applicationVersion     = "0.0.1"
	applicationDescription = "Mess up DNS responses that are otherwise perfectly valid."
)

var log = logrus.WithFields(logrus.Fields{"app": applicationName})

func initLogrus(ctx *cli.Context) {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC822,
		FullTimestamp:   true,
	})

	if level, err := logrus.ParseLevel(ctx.String("log_level")); err == nil {
		logrus.SetLevel(level)
	} else {
		log.Error(err)
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func start(ctx *cli.Context) {
	addr := fmt.Sprintf(":%d", ctx.Int("port"))
	net := "udp"
	if ctx.Bool("tcp") {
		net = "tcp"
	}
	handler := BaDNSHandler{hostPrefix: ctx.String("host-prefix")}

	enableGlitches(ctx, &handler)

	server := dns.Server{Addr: addr, Net: net, Handler: handler}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start the server", err)
	}
}

func enableGlitches(ctx *cli.Context, handler *BaDNSHandler) {
	g := glitch.Ttl{TTL: ctx.Int("ttl")}
	handler.AddGlitch(g)

	if ctx.Int("delay") > 0 {
		g := glitch.Delay{Duration: 3 * time.Second}
		handler.AddGlitch(g)
	}

	if ctx.Bool("no-answer") {
		g := glitch.NoAnswer{}
		handler.AddGlitch(g)
	}

	if ctx.String("replace-type") != "" {
		g := glitch.ReplaceType{Type: ctx.String("replace-type")}
		handler.AddGlitch(g)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = applicationName
	app.Version = applicationVersion
	app.Usage = applicationDescription
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 53,
			Usage: "Listen port(default=53)",
		},
		cli.BoolFlag{
			Name:  "tcp",
			Usage: "Listen on TCP",
		},
		cli.StringFlag{
			Name:  "log_level, l",
			Value: "info",
			Usage: "Logging level (debug, info, warn, error, fatal, panic) (default=info)",
		},
		cli.IntFlag{
			Name:  "ttl",
			Value: 1,
			Usage: "Force TTL to the given value for all responses (default=1)",
		},
		cli.StringFlag{
			Name:  "host-prefix",
			Value: "",
			Usage: "Only apply glitches if query hostname matches the prefix",
		},
		cli.IntFlag{
			Name:  "delay",
			Value: 0,
			Usage: "Delay DNS responses (default=0 seconds)",
		},
		cli.BoolFlag{
			Name:  "no-answer",
			Usage: "Remove answer section from DNS responses",
		},
		cli.StringFlag{
			Name:  "replace-type",
			Value: "",
			Usage: "Replace RR types with set value (i.e. MX) or define a mapping (i.e. A:AAAA)",
		},
	}
	app.Action = start
	app.Run(os.Args)

}
