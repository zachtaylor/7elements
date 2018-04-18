SE.widget.control('se-app-login', function() {
	var me = this;

	me.update = function() {
		if (window.location.search == "?account") {
			$("#form-group-username input").addClass("form-error");
		}
		if (window.location.search == "?password") {
			$("#form-group-password input").addClass("form-error");
		}
	};
});
