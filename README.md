# xkcd Tool
##### Exercize 4.12 - The Go Programming Language

An example tool for fetching xkcd API endpoint and saving the content into different files.
These files can be used as an offline index for searching a keyword inside all the transcripts.

**Features**
* Fetch API
* Save output to file
* Read local JSON files
* Show output as HTML template

**Usage:**
* First of all you need to save all the episodes in local:
```bash
  go run getAll.go
```
Note: consider launching the command inside a directory called *offlineIndex*, so that you don't need to change the code.

* Now you can run the following command to search for the keyword on your local DB (inside *offlineIndex*):
```bash
  go run xkcd.go <keyword> > out.html
```
Open *out.html* with your favourite browser. 