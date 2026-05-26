package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorAlert(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	escapedMsg, _ := json.Marshal(message)
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<script>
		alert(%s);
		window.history.back();
	</script>
</head>
<body></body>
</html>`, string(escapedMsg))
}
