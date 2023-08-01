## TODO:
  
 * Write tests
  
 * Flags as described at bottom of readme

 * Write proper readme

 * Test and write proper errors for non functioning regex/input

 * Concurrently search through files

Check TODO: comments in .go files
## Readme:

  This is supposed to be a CLI program that can be used to run a regular expression on a single file, or recursively on several files. 
      
  Wildcards can be used to limit which files are read.  
  
  Regular expression is read from file or directory. If read from directory, it will parse recursively.  
  
  Output is then printed to terminal window. This output can of course be piped further to the likes of grep or stdout.  
  
  Output is printed as csv expect for when just printing filepaths. 
  Regex group names and file names of parsed file can be printed out at first line as column headers.

## Currently implemented flags

|   | Description                                           | flag   | Defaults      |
|---|-------------------------------------------------------|--------|---------------|
|   | Parsing single file or directory                      | i      | ./input.txt   |
|   | Limiting recursive parsing with wildcard              | w      | *             |
|   | Choosing which file to read regular expression from   | r      | ./regex.txt   |
|   | Appends filepath as last column in csv printing       | p      | false         |


## There are planned flags for: 
| Description                                                                                        | flag | Defaults |
|----------------------------------------------------------------------------------------------------|------|----------|
| Allowing printing of whole regular expressions, not just groups                                    |      |          |
| Allow for unnamed capture groups. Will then print index number for capture groups at top of file   |      |          |
| Function to print all paths to run regular expression on. For sanity checking purposes             |      |          |
| Select which character to use as seperator                                                         |      |      ,   |
