### How To
Install go toolchain

```
git clone https://github.com/gronka/butter
cd butter
go install butter
```

Then run from anywhere with `butter {imagename}`

### Commands
* e -> Up a directory (to the very last child of the previous sibling dir)
* d -> Down a directory (to the first child of this directory or the next silbing dir)
* f -> Next image (in this dir)
* s -> Previous image (in this dir)


### TODO
* disable image wrapping
* goroutines for thumbnail generation by the crawler
* config file
* configurable input binding
* goroutines to preload images
* remember image position in each directory

### BUGS
* image scaling bugs on window resize
* decrementFolderPath cannot escape current parentPath
* keystrokes don't feel buttery
