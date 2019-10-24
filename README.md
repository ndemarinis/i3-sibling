# i3-sibling

`i3-sibling` is a tiny script based on `i3ipc-go` to switch between sibling tabs in i3.  

I use it to cycle between tabs of the same container.  Depending on your cofiguration, your mileage may vary.  

## Installation

To build
```
go build
go install # Should install to $GOPATH/bin
```

## Usage
 
To switch to the next sibling:
```
$ i3-sibling next
```

To switch to the previous sibling:
```
$ i3-sibling prev
```

When reaching the end of the sibling list, the switcher will wrap around to the first sibling.  


This is most useful when bound to keys (eg. Super-Tab)
```
bindsym $mod+Tab exec "~/bin/i3-sibling next"
bindsym $mod+Shift+Tab exec "~/bin/i3-sibling prev"
```

## Known bugs

None yet.  If you have any issues, please feel free to submit a pull request!
