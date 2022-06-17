package coraza

import (
	"io"

	"github.com/corazawaf/coraza/v2"
	"github.com/corazawaf/coraza/v2/types"
	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	tx *coraza.Transaction

	headersProcessed bool
	size             int
}

func (w responseWriter) Write(b []byte) (n int, err error) {
	if it := w.processResponseHeaders(); it != nil {
		// transaction was interrupted :(
		return
	}
	w.WriteHeaderNow()
	n, err = w.tx.ResponseBodyBuffer.Write(b)
	w.size += n
	return
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	if it := w.processResponseHeaders(); it != nil {
		// transaction was interrupted :(
		return
	}
	w.WriteHeaderNow()
	n, err = io.WriteString(w.tx.ResponseBodyBuffer, s)
	w.size += n
	return
}

func (w *responseWriter) processResponseHeaders() *types.Interruption {
	if w.headersProcessed || w.tx.Interruption != nil {
		return w.tx.Interruption
	}
	w.headersProcessed = true
	for k, vv := range w.ResponseWriter.Header() {
		for _, v := range vv {
			w.tx.AddResponseHeader(k, v)
		}
	}
	return w.tx.ProcessResponseHeaders(w.ResponseWriter.Status(), "http/1.1")
}

func (w *responseWriter) Status() int {
	if w.tx.Interruption != nil {
		return w.tx.Interruption.Status
	}
	return w.ResponseWriter.Status()
}

func (w *responseWriter) Size() int {
	return w.size
}
