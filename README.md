This is a test middleware for Ginonic powered by [Coraza Web Application Firewall](https://github.com/jptosso/coraza-waf). You may check the WAF documentation at [coraza.io](https://coraza.io/)

Looking for contributors and testers.

## How to use

```go
import(
    //...
    coraza"github.com/jptosso/coraza-waf"
    "github.com/jptosso/coraza-waf/seclang"
    corazagin"github.com/jptosso/coraza-gin"
)
func main() {
    // Creates a router without any middleware by default
    r := gin.New()
    waf := coraza.NewWaf()
    parser := seclang.NewParser(waf)
    //parser.FromString(`#... some rules`)
    r.Use(corazagin.Coraza(waf))

    // Per route middleware, you can add as many as you desire.
    r.GET("/mypath", MyFunction(), Endpoint)

    // Listen and serve on 0.0.0.0:8080
    r.Run(":8080")
}
```