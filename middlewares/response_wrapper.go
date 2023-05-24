package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

func ResponseWrapperMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bw := &responseBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: ctx.Writer,
		}

		ctx.Writer = bw

		ctx.Next()

		for _, err := range ctx.Errors {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"data":   err.Error(),
			})
			return
		}

		status := bw.Status()
		data := bw.body.String()

		var response interface{}

		err := json.Unmarshal(bw.body.Bytes(), &response)
		if err != nil {
			log.Error(err)
			return
		}

		wrapper := make(map[string]interface{})
		wrapper["status"] = status
		wrapper["data"] = response

		newBody, _ := json.Marshal(wrapper)

		bw.ResponseWriter.Write(newBody)

		log.Info(status, data)
	}
}
