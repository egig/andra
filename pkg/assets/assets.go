package assets

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func Build(dir string) string {
	e := fmt.Sprintf("%s/%s", dir, "web/app/admin/src/entry.tsx")
	o := fmt.Sprintf("%s/%s", dir, "assets")
	options := api.BuildOptions{
		EntryPoints: []string{e},
		//Outbase: ,
		Outdir: o,
		Bundle: true,
		//MinifyWhitespace:  true,
		//MinifyIdentifiers: true,
		//MinifySyntax:      true,
		Write: true,
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
		JSX:        api.JSXAutomatic,
		// Might want to set this for production build
		// PublicPath:
	}

	// TODO watch is not writing file
	//if watch {
	//	options.Watch = &api.WatchMode{
	//		OnRebuild: func(result api.BuildResult) {
	//			if len(result.Errors) > 0 {
	//				fmt.Printf("watch build failed: %d errors\n", len(result.Errors))
	//			} else {
	//				fmt.Printf("watch build succeeded: %d warnings\n", len(result.Warnings))
	//			}
	//		},
	//	}
	//}

	result := api.Build(options)
	if len(result.Errors) > 0 {
		log.Println(result.Errors)
		os.Exit(1)
	}
	err := ioutil.WriteFile(path.Join(o, "meta.json"), []byte(result.Metafile), 0644)
	if err != nil {
		panic(err)
	}

	return result.Metafile

	//if watch {
	//	<-make(chan bool)
	//}
}
