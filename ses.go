package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var VERSION = "0.0.1"
var EPOCH_ = "20211015"

func main() {
	// TODO: SetProcTitle
	app := &cli.App{
		Name:  "ses",
		Usage: "Google/Bing/Baidu/360 from command-line",
		Action: func(c *cli.Context) error {
			fmt.Println("ses!")
			return nil
		},
	}
	app.Version = VERSION
	app.Authors = []*cli.Author{{
		Name:  "Tacey Wong",
		Email: "xinyong.wang@qq.com",
	}}
	app.Copyright = fmt.Sprintf("© 2021-Now %s(%s)", app.Authors[0].Name, app.Authors[0].Email)

	app.Flags = []cli.Flag{
		&cli.Int64Flag{Name: "s", Aliases: []string{"start"}, Usage: "Start at the Nth result"},
		&cli.Int64Flag{Name: "n", Aliases: []string{"count"}, Usage: "Show N results", Value: 10},
		&cli.BoolFlag{Name: "N", Aliases: []string{"news"}, Usage: "Show resultsfrom news section"},
		&cli.BoolFlag{Name: "V", Aliases: []string{"videos"}, Usage: "Show results from videos section"},
		&cli.StringFlag{Name: "c", Aliases: []string{"country"}, Usage: "country-specific search with top-level domain .TLD, e.g., 'in' for India "},
		&cli.StringFlag{Name: "l", Aliases: []string{"lang"}, Usage: "display in language LANG"},
		&cli.StringFlag{Name: "g", Aliases: []string{"geoloc"}, Usage: "country-specific geolocation search with country code CC, e.g. 'in' for India. Country codes are the same as top-level domains"},
		&cli.BoolFlag{Name: "x", Aliases: []string{"exact"}, Usage: "disable automatic spelling correction"},
		&cli.StringFlag{Name: "CC", Aliases: []string{"colorize"}, Usage: "whether to colorize output; defaults to 'auto',which enables color when stdout is a tty device.using --colorize without an argument is equivalent to --colorize=always"},
		&cli.BoolFlag{Name: "C", Aliases: []string{"nocolor"}, Usage: "equivalent to --colorize=never"},
		&cli.StringFlag{Name: "CS", Aliases: []string{"colors"}, Usage: " set output colors (see man page for details)"},
		&cli.BoolFlag{Name: "j", Aliases: []string{"first", "lucky"}, Usage: "open the first result in web browser and exit"},
		&cli.StringFlag{Name: "tl", Aliases: []string{"time"}, Usage: "time limit search [h5 (5 hrs), d5 (5 days), w5 (5 weeks), m5 (5 months), y5 (5 years)]"},
		&cli.StringFlag{Name: "f", Aliases: []string{"from"}, Usage: "starting date/month/year of date range; must use American date format with slashes, e.g.,2/24/2020, 2/2020, 2020; can be used in conjunction with --to, and overrides -t, --time "},
		&cli.StringFlag{Name: "t", Aliases: []string{"to"}, Usage: "ending date/month/year of date range; see --from"},
		&cli.StringFlag{Name: "w", Aliases: []string{"site"}, Usage: "search a site using Google"},
		&cli.StringFlag{Name: "e", Aliases: []string{"exclude"}, Usage: "exclude site from results"},
		&cli.BoolFlag{Name: "uf", Aliases: []string{"unfilter"}, Usage: "do not omit similar results"},
		&cli.StringFlag{Name: "p", Aliases: []string{"proxy"}, Usage: "tunnel traffic through an HTTP proxy; PROXY is of the form [http://][user:password@]proxyhost[:port]"},
		&cli.BoolFlag{Name: "nt", Aliases: []string{"notweak"}, Usage: "disable TCP optimizations and forced TLS 1.2"},
		&cli.BoolFlag{Name: "js", Aliases: []string{"json"}, Usage: "output in JSON format; implies --noprompt"},
		&cli.StringFlag{Name: "uh", Aliases: []string{"url-handler"}, Usage: "custom script or cli utility to open results"},
		&cli.BoolFlag{Name: "sbl", Aliases: []string{"show-browser-logs"}, Usage: "do not suppress browser output (stdout and stderr)"},
		&cli.BoolFlag{Name: "np", Aliases: []string{"noprompt"}, Usage: "search and exit, do not prompt"},
		&cli.BoolFlag{Name: "4", Aliases: []string{"ipv4"}, Usage: "only connect over IPv4 (by default, IPv4 is preferred but IPv6 is used as a fallback"},
		&cli.BoolFlag{Name: "6", Aliases: []string{"ipv6"}, Usage: "only connect over IPv6"},
		&cli.BoolFlag{Name: "d", Aliases: []string{"debug"}, Usage: " enable debugging"},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
