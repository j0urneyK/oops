# Logrus formatter for Oops

```go
import oopslogrus "github.com/j0urneyK/oops/loggers/logrus"

func init() {
    logrus.SetFormatter(
        oopslogrus.NewOopsFormatter(
	    &logrus.JSONFormatter{},
        )
    )
}

func main() {
    err := oops.
        With("driver", "postgresql").
        With("query", query).
        With("query.duration", queryDuration).
        Errorf("could not fetch user")

    if err != nil {
    	logrus.WithError(err).Error(err)
    }
}
```
