package procmeminfo

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MemInfo is a map[string]uint64 with all the values found in /proc/meminfo
type MemInfo map[string]uint64

// Update s with current values, usign the pid stored in the Stat
func (m *MemInfo) Update() error {
	var err error

	path := filepath.Join("/proc/meminfo")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		n := strings.Index(text, ":")
		if n == -1 {
			continue
		}

		key := text[:n] // metric
		data := strings.Split(strings.Trim(text[(n+1):], " "), " ")
		if len(data) == 1 {
			value, err := strconv.ParseUint(data[0], 10, 64)
			if err != nil {
				continue
			}
			(*m)[key] = value
		} else if len(data) == 2 {
			if data[1] == "kB" {
				value, err := strconv.ParseUint(data[0], 10, 64)
				if err != nil {
					continue
				}
				(*m)[key] = value * 1024
			}
		}

	}
	return nil

}

// Total return the size of the memory in bytes.
// It is an alias of (*m)["MemInfo"]
func (m *MemInfo) Total() uint64 {
	return (*m)["MemTotal"]
}

// Available return the available memory following this formula:
//	Available = MemFree + Buffers + Cached
func (m *MemInfo) TotalAvailable() uint64 {
	return m.Available() + m.Buffers() + m.Cached()
}

// Available return MemFree from /proc/meminfo:
//
//	Available = MemFree
func (m *MemInfo) Available() uint64 {
	return (*m)["MemFree"]
}

// Used is a generic way of reporting used memory. It follows the next formula:
//
//	TotalUsed = MemTotal - TotalAvailable()
func (m *MemInfo) TotalUsed() uint64 {
	return m.Total() - m.TotalAvailable()
}

// Used is a generic way of reporting used memory. It follows the next formula:
//
//	Used = MemTotal - Available()
func (m *MemInfo) Used() uint64 {
	return m.Total() - m.Available()
}

// Cached /proc/meminfo Cached:
func (m *MemInfo) Cached() uint64 {
	return (*m)["Cached"]
}

// Buffers /proc/meminfo Buffers
func (m *MemInfo) Buffers() uint64 {
	return (*m)["Buffers"]
}

// Swap returns the % of swap used
func (m *MemInfo) Swap() int {
	total := (*m)["SwapTotal"]
	free := (*m)["SwapFree"]
	if total == 0 {
		return 0
	}
	return int((100 * (total - free)) / total)
}
