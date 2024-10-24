package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var FileInfoCmd = &cobra.Command{
	Use:     "file_info",
	Short:   "获取文件或文件夹信息",
	Long:    "文件或文件夹信息",
	Example: "file_info --path=./",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			count int64
			size  int64
		)
		GetPathInfo(FilePath, &count, &size)
		fmt.Printf("文件个数：%d，文件大小：%d b\n", count, size)
	},
}

var FilePath string

func init() {
	FileInfoCmd.PersistentFlags().StringVarP(&FilePath, "path", "p", "", "文件或文件夹路径")
	RootCmd.AddCommand(FileInfoCmd)
}

func GetPathInfo(file string, count, size *int64) {
	// 获取文件或文件夹信息
	pathInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 判断是否为文件夹
	if !pathInfo.IsDir() {
		*size = pathInfo.Size()
		*count++
		return
	}
	// 读取文件夹下的文件
	list, err := os.ReadDir(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range list {
		info, _ := v.Info()
		*size += info.Size()
		if info.IsDir() {
			GetPathInfo(file+"/"+v.Name(), count, size)
		} else {
			*count++
		}
	}
}
