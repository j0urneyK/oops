# Logrus formatter for Oops

```go
import oopsrecoverygin "github.com/j0urneyk/oops/recovery/gin"

func main() {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(oopsrecoverygin.GinOopsRecovery())

    // ...
}
```
