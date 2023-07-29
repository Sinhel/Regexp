package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	//Set default value for input files as input.txt, regex.txt in same the folder as the program. Files can also be specified as arguments in cli
	txtinputarg := "./input.txt"
	reinputarg := "./regex.txt"
	if len(os.Args) >= 2 {
		txtinputarg = os.Args[1]
		reinputarg = os.Args[2]
	}

	//Read files into memory
	reinput, err := os.ReadFile(reinputarg)
	if err != nil {
		log.Fatal(err)
	}
	txtinput, err := os.ReadFile(txtinputarg)
	if err != nil {
		log.Fatal(err)
	}
	runRegex(string(txtinput), string(reinput))
	recursivepathsearch(".", "*")
}

func runRegex(txtinput string, reinput string) (totres []map[string]string) {
	re := regexp.MustCompile(reinput)
	for _, match := range re.FindAllStringSubmatch(txtinput, -1) {
		res := map[string]string{}
		totres = append(totres, res)
		for i, subexpName := range re.SubexpNames() {
			if subexpName != "" {
				res[subexpName] = match[i]
			}
		}
	}
	keys := re.SubexpNames()
	for _, k := range keys {
		if k != "" {
			fmt.Printf("%s,", k)
		}
	}
	fmt.Printf("\n")
	for _, m := range totres {
		for _, k := range keys {
			if k != "" {
				fmt.Printf("%s,", m[k])
			}
		}
		fmt.Printf("\n")
	}
	return
}

func printSubexpNames(subexpNames []string) {
	for i, name := range subexpNames {
		if i != 0 && name != "" {
			fmt.Printf("%v,",
				name)
		}
	}
	fmt.Print("\n")
	for i := range subexpNames {
		for _, name := range subexpNames {
			if name != "" {
				fmt.Printf("%v,",
					subexpNames)
			}
		}
		fmt.Print("\n")
		i += i
	}
}

// test funtcion to print paths to console.
// TODO: call with flag
func printpaths(paths []string) {
	for i := range paths {
		fmt.Printf("%s",
			paths[i])
		fmt.Printf("\n")
	}
}

// TODO: call with flag
func recursivepathsearch(ipath string, matchpattern string) (opath []string) {
	filepath.WalkDir(ipath,
		func(s string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %v: %v\n",
					s,
					err)
				return err
			}
			if match(matchpattern, s) {
				opath = append(opath, s)
			}
			return nil
		})
	return opath
}

// wildCardToRegexp converts a wildcard pattern to a regular expression pattern.
// TODO: call with flag
func wildCardToRegexp(pattern string) string {
	components := strings.Split(pattern, "*")
	if len(components) == 1 {
		// if len is 1, there are no *'s, return exact match pattern
		return "^" + pattern + "$"
	}
	var result strings.Builder
	for i, literal := range components {

		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}

		// Quote any regular expression meta characters in the
		// literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return "^" + result.String() + "$"
}

func match(pattern string, value string) bool {
	result, _ := regexp.MatchString(wildCardToRegexp(pattern), value)
	return result
}
