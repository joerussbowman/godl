--- THIS DOES NOT WORK YET, WORK IN PROGRESS ---

Writes content from stdin to FILENAME.

FILENAME can include a date template using the date
Mon Jan 2 15:04:05 -0700 MST 2006 as the layout surrounded
by %%. For example
    
    godl http_%%20060101%%.log

OPTIONS
    -z flag sets the file to be compressed as it's written.
