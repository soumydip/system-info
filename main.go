package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type commandRunner func(name string, args ...string) ([]byte, error)

func defaultCommandRunner(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}

func parseWMICOutput(output string) string {
	lines := strings.Split(output, "\n")
	cleaned := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}

	if len(cleaned) == 0 {
		return ""
	}
	if len(cleaned) == 1 {
		return cleaned[0]
	}
	return strings.Join(cleaned[1:], "; ")
}

func wmicInfo(runner commandRunner, args ...string) string {
	output, err := runner("wmic", args...)
	if err != nil {
		return "Unavailable"
	}

	result := parseWMICOutput(string(output))
	if result == "" {
		return "Unavailable"
	}
	return result
}

func formatBytes(raw string) string {
	raw = strings.TrimSpace(raw)
	bytes, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return raw
	}
	const gb = 1024 * 1024 * 1024
	return fmt.Sprintf("%.2f GB", float64(bytes)/gb)
}

func printRuntimeInfo() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unavailable"
	}

	fmt.Println("=== Runtime & OS Details ===")
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Memory Alloc: %.2f MB\n", float64(mem.Alloc)/(1024*1024))
	fmt.Printf("Memory Sys: %.2f MB\n", float64(mem.Sys)/(1024*1024))
	fmt.Printf("GC Runs: %d\n", mem.NumGC)
}

func printHardwareInfo(runner commandRunner) {
	fmt.Println("\n=== Hardware Info (WMIC) ===")
	fmt.Printf("CPU: %s\n", wmicInfo(runner, "cpu", "get", "name"))
	fmt.Printf("RAM: %s\n", formatBytes(wmicInfo(runner, "computersystem", "get", "TotalPhysicalMemory")))
	fmt.Printf("GPU: %s\n", wmicInfo(runner, "path", "win32_VideoController", "get", "name"))
	fmt.Printf("Disk: %s\n", wmicInfo(runner, "logicaldisk", "get", "caption,size,freespace"))
}

func main() {
	printRuntimeInfo()
	printHardwareInfo(defaultCommandRunner)
}
