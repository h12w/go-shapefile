go-shapefile
============

Go binding to the shapefile library.

Dependencies
------------

Requires the https://github.com/mousebird/shapelib C library.

For OS X, install with:

```
brew install shapelib
```

Installation
------------

```
go get github.com/hailiang/go-shapefile
```

Example usage
-------------

See `examples/example.go` for basic example usage to load a `.shp` file (and its `.dbf` companian file).

```
go run examples/example.go map/bou2_4p.shp
```
