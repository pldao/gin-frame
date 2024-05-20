package command

import (
	"fmt"
	"github.com/PLDao/gin-frame/data"
	"github.com/PLDao/gin-frame/internal/console/demo"
	"github.com/PLDao/gin-frame/internal/routers"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "command",
		Short:   "The control head runs the command",
		Example: "go-layout command demo",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 初始化数据库
			data.InitData()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	host = "0.0.0.0"
	port = 9001
)

func init() {
	Cmd.AddCommand(demo.DemoCmd)
}

func run() error {
	r := routers.SetRouters()

	err := r.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	return nil
}
