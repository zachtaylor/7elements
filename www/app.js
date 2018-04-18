$(function() {
	var app = {};
	app.drawer = {
		open: function() {
			$('#nav-drawer').css({left:'0px'});
			app.drawer.isopen = true;
		},
		close: function() {
			$('#nav-drawer').css({left:'-'+(10+$('#nav-drawer').width())+'px'});
			app.drawer.isopen = false;
		},
		toggle: function() {
			if (app.drawer.isopen) app.drawer.close();
			else app.drawer.open();
		}
	};
	app.chat = {
		open: function() {
			$('#nav-chat').slideDown();
			$('#nav-chat input').focus();
			app.chat.isopen = true;
		},
		close: function() {
			$('#nav-chat').slideUp();
			app.chat.isopen = false;
		},
		toggle: function() {
			if (app.chat.isopen) app.chat.close();
			else app.chat.open();
		}
	};
	SE.event.on('nav-drawer', function() {
		app.drawer.toggle();
	});
	SE.event.on('nav-chat', function() {
		if (!app.data) {
			return console.warn('you must login to access chat');
		} else if ($('#nav-chat input').val()) {
			$('#nav-chat input').focus();
			return;
		}
		app.chat.toggle();
	});

	app.keybind = {
		appdrawer: 27,
		chat:13,
	};
	app.keybind.on = function(f) {
		$.each(app.keybind, function(action, key) {
			f(key, action);
		});
		SE.event.on('keybind', function(key, action) {
			f(key, action);
		});
	};
	app.keybind.set = function(key, action) {
		if (key != 13 && key != 27) {
			return console.warn('keybind key error', key);
		} else if (action != 'appdrawer' && action != 'chat') {
			return console.warn('keybind unrecognized action', action);
		}
		app.keybind[action] = key;
		SE.event.fire('keybind', key, action);
	};
	$('body').keyup(function(e) {
		if (e.which == app.keybind.appdrawer) SE.event.fire('nav-drawer');
		if (e.which == app.keybind.chat) SE.event.fire('nav-chat');
	});
	$('#nav-chat input').keyup(function(e) {
		e.stopPropagation();
		if (e.which == 13) {
			var message = $(this).val();
			if (message) {
				SE.websocket.send('chat', {
					channel: $('#nav-chat-title span')[0].innerHTML,
					message: $(this).val()
				});
				$(this).val('');
			} else {
				app.chat.close();
			};
		};
	});

	$('#content').click(function(e) {
		app.drawer.close();
	});

	window.onhashchange = function() {
		SE.event.fire('nav-'+app.hash());
		$('#nav-path')[0].innerHTML = '/'+app.hash();
	};
	app.hash = function(v) {
		if (v) window.location.hash='#'+v;
		else if (v === "") window.location.hash="";
		else return window.location.hash.substr(1);
	};

	$(window).resize(function() {
		var footerHeight = $('#footer').height();
		$('#content').css({
			'max-height':(window.innerHeight-64-footerHeight)+'px'
		});
	});

	app.getview = function(name) {
		app.viewcache = app.viewcache || {};
		if (app.viewcache[name]) return app.viewcache[name];
		app.viewcache[name] = new Promise(function(resolve, reject) {
			console.debug('app.build:',name);
			SE.widget.new('se-app-'+name, app).then(function(widget) {
				resolve(widget);
			}, function() {
				reject();
			});
		});
		return app.viewcache[name];
	};
	app.gotoview = function(name) {
		if (app.view == name) return app.drawer.close();
		if (SE.dirtypage.state()) return;
		console.debug('app.goto:',name);
		app.view = name;
		app.hash(name);
		SE.event.fire('nav-reset');
		app.drawer.close();
		app.getview(name).then(function(view) {
			$('#content').append(view);
			if (view.footer) $('#footer').append(view.footer);
			if (view.reload) view.reload();
			$(window).resize();
		});
	};
	app.declareview = function(name) {
		SE.event.on('nav-'+name, function() {
			app.gotoview(name);
		});
	};

	app.declareview('home');
	app.declareview('learn');
	app.declareview('login');
	app.declareview('signup');
	app.declareview('decks');
	app.declareview('edit');
	app.declareview('games');
	app.declareview('play');
	app.declareview('search');
	app.declareview('cards');
	app.declareview('settings');
	app.declareview('packs');
	app.declareview('patch-notes');

	SE.event.on('/ping', function(data) {
		if (data&&data.username) app.data = data;
		else {
			document.cookie = 'SessionId=0;path=/;'
			$('#nav-top-login').show();
			return;
		};
		var cc = 0;
		$.each(app.data.cards, function(id, carddata) {
			if (carddata.copies > 0) cc += carddata.copies;
		});
		$('#nav-top-login').hide();
		$('#nav-top-account').show();
		$('nav [data-ctrl="account-cards"]')[0].innerHTML = cc;
		$('nav [data-ctrl="account-username"]')[0].innerHTML = app.data.username; 
		$('nav [data-ctrl="account-decks"]')[0].innerHTML = Object.keys(app.data.decks).length;
		$('nav [data-ctrl="account-packs"]')[0].innerHTML = app.data.packs;
		$('nav [data-ctrl="account-coins"]')[0].innerHTML = app.data.coins;
		$('nav [data-ctrl="account-games"]')[0].innerHTML = Object.keys(app.data.games).length;
		SE.event.fire('nav-'+app.hash());
	});
	SE.event.on('/chat', function(data) {
		$('#nav-top-account img')[0].src='/img/icon/chat.green.128px.png';
		vii.sound.play('chat');
		SE.event.fire('/notification', {
			class:'chat',
			title:data.username,
			message:data.message,
			timeout:7000,
		});
	});
	SE.event.on('/notification', function(data) {
		SE.widget.new('se-app-notification', data).then(function(notification) {
			$('#nav-notifications').append(notification);
		});
	});

	SE.event.on('websocket.close', function() {
		console.error('websocket closed');
		$('#nav-top-account img')[0].src='/img/icon/chat.red.128px.png';
	});

	SE.event.on('nav-reset', function() {
		$.each($('#content')[0].children, function(i, html) {
			$(html).detach();
		});
		$.each($('#footer')[0].children, function(i, html) {
			$(html).detach();
		});
	});

	SE.event.on('nav-coins', function() {
		app.drawer.close();
		if (app.view == 'coins') return;
		SE.event.fire('nav-reset');

		SE.widget.new('se-buy-menu').then(function(menu) {
			if (app.data) menu.coins = app.data.coins;
			else menu.coins = 0;
			$('#content').append(menu);
		});
		app.view = 'coins';
		app.hash('coins');
	});

	SE.event.on('buy-coins', function() {
		SE.api.get('buycoins').then(function(data) {
			SE.api.cache.buycoins = null;
			app.data.coins += 10;
			vii.ping();
		});
	});

	SE.event.on('buy-pack', function() {
		SE.api.get('buypack').then(function(data) {
			SE.api.cache.buypack = null;
			app.data.packs++;
			app.data.coins -= 7;
			vii.ping();
		});
	});

	SE.event.on('/match', function(data) {
		vii.ping.data.games[data.gameid] = data;
		SE.event.fire('/notification', {
			class:'match',
			title:'Match Found',
			message:data.gameid,
			timeout:7000,
		});
	});

	// PIPE
	SE.event.on('websocket.message', function(name, data) {
		SE.event.fire(name, data);
	});

	if (vii.cookie('SessionId')) vii.ping();
	else console.warn('no session found');
	if (app.hash()) window.onhashchange();
	else app.hash('home');
});
