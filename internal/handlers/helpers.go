package handlers

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
func SuccessAlert(w http.ResponseWriter, message string, redirectURL string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)

	escapedMsg, _ := json.Marshal(message)
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<script>
		alert(%s);
		window.location.href = %s;
	</script>
</head>
<body></body>
</html>`, string(escapedMsg), fmt.Sprintf(`"%s"`, redirectURL))
}
