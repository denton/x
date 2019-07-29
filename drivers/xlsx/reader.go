package xlsx

import (
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

// Driver represents structure needed
// for manipulationg data in xlsx files
type Driver struct {
	filename string
	file     *xlsx.File
}

// Open initialises driver from disc file
func (x *Driver) Open(filename string) error {

	var err error

	x.file, err = xlsx.OpenFile(filename)
	if err != nil {
		return errors.Wrap(err, "dsd")
	}
	return nil
}

//
func (x *Driver) HasTabs() bool {
	return true
}
