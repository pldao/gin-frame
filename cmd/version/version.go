package version

import (
	"fmt"
	"github.com/PLDao/gin-frame/internal/global"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "version",
		Short:   "GetUserInfo version info",
		Example: "go-layout version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(global.Version)
		},
	}
)
