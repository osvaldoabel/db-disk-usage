package du

// DiskUsage is the interface that wraps the basic methods to get disk usage
// information from a given path
type DiskUsage interface {
	// GetDiskUsage returns the disk usage of the given path
	GetDiskUsage(path string, args string) ([]StdRow, error)
	// GetExecutablePath returns the path of the given command
	GetExecutablePath(command string) (string, error)
}

type StdRow interface {
	// GetSize returns the size of the file or directory
	GetSize() string
	// GetPath returns the path of the file or directory
	GetPath() string
}
