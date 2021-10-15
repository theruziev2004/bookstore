package pkg

import (
	"fmt"
	"net/http"
)

func SendString(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	_, _ = fmt.Fprint(writer, message)
}
