package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type GhsOptions struct {
	Fields     string `short:"f"  long:"fields"     description:"limits what fields are searched. 'name', 'description', or 'readme'." default:"name,description"`
	Sort       string `short:"s"  long:"sort"       description:"The sort field. 'stars', 'forks', or 'updated'." default:"best match"`
	Order      string `short:"o"  long:"order"      description:"The sort order. 'asc' or 'desc'." default:"desc"`
	Language   string `short:"l"  long:"language"   description:"searches repositories based on the language they’re written in."`
	User       string `short:"u"  long:"user"       description:"limits searches to a specific user name."`
	Repository string `short:"r"  long:"repo"       description:"limits searches to a specific repository."`
	Max        int    `short:"m"  long:"max"        description:"limits number of result. range 1-1000" default:"100"`
	Version    bool   `short:"v"  long:"version"    description:"print version infomation and exit."`
	Enterprise string `short:"e"  long:"enterprise" description:"search from github enterprise."`
	Token      string `short:"t"  long:"token"      description:"Github API token to avoid Github API rate limit"`
}

func GhsOptionParser() ([]string, GhsOptions) {
	var opts GhsOptions
	parser := flags.NewParser(&opts, flags.HelpFlag)

	parser.Name = "ghs"
	parser.Usage = "[OPTION] \"QUERY\""
	args, err := parser.Parse()

	if err != nil {
		ghsOptionError(parser)
	}

	if opts.Version {
		fmt.Printf("ghs %s\n", Version)
		checkVersion(Version)
		os.Exit(0)
	}

	if (opts.User == "" && opts.Repository == "") && len(args) == 0 {
		ghsOptionError(parser)
	}

	if opts.Max < 1 || opts.Max > 1000 {
		ghsOptionError(parser)
	}

	return args, opts
}

func ghsOptionError(parser *flags.Parser) {
	printGhsHelp(parser)
	os.Exit(1)
}

func printGhsHelp(parser *flags.Parser) {
	parser.WriteHelp(os.Stdout)
	fmt.Printf("\n")
	fmt.Printf("Github search APIv3 QUERY infomation:\n")
	fmt.Printf("   https://developer.github.com/v3/search/\n")
	fmt.Printf("   https://help.github.com/articles/searching-repositories/\n")
	fmt.Printf("\n")
	fmt.Printf("Version:\n")
	fmt.Printf("   ghs %s (https://github.com/sona-tar/ghs.git)\n", Version)
	fmt.Printf("\n")
	checkVersion(Version)
}
