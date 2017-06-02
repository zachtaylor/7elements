package serverutil

import (
	"fmt"
	"net/http"
)

const redirectHomeTpl = `<html>
	<head><title>redirect...</title></head>
	<body>
		<h3>%s</h3>
		<span>redirecting to homepage in 5 seconds... <a href="/">Click here to go now</a></span>
		<script type="text/javascript">
			window.setTimeout(function() {
				window.location.pathname="";
			}, 5000);
		</script>
	</body>
</html>`

func WriteRedirectHome(w http.ResponseWriter, reason string) {
	var content = fmt.Sprintf(redirectHomeTpl, reason)

	w.Write([]byte(content))
}
