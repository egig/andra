package assets

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var watch bool

var buildCmd = &cobra.Command{
	Use: "build",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			log.Println(err)
		}

		e := fmt.Sprintf("%s/%s", dir, "pkg/admin/web/app/admin/src/entry.tsx")
		o := fmt.Sprintf("%s/%s", dir, "assets")
		options := api.BuildOptions{
			EntryPoints: []string{e},
			//Outbase: ,
			Outdir:            o,
			Bundle:            true,
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			MinifySyntax:      true,
			Write:             true,
			Loader: map[string]api.Loader{
				".eot":  api.LoaderFile,
				".ttf":  api.LoaderFile,
				".woff": api.LoaderFile,
				".svg":  api.LoaderFile,
				".png":  api.LoaderFile,
				".jpg":  api.LoaderFile,
				".jpeg": api.LoaderFile,
			},
			AssetNames: "[name]-[hash]",
			EntryNames: "[dir]/[name]-[hash]",
			Metafile:   true,
			// Might want to set this for production build
			// PublicPath:
		}

		// TODO watch is not writing file
		if watch {
			options.Watch = &api.WatchMode{
				OnRebuild: func(result api.BuildResult) {
					if len(result.Errors) > 0 {
						fmt.Printf("watch build failed: %d errors\n", len(result.Errors))
					} else {
						fmt.Printf("watch build succeeded: %d warnings\n", len(result.Warnings))
					}
				},
			}
		}

		result := api.Build(options)
		if len(result.Errors) > 0 {
			log.Println(result.Errors)
			os.Exit(1)
		}
		fmt.Println(result.Metafile)
		err = ioutil.WriteFile(path.Join(o, "meta.json"), []byte(result.Metafile), 0644)
		if err != nil {
			panic(err)
		}

		if watch {
			<-make(chan bool)
		}
	},
}

func init() {
	buildCmd.PersistentFlags().BoolVar(&watch, "watch", false, "watch file change")
}
