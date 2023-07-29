## TODO:
  
  Test code to "completetion"
  
  Flags
  
  Write proper readme

Check TODO: comments in .go files

This is supposed to be a CLI program that can be used to run a regular expression on a single file, or recursively on several files. Wildcards can be used to limit which files are read.
Regular expression is read from file. 
Output is then printed to terminal window. This output can of course be piped further to the likes of grep or stdout. 
Output is printed as csv, with regex group names printed at first line for grouping into collums. 

There are planned flags for:
*Setting if the program should be recursive or single file
*Printing path for file the printed line is read from
*More to come