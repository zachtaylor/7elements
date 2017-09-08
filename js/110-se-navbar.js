SE.widget.control('se-navbar', function() {
	var navbar = this;
	var pathname = window.location.pathname;

	SE.widget.controlProperty(navbar, 'packs');
	SE.widget.controlProperty(navbar, 'username');
	SE.widget.controlProperty(navbar, 'games');

	if (pathname == '/') {
		$('#se-navbar-header', this).addClass('active');
	} else if (pathname == '/cards/') {
		SE.widget.controlProperty(navbar, 'cards');
		$('#se-navbar-link-cards', this).addClass('active');
	} else if (pathname == '/packs/') {
		$('#se-navbar-link-packs', this).addClass('active');
	} else if (pathname.substr(0, 7) == '/decks/') {
		$('#se-navbar-link-decks', this).addClass('active');
	} else if (pathname.substr(0, 7) == '/games/') {
		$('#se-navbar-link-games', this).addClass('active');
	} else if (pathname.substr(0, 11) == '/myaccount/') {
		$('#se-navbar-link-myaccount', this).addClass('active');
	}

	SE.api.get('navbar').then(function(data) {
		navbar.username = data.username;
		navbar.packs = data.packs;
		navbar.games = Object.keys(data.games).length;
		$('#se-navbar-link-packs, #se-navbar-link-decks, #se-navbar-link-games, #se-navbar-link-myaccount', navbar).css({
			display:'block'
		});
		$('#se-navbar-link-login', navbar).hide();
	}, function(e) {if (e.responseText != 'session missing') console.error(e.responseText)});
});
