SE.widget.control('se-app-signup', function() {
	var me = this;

	me.update = function() {
		var search = window.location.search.substr(1).split('&');
		$.each(search, function(_, s) {
			var parts = s.split('=');
			if (parts[0] == 'email') {
				$("input[name='email']").val(parts[1]);
				$('#form-group-email').addClass("has-success");
			} else if (parts[0] == 'usernametaken') {
				$("input[name='username']").val(parts[1]);
				$('#form-group-username input').addClass("form-error");
				$('#form-group-username').prev('.elemen7s-font-fprint').append('<span style="color:red;">Username taken</span>');
			} else if (parts[0] == 'username') {
				$("input[name='username']").val(parts[1]);
				$('#form-group-username').addClass("has-success");
			} else if (parts[0] == 'passwordmatch') {
				$('#form-group-password1 input, #form-group-password2 input').addClass("form-error");
				$('#form-group-password1, #form-group-password2').prev('.elemen7s-font-fprint').append('<span style="color:red;">Password mismatch</span>');
			}
		});
	};
});
