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
  
  Output is then printed to terminal window. This output can be piped further to the likes of grep or stdout.  
  
  Output is printed as csv expect for when just printing filepaths.  
  Named capture groups are supported, and they are not used they will be numbered from 1 to n.  
  Regex group names and file names of parsed file can be printed out at first line as column headers.

## Currently implemented flags

|   | Description                                           | flag   | Defaults      |
|---|-------------------------------------------------------|--------|---------------|
|   | Parse single file or directory                        | i      | ./input.txt   |
|   | Choose which file to read regular expression from     | r      | ./regex.txt   |
|   | Limit recursive parsing with wildcard                 | w      | *             |
|   | Print paths that are read from input flag             | p      | false         |
|   | Appends filepath as last column in csv printing       | a      | false         |
|   | Select which character to use as seperator            | s      | ,             |

## There are planned flags for: 
|   | Description                                                                            | flag | Defaults |
|---|----------------------------------------------------------------------------------------|------|----------|
|   | Allowing printing of whole regular expressions, not just groups                        |      |          |