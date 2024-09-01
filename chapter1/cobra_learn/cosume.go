package cobra_learn

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "consume data",
	Long:  `consume data`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			consumeData()
			time.Sleep(time.Millisecond * 100)
		}
	},
}

func consumeData() {
	root := viper.GetString("DataPath")
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		log.Println(info.Name() + " consume success")
		err = os.Remove(path)
		if err != nil {
			log.Println(info.Name() + " consume fail")
		}
		return nil
	})
}
