What I've aimed to create here can be more or less duplicated with ripgrep as such:


rg --with-filename --no-heading --multiline --only-matching --replace '$1, $2, $3' -f pattern *.ps


If i'm to update this, there will most likely be a full rewrite at some point. 


## TODO:
  
 * Write tests
  
 * Write proper readme

 * Test and write proper errors for non functioning regex/input

 * Concurrently search through files

 * Include most of grep's functionality

Check TODO: comments in .go files
## Readme:

  This is a CLI program that can be used to run a regular expression on a single file, or recursively on several files. 
      
  Wildcards can be used to limit which files are read.  
  
  Regular expression is read from file or directory. If read from directory, it will parse recursively.  
  
  Output is then printed to terminal window. This output can be piped further to grep or stdout.  
  
  Output with capture groups is printed with a seperator character that can be specified. Uniqe matches are printed on new line. If printing paths with -p all paths will be seperated with new line.  

  Named capture groups are supported. If they are not used they will be numbered from 0 to n, where 0 is the complete match of regular expression.   
  
  Regex group names and file names of parsed file can be printed out at first line as column headers.

## Currently implemented flags

| Description                                       | flag | Defaults     |
|---------------------------------------------------|------|--------------|
| Parse single file or directory                    | i    | ./input.txt  |
| Choose which file to read regular expression from | r    | ./regex.txt  |
| Limit recursive parsing with wildcard             | w    | *            |
| Print paths that are read from input flag         | p    | false        |
| Appends filepath as last column in csv printing   | a    | false        |
| Omit full match from printed results              | o    | false        |
| Select which character to use as seperator        | s    | ,            |
| Charmap to decode from                            | c    | Windows 1252 |
| Lists all available charmaps                      | lc   | false        |
| Decode string in format chosen by argument c      | d    | false        |

## There are planned flags for: 
| Description                        | flag | Defaults |
|------------------------------------|------|----------|
|Printing line numbers next to match |      |          |
|Printing full line containing match |      |          |




