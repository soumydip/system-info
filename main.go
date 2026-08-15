package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func GetRuntimeData() {
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Compiler:", runtime.Compiler)
}

func GetGoMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Allocated Memory (Go program):", m.Alloc/1024, "KB")
	fmt.Println("Total Allocated (cumulative):", m.TotalAlloc/1024, "KB")
	fmt.Println("System Memory obtained from OS:", m.Sys/1024, "KB")
	fmt.Println("Number of GC runs:", m.NumGC)
}

func GetOsData() {
	hostName, _ := os.Hostname()
	fmt.Println("Host Name:", hostName)
	fmt.Println("Operating System:", os.Getenv("OS"))
	fmt.Println("Architecture:", os.Getenv("PROCESSOR_ARCHITECTURE"))
	fmt.Println("Number of CPUs:", os.Getenv("NUMBER_OF_PROCESSORS"))
	fmt.Println("User Name:", os.Getenv("USERNAME"))
	fmt.Println("User Domain:", os.Getenv("USERDOMAIN"))
	fmt.Println("Processor Identifier:", os.Getenv("PROCESSOR_IDENTIFIER"))
	fmt.Println("Processor Level:", os.Getenv("PROCESSOR_LEVEL"))
	fmt.Println("Processor Revision:", os.Getenv("PROCESSOR_REVISION"))
	fmt.Println("System Drive:", os.Getenv("SystemDrive"))
	fmt.Println("System Root:", os.Getenv("SystemRoot"))
	fmt.Println("Temp Directory:", os.TempDir())
	fmt.Println("Process ID:", os.Getpid())
	fmt.Println("Parent Process ID:", os.Getppid())

	dir, _ := os.Getwd()
	fmt.Println("Current Working Directory:", dir)
}

func GetTimeData() {
	now := time.Now()
	fmt.Println("Current Time:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Timezone:", now.Location())
}

func ExcuteScript(script string) {
	output, err := exec.Command("cmd", "/C", script).Output()
	if err != nil {
		fmt.Println("Error executing script:", err)
	} else {
		fmt.Println("Script output:", string(output))
	}
}

func GetCPUInfo() {
	ExcuteScript("wmic cpu get Name,NumberOfCores,NumberOfLogicalProcessors,MaxClockSpeed")
}

func GetRAMInfo() {
	ExcuteScript("wmic memorychip get Capacity,Speed,Manufacturer")
}

func GetGPUInfo() {
	ExcuteScript("wmic path win32_VideoController get Name,AdapterRAM,DriverVersion")
}

func GetDiskInfo() {
	ExcuteScript("wmic logicaldisk get DeviceID,Size,FreeSpace,VolumeName")
}

func main() {
	fmt.Println("=== Runtime Data ===")
	GetRuntimeData()

	fmt.Println("\n=== Go Memory Stats ===")
	GetGoMemoryStats()

	fmt.Println("\n=== OS Data ===")
	GetOsData()

	fmt.Println("\n=== Time Data ===")
	GetTimeData()

	fmt.Println("\n=== CPU Info (via wmic) ===")
	GetCPUInfo()

	fmt.Println("\n=== RAM Info (via wmic) ===")
	GetRAMInfo()

	fmt.Println("\n=== GPU Info (via wmic) ===")
	GetGPUInfo()

	fmt.Println("\n=== Disk Info (via wmic) ===")
	GetDiskInfo()
}
