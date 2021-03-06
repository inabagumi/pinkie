package main

import (
	"log"
	"os"
	"path/filepath"

	pinkie "github.com/inabagumi/pinkie/v4/pkg/client"
	"github.com/inabagumi/pinkie/v4/pkg/crawler"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version = "dev"

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "The Pinkie is a crawler that uses the YouTube Data API.")

	app.Version(version)
	app.VersionFlag.Short('v')

	app.HelpFlag.Short('h')

	all := app.Flag("all", "Fetch all videos of channel.").
		Short('a').
		Bool()

	channels := app.Flag("channel", "A channel ID to scrape.").
		Short('c').
		Required().
		PlaceHolder("<id>").
		Strings()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	opts := &crawler.Options{
		AlgoliaAPIKey:        os.Getenv("ALGOLIA_API_KEY"),
		AlgoliaApplicationID: os.Getenv("ALGOLIA_APPLICATION_ID"),
		AlgoliaIndexName:     os.Getenv("ALGOLIA_INDEX_NAME"),
		GoogleAPIKey:         os.Getenv("GOOGLE_API_KEY"),
	}

	c, err := pinkie.New(opts)
	if err != nil {
		log.Fatal(err)
	}

	c.Run(*channels, *all)
}
