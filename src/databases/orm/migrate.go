package orm

import (
	"log"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "For migrating database",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", true, "for running auto migration")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "for running auto reset migration")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := NewDb()
	if err != nil {
		return err
	}

	if migDown {
		log.Println("migration reset successfully !")
		return db.Migrator().DropTable(&models.User{}, &models.Category{}, &models.Vehicle{}, &models.Reservation{}, &models.History{})
	}

	if migUp {

		db.AutoMigrate(&models.User{}, &models.Category{}, &models.Vehicle{}, &models.Reservation{}, &models.History{})
		log.Println("Migration successfully !")

	}

	return nil
}
