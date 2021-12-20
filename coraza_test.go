package coraza

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jptosso/coraza-waf/v2"
)

func TestMiddleware1(t *testing.T) {
	waf := coraza.NewWaf()
	router := setupTestRouter(waf)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("failed to set status")
	}
	if w.Body.String() != "pong" {
		t.Error("failed to set body, got: " + w.Body.String())
	}
}

func setupTestRouter(waf *coraza.Waf) *gin.Engine {
	r := gin.Default()
	r.Use(Coraza(waf))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
