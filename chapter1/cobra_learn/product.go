package cobra_learn

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var productCmd = &cobra.Command{
	Use:     "product [name]",
	Short:   "product data",
	Long:    `Manage data products`,
	Example: `product report --date=20240701`,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Print("product start")
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		productData(args[0], Date, Count)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Print("product end")
	},
}

var Count int

func init() {
	productCmd.Flags().IntVarP(&Count, "count", "c", 1, "product count")
}

func productData(name, data string, count int) {
	path := viper.GetString("DataPath")
	timestamp := time.Now().UnixMilli()
	for i := 1; i <= count; i++ {
		filename := fmt.Sprintf("%s/%s_%s_%d_%d", path, name, data, i, timestamp)
		file, err := os.Create(filename)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("create file:", filename)
	}
}
