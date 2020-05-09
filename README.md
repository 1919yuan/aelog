## AppEngine Logger Wrapper

Packge `github.com/1919yuan/aelog` provides a wrapping utility for logging to
stdout when testing locally and logging to StackDriver when running on
AppEngine.

### Example Usage

```go
import (
  "github.com/1919yuan/aelog"
)

...

aelog.Debug("Testing: Debug")
aelog.Info("Testing: Info")
aelog.Warning("Testing: Warning")
aelog.Error("Testing: Error")
aelog.Fatal("Testing: Fatal")
```

The logging will be done using the `log` package when the program is running
locally and will be done using `cloud.google.com/go/logging.Logger` when the
program is running in AppEngine environment.
