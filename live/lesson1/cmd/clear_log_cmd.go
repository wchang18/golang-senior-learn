package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

var ClearLogCmd = &cobra.Command{
	Use:     "clear_log",
	Short:   "清理日志",
	Example: "clear_log --dir=./data --date=20241024 --day=3",
	Run: func(cmd *cobra.Command, args []string) {
		if Date != "" {
			if !CheckDate() {
				fmt.Printf("%s时间格式错误，例子：2024-10-01", Date)
				return
			}
		}
		ClearLogFunc()
	},
}

var (
	DirPath string
	Date    string
	Day     int
)

func init() {
	ClearLogCmd.Flags().StringVarP(&DirPath, "dir", "d", "", "指定目录")
	ClearLogCmd.Flags().StringVarP(&Date, "date", "t", "", "指定日期")
	ClearLogCmd.Flags().IntVarP(&Day, "day", "n", 0, "指定多少天之前")
	RootCmd.AddCommand(ClearLogCmd)
}

func ClearLogFunc() {

	pathInfo, err := os.Stat(DirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !pathInfo.IsDir() {
		fmt.Println("path is not dir")
		return
	}

	count := 0
	total := 0

	filepath.Walk(DirPath, func(path string, info os.FileInfo, err error) error {
		fileTime := info.ModTime()
		fmt.Println(path)
		total++
		if Date != "" {
			date, _ := time.Parse("2006-01-02", Date)
			if fileTime.Before(date) {
				err = os.Remove(path)
				if err == nil {
					count++
				}
			}
			return nil
		}

		if Day > 0 {
			if fileTime.Before(time.Now().AddDate(0, 0, -Day)) {
				err = os.Remove(path)
				if err == nil {
					count++
				}
				return nil
			} else {
				return nil
			}
		}

		err = os.Remove(path)
		if err == nil {
			count++
		}
		return nil
	})
	tmp, _ := filepath.Abs(DirPath)
	fmt.Printf("路径：%s，清理完成，共查询%d个文件，共%d个文件被清理\n", tmp, total, count)
}

func CheckDate() bool {
	_, err := time.Parse("2006-01-02", Date)
	return err == nil
}
