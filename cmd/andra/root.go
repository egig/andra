package andra

import (
	"context"
	"fmt"
	"github.com/egig/andra/cmd/andra/assets"
	"github.com/egig/andra/pkg/server"
	"github.com/spf13/cobra"
	"os"
)

const ConfigDefaultFileName = "andra.config.js"
const ConfigDefaultEnvName = "production"
const ConfigDefaultPort = ":8080"

var (
	projDir string
	cfgFile string
	env     string
	dev     bool
	dt      *cobra.Command
	port    string
)

func init() {
	dt = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.WithValue(context.Background(), "port", port)
			ctx = context.WithValue(ctx, "dir", projDir)
			server.Run(ctx)
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dt.PersistentFlags().StringVar(&projDir, "dir", wd, "config file")
	dt.PersistentFlags().StringVar(&cfgFile, "config", ConfigDefaultFileName, "preject directory")
	dt.PersistentFlags().StringVar(&env, "env", ConfigDefaultEnvName, "env name")
	dt.PersistentFlags().BoolVar(&dev, "dev", false, "alias for --env=develop")
	dt.PersistentFlags().StringVar(&port, "port", ConfigDefaultPort, "port")

	dt.AddCommand(assets.RootCmd())
}

func Execute() {
	if err := dt.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
