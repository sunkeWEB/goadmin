package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"go-admin/common/global"
)

var (
	configYml string
	port      string
	mode      string
	StartCmd  = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "go-admin version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
			return run()
		},
	}
)

func run() error {
	fmt.Printf("版本号：%s\n", global.Version)
	return nil
}
