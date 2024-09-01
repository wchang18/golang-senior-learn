package cobra_learn

import (
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

var checkInfoCmd = &cobra.Command{
	Use:     "check_info",
	Short:   "check host info",
	Long:    `check host info`,
	Example: "check_info",
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
	checkInfoCmd.Flags().StringVarP(&fileType, "file_type", "t", "txt", "文件类型")
	checkInfoCmd.Flags().StringVarP(&filePath, "file_path", "p", "./", "文件路径")
	checkInfoCmd.Flags().StringVarP(&fileName, "file_name", "n", "info.txt", "文件名称")
}

type OsInfo struct {
	HostInfo   *host.InfoStat
	CpuPercent []float64
	MemStats   *mem.VirtualMemoryStat
	DiskPart   []disk.PartitionStat
	DiskUsage  *disk.UsageStat
}

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

var header = []string{"DateTime", "HostName", "CPU使用比", "内存使用比", "磁盘使用比"}

func WriteTxtFile(fileType string, filePath string, fileName string, info OsInfo) {
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

func CheckInfo() {
	info := GetOsInfo()
	if !strings.HasSuffix(fileName, fileType) {
		fileName = fileName + "." + fileType
	}
	if fileType == "txt" {
		WriteTxtFile(fileType, filePath, fileName, info)
	}
}
