## TODO:
  
  Write tests
  
  Complete flags
  
  Write proper readme

  Test and write proper errors for non functioning input

  Write flag for outputting all files that are read recursively. Can then be used as a sanity check before running program

Check TODO: comments in .go files

This is supposed to be a CLI program that can be used to run a regular expression on a single file, or recursively on several files. Wildcards can be used to limit which files are read.  
Regular expression is read from file or directory. If read from directory, it will parse recursively.  
Output is then printed to terminal window. This output can of course be piped further to the likes of grep or stdout.  
Output is printed as csv, with regex group names, and file name of parsed file printed at first line for grouping into collums.  

## Currently implemented flags

| Description                                         | flag | Defaults    |
|-----------------------------------------------------|------|-------------|
| Parsing single file or directory                    | i    | ./input.txt |
| Limiting recursive parsing with wildcard            | w    | *           |
| Choosing which file to read regular expression from | r    | ./regex.txt |


## There are planned flags for: 
| Description                                                                                        | flag | Defaults |
|----------------------------------------------------------------------------------------------------|------|----------|
| Allowing printing of whole regular expressions, not just groups                                    |      |          |
| Allow for unnamed capture groups. Will then print index number for capture groups at top of file   |      |          |
| Printing path for file the printed line is read from                                               |      |          |
| More to come?                                                                                      |      |          |