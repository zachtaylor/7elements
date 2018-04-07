SE.widget.control('se-app-search', function() {
	var me = this;
	me.deckselect = 1;

	me.spinner = Promise.Fake();

	SE.widget.new('se-card-spinner').then(function(spinner){
		$(me['spinner-replace']).replaceWith(spinner);
		me.spinner.resolve(spinner);
	});

	me.append = function(card) {
		me.spinner.then(function(spinner) {
			spinner.append(card);
		});
	};
	me.prepend = function(card) {
		me.spinner.then(function(spinner) {
			spinner.prepend(card);
		});
	};
	me.empty = function() {
		me.spinner.then(function(spinner) {
			spinner.empty();
		});
	};
	me.update = function(data) {
		console.log('se-app-search update', data);
		if (data) me.data = data;
		me.deckid = me.deckselect;
		me.deckname = me.data[me.deckselect].name;
		me.deckcount = Object.keys(me.data).length;
		me.deckwins = me.data[me.deckselect].wins;

		me.empty();
		$.each(me.data[me.deckselect].cards, function(cardid, cardcount) {
			SE.widget.new('se-card', cardid).then(function(card) {
				card.banner = cardcount + ' copies';
				me.append(card);
			});
		});
	};

	$(me.$deckleft).click(function() {
		if (me.deckselect > 1) me.deckselect--;
		else me.deckselect = me.deckcount - 1;
		me.update(me.data);
	});
	$(me.$deckright).click(function() {
		if (me.deckselect < me.deckcount) me.deckselect++;
		else me.deckselect = 1;
		me.update(me.data);
	});

	// SE.event.on('animate-game-found', function() {
	// 	SE.dirtypage.off();
	// 	$('#menu-queue-progress').removeClass('progress-bar-warning').removeClass('progress-bar-striped').addClass('progress-bar-success');
	// 	$('#search-progress span:first-child')[0].innerHTML = 'Game found!';
	// 	$('#search-progress')[0].timer = '';
	// 	$('#menu-queue-cancel').fadeOut();
	// });

	var vai = $('<button class="vii vii-blue" style="margin-right:7px;"><img src="/img/icon/use.20px.png">Play now vs A.I.</button>');
	vai.click(function() {
		SE.websocket.send('newgame', {
			deckid: me.deckselect,
			ai:true
		});
		window.location.href = '/#games';
	});
	var pvp = $('<button class="vii vii-blue" style="margin-right:7px;"><img src="/img/icon/timer.20px.png">Search for a human</button>');
	pvp.click(function() {
		SE.websocket.send('newgame', {
			deckid: me.deckselect
		});
		window.location.href = '/#games';
	});
	me.footer = $('<div></div>')[0];
	$(me.footer).append(vai);
	$(me.footer).append(pvp);

	vii.ping().then(function(data) {
		if (data.decks) me.update(data.decks);
		else $('#content').append('<h2>Search Game: <a href="/login/">Login</a> Required</h2>')
	});
});
