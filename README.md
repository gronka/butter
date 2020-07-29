# butter
butter is a lightweight, home-row-based image gallery written in golang with 
opengl bindings and threading to create a buttery smooth experience.

### How To
Install go toolchain, then install go-gl. It has to compile, so takes a couple minutes:

```
go get -u github.com/go-gl/glfw/v3.3/glfw
go get -u github.com/go-gl/gl/v2.1/gl
```

Install butter:

```
git clone https://github.com/gronka/butter
cd butter
go install butter
```

Then run from anywhere with `butter {imagename}`

### Commands
Commands revolve around esdf (and later ijkl)

* f -> Next image (in this dir)
* s -> Previous image (in this dir)
* e -> Up a directory (to the very last child of the previous sibling dir)
* d -> Down a directory (to the first child of this directory or the next silbing dir)
