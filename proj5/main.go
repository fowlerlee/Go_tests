package proj5

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var regexes = []*regexp.Regexp{
	regexp.MustCompile(`(?i)user`),
	regexp.MustCompile(`(?)password`),
	regexp.MustCompile(`kdb`),
	regexp.MustCompile(`login`),
}

func walkFunc(path string, f os.FileInfo, err error) error  {
	for _, r := range regexes {
		if r.MatchString(path) {
			fmt.Printf("[+] HIT: %s\n", path)
		}
	}
	return nil
}

func main() {
	root := os.Args[1]
	if err := filepath.Walk(root, walkFunc); err != nil {
		log.Panicln(err)
	}
}