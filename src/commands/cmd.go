package commands

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm"
	"github.com/AnggaArdhinata/backend-go/src/databases/server"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "vehicle rental back-end golang",
	Long:  "golang back-end vehicle rental with gorilla/mux",
}

func init() {
	initCommand.AddCommand(server.ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
