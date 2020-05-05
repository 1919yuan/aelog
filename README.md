## AppEngine Logger Wrapper

Packge `github.com/1919yuan/aelog` provides a wrapping utility for logging to
stdout when testing locally and logging to StackDriver when running on
AppEngine.

### Example Usage

```go
import (
  "context"
  "github.com/1919yuan/aelog"
)

...

Log = aelog.CreateStackDriverLogger(context.Background(), nil)

...

Log.Debug("Testing StackDriverLogger: Debug")
Log.Info("Testing StackDriverLogger: Info")
Log.Warning("Testing StackDriverLogger: Warning")
Log.Error("Testing StackDriverLogger: Error")
Log.Fatal("Testing StackDriverLogger: Fatal")
```

The logging will be done using the `log` package when the program is running
locally and will be done using `cloud.google.com/go/logging.Logger` when the
program is running in AppEngine environment.
