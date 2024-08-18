package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println()
	fmt.Println("runtime.MemStats")
	fmt.Printf("Alloc: %s\n", toMB(mem.Alloc))
	fmt.Printf("TotalAlloc: %s\n", toMB(mem.TotalAlloc))
	fmt.Printf("Sys: %s\n", toMB(mem.Sys))
	fmt.Printf("NumGC: %d\n", mem.NumGC)
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("HeapAlloc: %s\n", toMB(mem.HeapAlloc))
	fmt.Printf("HeapSys: %s\n", toMB(mem.HeapSys))
	fmt.Printf("HeapIdle: %s\n", toMB(mem.HeapIdle))
	fmt.Printf("HeapInuse: %s\n", toMB(mem.HeapInuse))
	fmt.Printf("HeapReleased: %s\n", toMB(mem.HeapReleased))
	fmt.Printf("HeapObjects: %d\n", mem.HeapObjects)
	fmt.Printf("StackInuse: %s\n", toMB(mem.StackInuse))
	fmt.Printf("StackSys: %s\n", toMB(mem.StackSys))
	fmt.Printf("MSpanInuse: %s\n", toMB(mem.MSpanInuse))
	fmt.Printf("MSpanSys: %s\n", toMB(mem.MSpanSys))
	fmt.Printf("MCacheInuse: %s\n", toMB(mem.MCacheInuse))
	fmt.Printf("MCacheSys: %s\n", toMB(mem.MCacheSys))
	fmt.Printf("BuckHashSys: %s\n", toMB(mem.BuckHashSys))
	fmt.Printf("GCSys: %s\n", toMB(mem.GCSys))
	fmt.Printf("OtherSys: %s\n", toMB(mem.OtherSys))
	fmt.Printf("NextGC: %d\n", mem.NextGC)
	fmt.Printf("LastGC: %d\n", mem.LastGC)
	fmt.Printf("PauseTotalNs: %d\n", mem.PauseTotalNs)
	fmt.Printf("PauseNs: %d\n", mem.PauseNs[(mem.NumGC+255)%256])
	fmt.Printf("PauseEnd: %d\n", mem.PauseEnd[(mem.NumGC+255)%256])
	fmt.Printf("NumForcedGC: %d\n", mem.NumForcedGC)
	fmt.Printf("GCCPUFraction: %.2f\n", mem.GCCPUFraction)
	fmt.Printf("EnableGC: %v\n", mem.EnableGC)
	fmt.Printf("DebugGC: %v\n", mem.DebugGC)
	fmt.Printf("BySize: %v\n", mem.BySize)
	fmt.Printf("Frees: %v\n", mem.Frees)
}

type Size uint64

const (
	Byte Size = 1 << (10 * iota)
	KB
	MB
)

func megabytes(b uint64) float64 {
	return float64(b) / float64(MB)
}

func toMB(b uint64) string {
	return fmt.Sprintf("%.2f MB", megabytes(b))
}
