package test

import (
	"fmt"
	"testing"
)

// TestHashIp 测试落在各个IP上的概率
func TestHashIp(t *testing.T) {
	// 负载
	load := 10
	index := make([]int, load)
	count := 0
	for i := 1; i < 255; i++ {
		for j := 1; j < 255; j++ {
			for k := 1; k < 255; k++ {
				local := (i + j + k) % load
				index[local]++
				count++
			}
		}
	}
	for i, v := range index {
		fmt.Print(i, v)
		fmt.Printf(" %.4f\n", float64(v)/float64(count))
	}
}
