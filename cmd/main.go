package main

import (
	"fmt"

	"github.com/osvaldoabel/db-disk-usage/pkg/du"
)

type Row struct {
	Size string
	Path string
}

var (
	Rows []string
)

func main() {
	du := du.NewDiskUsage()
	mysqlPath := "/opt/homebrew/var/mysql/"
	args := "-hxd 1"

	result, er := du.GetDiskUsage(mysqlPath, args)
	if er != nil {
		fmt.Println("Error 1")
	}
	for _, r := range result {
		Rows = append(Rows, r.GetSize()+"\t"+r.GetPath())
	}
	fmt.Println(Rows)
}
