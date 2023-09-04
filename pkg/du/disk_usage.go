package du

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

type diskUsage struct{}

type stdRow struct {
	Size string
	Path string
}

func NewStdRow(size, path string) StdRow {
	return &stdRow{
		Size: size,
		Path: path,
	}
}

func (r *stdRow) GetSize() string {
	return r.Size
}

func (r *stdRow) GetPath() string {
	return r.Path
}

func NewDiskUsage() DiskUsage {
	return &diskUsage{}
}

// GetExecutablePath returns the path of the given command
func (du *diskUsage) GetExecutablePath(command string) (string, error) {
	return exec.LookPath(command)
}

// GetDiskUsage returns the disk usage of the given path
func (du *diskUsage) GetDiskUsage(path string, args string) ([]StdRow, error) {
	dhPath, er := du.GetExecutablePath("du")
	if er != nil {
		return nil, ErrNotFoundDiskUsageExecutable
	}

	cmd := exec.Command(dhPath, args, path)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	pipeOutput, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	byteResult, err := io.ReadAll(pipeOutput)
	if err != nil {
		return nil, err
	}

	var rows []StdRow
	splitResult := strings.Split(string(byteResult), "\n")

	for _, line := range splitResult {
		splitRow := strings.Split(line, "\t")
		if len(splitRow) == 0 {
			continue
		} else {
			if splitRow[0] == "" {
				continue
			}
		}
		stdRow := NewStdRow(splitRow[0], splitRow[1])
		rows = append(rows, stdRow)
	}

	return rows, nil
}
