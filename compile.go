package oh_sass

import (
	"github.com/yosssi/gcss"
	"os"
)

func SassCompile(source string) (string, error) {
	path,err := gcss.CompileFile(source)
	if err != nil {
		return path,err
	}
	return path,err
}

func MoveTo(source, output string) error {
	err := os.Rename(source,output)
	if err != nil {
		return err
	}
	return err
}

func SassComplieTo(source, output string) error {
	path,err := gcss.CompileFile(source)
	if err != nil {
		return err
	}

	err = MoveTo(path, output)
	if err != nil {
		return err
	}
	return err
}
