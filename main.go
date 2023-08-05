package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//Initiate flags for program, so that input can be selected at runtime
	fReinput := flag.String("r", "./regex.txt", "Specify file regular expression is read from")
	fTxtinput := flag.String("i", "./input.txt", "Specify input file here\n If file is specified, file is parsed\n If directory is specified, directory is parsed recursively")
	fWildcard := flag.String("w", "*", "Limit what kind of files recursive parsing should go through")
	fPrintpath := flag.Bool("p", false, "Print paths that are read from input flag")
	fAppendpath := flag.Bool("a", false, "Print file path as seperate column in csv")
	fSeperatorchar := flag.String("s", ",", "Specify which character to use as seperator between matches")
	flag.Parse()

	if *fPrintpath {
		printpaths(recursivepathsearch(*fTxtinput, *fWildcard))
	} else {

		//Read regular expression into memory, and store SubexpNames into keys slice
		reinput, err := os.ReadFile(*fReinput)
		if err != nil {
			fmt.Printf("Regular expression can't be a directory\n")
			log.Fatal(err)
		}
		keys := printReSubexpNames(string(reinput), *fAppendpath, *fSeperatorchar)

		for _, v := range recursivepathsearch(*fTxtinput, *fWildcard) {
			if !isDir(v) {
				txtinput, err := os.ReadFile(v)
				if err != nil {
					log.Fatal(err)
				}
				totres := runRegexAllStringSubmatch((string(txtinput)), string(reinput), keys)
				printReSubexContents(totres, keys, v, *fSeperatorchar, *fAppendpath)
			}
		}
	}
}

func runRegexAllStringSubmatch(txtinput string, reinput string, keys []string) (totres []map[string]string) {
	re := regexp.MustCompile(reinput)
	for _, match := range re.FindAllStringSubmatch(txtinput, -1) {
		matchres := map[string]string{}
		totres = append(totres, matchres)
		for i, k := range keys {
			if k != "" && i != 0 {
				matchres[k] = match[i]
			}
		}
	}
	return
}

// print SubexNames, typically at top of file to get collums for csv
func printReSubexpNames(reinput string, fPrintpath bool, seperator string) (keys []string) {
	re := regexp.MustCompile(string(reinput))
	names := re.SubexpNames()
	for i, k := range names {
		if k != "" && i != 0 {
			fmt.Printf("%s%s", k, seperator)
		} else if i != 0 {
			fmt.Printf("%v%s", i, seperator)
		}
		if k != "" {
			keys = append(keys, k)
		} else {
			keys = append(keys, strconv.Itoa(i))
		}
	}
	if fPrintpath {
		fmt.Printf("filepath")
	}
	fmt.Printf("\n")
	return
}

// print contents of Subexpressions, typically done after printing SubexNames
func printReSubexContents(totres []map[string]string, keys []string, path string, seperator string, fPrintpath bool) {
	for _, m := range totres {
		for _, v := range keys {
			if v != "0" {
				fmt.Printf("%s%s", m[v], seperator)
			}
		}
		if fPrintpath {
			fmt.Printf("%s", path)
		}
		fmt.Printf("\n")
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
func recursivepathsearch(ipath string, matchpattern string) (paths []string) {
	filepath.WalkDir(ipath, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %v: %v\n",
				s,
				err)
			return err
		}
		if match(matchpattern, s) {
			paths = append(paths, s)
		}
		return nil
	})
	return paths
}

// wildcardToRegexp converts a wildcard pattern to a regular expression pattern.
// TODO: call with flag
func wildcardToRegexp(pattern string) string {
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
	result, _ := regexp.MatchString(wildcardToRegexp(pattern), value)
	return result
}

// check if path is dir or file
func isDir(path string) (isDir bool) {
	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	isDir = fileinfo.IsDir()
	return
}
