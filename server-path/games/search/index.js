$(function() {
	var deckid = 1, data = {};

	var startcountup = function(me) {
		SE.widget.controlProperty(me, 'timer');
		me.countup = function() {
			if (me.timer > 0) me.timer++;
			window.setTimeout(me.countup, 1000);
		};
		me.countup();
	};

	startcountup($('#search-progress')[0]);

	SE.api.get('navbar').then(function(navbardata) {
		data.navbar = navbardata;
		SE.event.fire('update');
	});

	SE.event.on('update', function() {
		if (!data.navbar) return;
		var deck = data.navbar.decks[deckid];
		$('.se-deck-details')[0].update(deck);
	});

	$('#deck-left').click(function() {
		deckid--;
		if (deckid == 0) {
			deckid = Object.keys(data.navbar.decks).length;
		}
		SE.event.fire('update');
	});
	$('#deck-right').click(function() {
		deckid++;
		if (!data.navbar.decks[deckid]) {
			deckid = 1;
		}
		SE.event.fire('update');
	});

	$('#deck-ready').click(function() {
		SE.event.fire('animate-search-start');
		$.getJSON('/api/newgame.json', {
			deckid: deckid
		}).done(function(data) {
			SE.event.fire('game-found', data.gameid);
			SE.event.fire('animate-game-found');
		}).fail(function(e) {
			SE.event.fire('animate-search-fail');
		});
	});

	SE.event.on('game-found', function(gameid) {
		window.setTimeout(function() {
			window.location.href = ('/games/play/?gameid='+gameid)
		}, 5000);
	});

	SE.event.on('animate-search-fail', function(data) {
		console.log('search-fail');
	});

	SE.event.on('animate-search-start', function() {
		$('#menu-search .footer').hide();
		$('#search-progress').slideDown();
		$('#search-progress')[0].timer = 1;
		SE.dirtypage.on();
	});

	SE.event.on('animate-game-found', function() {
		SE.dirtypage.off();
		$('#menu-queue-progress').removeClass('progress-bar-warning').removeClass('progress-bar-striped').addClass('progress-bar-success');
		$('#search-progress span:first-child')[0].innerHTML = 'Game found!';
		$('#search-progress')[0].timer = '';
		$('#menu-queue-cancel').fadeOut();
	});
});
