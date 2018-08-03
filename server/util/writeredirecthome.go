package serverutil

import (
	"fmt"
	"net/http"
)

const redirectHomeTpl = `<html>
	<head>
		<title>Redirect</title>
	</head>
	<body>
		<img src="/img/banner-black.64px.png">
		<h3>%s</h3>
		<span style="font-size:21px">
			Redirecting to game in <b>30 s</b><br/>
			Click anywhere or press any key to go now
		</span>
		<script type="text/javascript">
			window.setTimeout(function() {
				window.location.pathname="/";
			}, 30000);
			document.addEventListener("click", function(e) {
				window.location.pathname="/";
			});
			document.addEventListener("keydown", function(e) {
				window.location.pathname="/";
			});
		</script>
	</body>
</html>
`

func WriteRedirectHome(w http.ResponseWriter, reason string) {
	w.Write([]byte(fmt.Sprintf(redirectHomeTpl, reason)))
}
