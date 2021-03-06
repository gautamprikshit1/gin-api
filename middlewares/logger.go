package middlewares

import (
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("%s - [%s] %s %v %s %d %v %s\n",
        param.ClientIP,
        param.TimeStamp.Format(time.RFC822),
        param.Method,
        param.MethodColor(),
        param.Path,
        param.StatusCode,
        param.StatusCodeColor(),
        param.Latency,
    )
})
}
