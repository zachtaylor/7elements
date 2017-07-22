SE.widget.control('se-navbar', function() {
	var navbar = this;
	var pathname = window.location.pathname;

	if (pathname == '/') {
		$('#se-navbar-header', this).addClass('active');
	} else if (pathname == '/cards/') {
		$('#se-navbar-link-cards', this).addClass('active');
	} else if (pathname == '/packs/') {
		$('#se-navbar-link-packs', this).addClass('active');
	} else if (pathname.substr(0, 7) == '/decks/') {
		$('#se-navbar-link-decks', this).addClass('active');
	} else if (pathname.substr(0, 11) == '/myaccount/') {
		$('#se-navbar-link-myaccount', this).addClass('active');
	}

	SE.api.get('myaccount', function(data) {
		$('#se-navbar-link-myaccount', navbar)[0].innerHTML = 'Account:'+data.username;

		if (data.packs) {
			$('#se-navbar-link-packs', navbar)[0].innerHTML = 'Packs ('+data.packs+')';
		} else {
			$('#se-navbar-link-packs', navbar)[0].innerHTML = 'Packs';
		}
	}).then(null, function() {
		$('#se-navbar-link-packs, #se-navbar-link-decks, #se-navbar-link-myaccount', navbar).not('.active').hide();
		$('#se-navbar-link-login', navbar).show();
	});
});
