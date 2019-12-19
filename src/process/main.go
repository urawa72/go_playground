package main

import (
  "fmt"
  "github.com/shirou/gopsutil/process"
)

func main() {
  pid := 1
  ps, err := process.NewProcess(int32(pid))
  if err != nil {
    panic(err)
  }

  cpu, err := ps.CPUPercent()
  if err != nil {
    panic(err)
  }

  pm, err := ps.MemoryInfo()
  if err != nil {
    panic(err)
  }

  // fmt.Printf("Memory Use(RSS): %d [kB]:", pm.RSS/1024)
  fmt.Printf("CPU Use: %.2f [%%], Memory Use(RSS): %d [kB]:", cpu, pm.RSS/1024)
}
