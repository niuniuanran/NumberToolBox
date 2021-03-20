# Number Toolbox
The first command line tool I build with Go.

## Add it to your command line
```
$ git clone https://github.com/niuniuanran/NumberToolBox.git
$ cd NumberToolBox
$ go build numbertool.go
$ cp numbertool /usr/local/bin
```

## Example use case
### Factors
``` 
$ numbertool -tool factors -number 3948102
```
Output:
``` 
The factors of 3948102 are: [2 3 3 3 3 24371]
```

### Prime
```
$ numbertool -tool prime -number 3413453
```
Output:
```
3413453 is not a prime number
```

### Prime showing factors
``` 
$ numbertool -tool prime -number 3413453 -v
```
Output:
``` 
3413453 is not a prime number
The factors of 3413453 are: [23 148411]
```




