// v0.2.0b
// Author: DIEHL E.
// (©) Sony Pictures Entertainment, Feb 2021

package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/wunderbarb/mypkg/toolbox"
)

func CreateHeader(name string) error {

	name = toolbox.SetExtension(name, "go")
	if toolbox.IsExist(name) {
		return fmt.Errorf("%s already exists", name)
	}

	packageName := findPackageName()

	f, err := os.Create(name)
	if err != nil {
		return errors.Wrap(err, "cannot create file")
	}
	defer f.Close()

	f.WriteString("// v0.1.0\n// Author: DIEHL E.\n")
	f.WriteString(fmt.Sprintf("// © Sony Pictures Entertainment, %s\n",
		time.Now().Format("Jan 2006")))

	f.WriteString(fmt.Sprintf("\npackage %s\n\n", packageName))
	return nil
}

func findPackageName() string {
	ls, n := toolbox.ListOfFilesWithExt(".", ".go")
	if n == 0 {
		return "main"
	}
	fh, err := os.Open(ls[0])
	if err != nil {
		return "main"
	}

	defer fh.Close()
	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {

		s := strings.Trim(fileScanner.Text(), "\r\n") // It is mandatory to strip also \r added by Windows
		sParsed := strings.Fields(s)
		if len(sParsed) == 2 {
			if sParsed[0] == "package" {
				return sParsed[1]
			}
		}
	}
	return "main"
}

func FindFunctionDeclaration(fn string, fileName string) ([]string, int, error) {
	fh, err := os.Open(fileName)
	if err != nil {
		return nil, 0, errors.Wrap(err, "could not open file")
	}
	defer fh.Close()
	fileScanner := bufio.NewScanner(fh)
	for fileScanner.Scan() {
		s := strings.Trim(fileScanner.Text(), "\r\n") // It is mandatory to strip also \r added by Windows
		sParsed := strings.Fields(s)
		if len(sParsed) >= 3 {
			detect := fn + "("
			if sParsed[0] == "func" {
				for i, s := range sParsed[1:] {
					if strings.Contains(s, detect) {
						return sParsed[1:], i, nil
					}
				}

			}
		}
	}
	return nil, 0, errors.New("not here")

}
