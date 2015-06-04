# uwAPI.go
"OO" wrapper for the UWaterloo Open Data API, written in Golang

## Installation
As per usual:

```
go get github.com/SaintDako/uwAPI.go
```

## Dependencies
This wrapper depends on the glorious [gabs package](https://github.com/Jeffail/gabs), so go get gabs (pun intended) first.

## Usage
Pretty straightforward:

```golang
package main

import uwapi "github.com/SaintDako/uwAPI.go"

func main() {
	API_KEY := "YOUR_API_KEY_HERE"
	uw := uwapi.Create()

	jsonObj, _ := uw.FoodServices.Menu()
}
```

Then do whatever with the returned `gabs.Container` (check out their docs for getting, setting, etc...).

As seen in the example, the structs (e.g. `FoodServices`) are all named appropriately, i.e. the same as they are in the UW API's docs. However, the methods are not always of the same name due to conflicts or parameters.

Check out the [documentation](https://godoc.org/github.com/SaintDako/uwAPI.go) for more info.

## TODO

- add tests