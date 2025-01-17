package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io"
	"os"
	"runtime"
	"time"
)

// 检测数据存储位置
var savepath = flag.String("o", "./data.csv", "数据存储位置")

// 检测进程ID
var pid = flag.Int("p", 0, "进程ID")

// 采样间隔
var intv = flag.Int("i", 1000, "采样间隔，单位毫秒")

var v = flag.Bool("v", false, fmt.Sprintf("版本信息 %s", VERSION))

const VERSION = "1.0.0"

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		_, _ = fmt.Fprintf(os.Stderr, "  plog -p 745 -o ./data.csv\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *v {
		_, _ = fmt.Fprintf(os.Stderr, "plog version %s\n", VERSION)
		return
	}

	if *pid <= 0 {
		_, _ = fmt.Fprintf(os.Stderr, "PID 错误\n")
		return
	}
	if *savepath == "" {
		_, _ = fmt.Fprintf(os.Stderr, "数据存储位置错误\n")
		return
	}
	if *intv <= 0 {
		*intv = 1000
	}

	file, _ := os.Create(*savepath)
	defer file.Close()
	out := io.MultiWriter(os.Stdout, file)

	fmt.Printf("开始监控进程 %d\n", *pid)
	p, err := process.NewProcess(int32(*pid))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "PID [%s] 不存在\n", pid)
		return
	}
	_, _ = fmt.Fprintf(out, "ts, CPU(%%), Mem(KB)\n")
	for {
		time.Sleep(time.Duration(*intv) * time.Millisecond)
		cpuUsage, err := p.CPUPercent()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "CPU Usage ERR %v\n", err)
			return
		}
		memUsage, err := p.MemoryInfo()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Memory Usage ERR %v\n", err)
			return
		}
		_, _ = fmt.Fprintf(out, "%d, %.2f, %d\n", time.Now().UnixMilli(), cpuUsage/float64(runtime.NumCPU()), memUsage.RSS/1024)
	}

}
