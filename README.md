go-shapefile
============

Go binding to the shapefile library.

About shapefile
---------------
The [Esri shapefile](http://en.wikipedia.org/wiki/Shapefile), or simply a shapefile, is a popular geospatial vector data format for geographic information system software.

ESRI Shapefile Technical Description:
http://www.esri.com/library/whitepapers/pdfs/shapefile.pdf

Dependencies
------------

Shapefile C Library: http://shapelib.maptools.org

For Ubuntu Linux, install with:
```
apt-get install shapelib
```

For OS X, install with:

```
brew install shapelib
```

Installation
------------

```
go get -u github.com/hailiang/go-shapefile
```

Example usage
-------------

See `examples/example.go` for basic example usage to load a `.shp` file (and its `.dbf` companian file).

```
go run examples/example.go map/bou2_4p.shp
```
