package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var CheckSys = &cobra.Command{
	Use:     "check_sys",
	Short:   "检查系统信息",
	Long:    "检查系统信息",
	Example: "check_sys --file_path=./data --file_type=csv --file_name=os-info",
	Run: func(cmd *cobra.Command, args []string) {
		CheckInfo()
	},
}

var (
	fileType string
	filePath string
	fileName string
)

func init() {
	RootCmd.AddCommand(CheckSys)
	CheckSys.Flags().StringVarP(&fileType, "file_type", "t", "txt", "文件类型")
	CheckSys.Flags().StringVarP(&filePath, "file_path", "p", "./", "文件路径")
	CheckSys.Flags().StringVarP(&fileName, "file_name", "n", "os_info", "文件名称")
}

type OsInfo struct {
	HostInfo   *host.InfoStat
	CpuPercent []float64
	MemStats   *mem.VirtualMemoryStat
	DiskPart   []disk.PartitionStat
	DiskUsage  *disk.UsageStat
}

var header = []string{"DateTime", "HostName", "CPU使用比", "内存使用比", "磁盘使用比"}

func GetOsInfo() OsInfo {
	var osInfo OsInfo
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println(err)
	}
	osInfo.CpuPercent = cpuPercent
	//fmt.Println(cpuPercent, err)

	memStats, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	osInfo.MemStats = memStats
	//fmt.Println(memStats, err)

	diskPart, err := disk.Partitions(true)
	if err != nil {
		fmt.Println(err)
	}
	osInfo.DiskPart = diskPart
	//fmt.Println(diskPart, err)

	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}
	osInfo.DiskUsage = diskUsage
	//fmt.Println(diskUsage, err)

	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	osInfo.HostInfo = hostInfo
	return osInfo
}

func WriteTxtFile(filePath string, fileName string, info OsInfo) {
	var content string
	content += fmt.Sprintf("%s:%v\n", header[0], time.Now().Format("2006-01-02 15:04:05"))
	content += fmt.Sprintf("%s:%v\n", header[1], info.HostInfo.Hostname)
	content += fmt.Sprintf("%s:%.2f\n", header[2], info.CpuPercent[0])
	content += fmt.Sprintf("%s:%v\n", header[3], info.MemStats.UsedPercent)
	content += fmt.Sprintf("%s:%.2f\n", header[4], info.DiskUsage.UsedPercent)
	file, err := os.Create(filePath + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(content)
}

func WriteCsvFile(filePath string, fileName string, info OsInfo) {
	fileName = filePath + fileName
	var (
		file   *os.File
		writer *csv.Writer
	)
	if _, err := os.Stat(fileName); err != nil {
		file, _ = os.Create(fileName)
		writer = csv.NewWriter(file)
		err = writer.Write(header)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		file, _ = os.OpenFile(fileName, os.O_APPEND, 0666)
		writer = csv.NewWriter(file)
	}
	defer writer.Flush()
	list := []string{time.Now().Format("2006-01-02 15:04:05"), info.HostInfo.Hostname, fmt.Sprintf("%.2f", info.CpuPercent[0]), fmt.Sprintf("%.2f", info.MemStats.UsedPercent), fmt.Sprintf("%.2f", info.DiskUsage.UsedPercent)}
	err := writer.Write(list)
	if err != nil {
		fmt.Println(err)
	}
}

func CheckInfo() {
	info := GetOsInfo()
	if !strings.HasSuffix(fileName, fileType) {
		fileName = fileName + "." + fileType
	}
	if fileType == "txt" {
		WriteTxtFile(filePath, fileName, info)
		return
	}

	if fileType == "csv" {
		WriteCsvFile(filePath, fileName, info)
		return
	}
}
