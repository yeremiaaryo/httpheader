package httpheader

import "net/http"

func AddHeader(req *http.Request, key string, value string) {
	req.Header[key] = []string{value}
	return
}
