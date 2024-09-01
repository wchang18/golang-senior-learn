package cobra_learn

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear log",
	Long:  `clear log`,
	Run: func(cmd *cobra.Command, args []string) {
		clearLog(delhour)
	},
}

var delhour int

func init() {
	clearCmd.Flags().IntVarP(&delhour, "delhour", "d", 24, "delete time")
}

func clearLog(hour int) {
	logPath := viper.GetString("LogPath")
	deltime := time.Now().Add(-1 * time.Duration(hour) * time.Hour)
	filepath.Walk(logPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.ModTime().Before(deltime) && strings.HasSuffix(info.Name(), ".log") {
			err = os.Remove(path)
			if err != nil {
				log.Println("delete file error:", err)
			}
			log.Println("delete file:", path)
		}

		return nil
	})
}
