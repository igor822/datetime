DateTime
========

This is a simple library to format date and time strings.  
Here is some example

```go
import(
    "github.com/igor822/datetime"
    "fmt"
)

dt := datetime.NewDateTime()
output := dt.Format("Y-m-d", "d/m/Y", "2018-01-30 15:30:14")
fmt.Println(output)
```

So the `Format()` receives 3 parameters `fromFormat`, `toFormat` and `datetime` string.
