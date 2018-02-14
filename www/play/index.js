$(function() {
	var app = {};
	var seatsready = Promise.Fake();
	var username, gameid, deckid = 1, games = {}, hand = {}, decks = [], canDeclareAttack;

	SE.widget.controlProperty($('#start-menu')[0], 'timer');
	SE.widget.controlProperty($('#elements-menu')[0], 'timer');
	SE.widget.controlProperty($('#game-menu-pass')[0], 'timer');

	$('#chat-menu').keydown(function(e) {
		if (e.keyCode == 13) {
			SE.websocket.send('chat', {
				'message':$(this).val()
			});
			$(this).val('');
		}
	});

	var updatetimer = function(me) {
		me.updatetimer = function() {
			if (me.timer > 0) me.timer--;
			window.setTimeout(me.updatetimer, 1000);
		};
		me.updatetimer();
	};
	updatetimer($('#game-menu-pass')[0]);

	window.newAlert = function(msg) {
		SE.widget.new('se-alert', msg).then(function(alert) {
			$('#alert-menu').append(alert);
		});
	}
	var eventStart = function(data) {
		var startMenu = $('#start-menu');
		var handMenu = $('#hand-menu');
		startMenu.slideDown();
		startMenu.css({opacity:1});
		startMenu[0].timer = data.timer;
		updatetimer($('#start-menu')[0]);
		handMenu.css({opacity:0});
		handMenu.css({opacity:1});
	};
	var eventGame = function(data) {
		var seats = $('#game-menu-seats');

		username = data.username;
		gameid = data.gameid;

		SE.widget.get('se-gameseat').then(function(seatFactory) {
			var seatsmap = {};
			$.each(data.seats, function(seatid, seatdata) {
				var seat = seatFactory(seatdata);
				seat.update(seatdata);
				seat.timer = data.timer;
				seat.turnphase = 'starting';
				seats.append(seat);
				seatsmap[seat.username] = seat;
			});
			seatsready.resolve(seatsmap);
		});

		SE.event.fire('reset-cards-in-hand', data.hand);
		var gameMenu = $('#game-menu');
		gameMenu.css({opacity:0.5}); // warm up transition prop
		gameMenu.slideDown();
		gameMenu.css({opacity:1});
	};
	var eventHand = function(data) {
		SE.event.fire('reset-cards-in-hand', data.cards);
	};
	var eventSunrise = function(data) {
		$('#start-menu').slideUp();
		$('#game-menu-meta-message')[0].innerHTML = 'sunrise';
		$('#game-menu-meta img')[0].src = '/img/icon/sunrise.1.64px.png';
		$('#game-menu-meta img')[4].src = '/img/icon/moon.0.64px.png';

		if (data.username == username) {
			var elementsMenu = $('#elements-menu');
			elementsMenu.fadeIn();
			elementsMenu[0].timer = data.timer;
			updatetimer(elementsMenu[0]);
		}

		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				if (seat.username == data.username) {
					seat.reactivate();
					seat.turnphase = 'sunrise';
					seat.timer = data.timer;
				} else {
					seat.turnphase = 'wait';
					seat.timer = 0;
				}
			});
		});
	};
	var eventMain = function(data) {
		$('#game-menu-pass')[0].timer = data.timer;
		$('#game-menu-meta-message')[0].innerHTML = 'main';
		$('#game-menu-meta img')[1].src = '/img/icon/sun.1.64px.png';
		$('#game-menu-meta img')[0].src = '/img/icon/sunrise.0.64px.png';
		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				seat.timer = data.timer;
				if (seat.username == data.username) {
					seat.turnphase = 'main';
					seat.life = data.life;
					seat.hand = data.hand;
					seat.deck = data.deck;
					seat.spent = data.spent;
				} else {
					seat.turnphase = 'respond(main)';
				}
			});
		});
	};
	var eventSunset = function(data) {
		$('#game-menu-pass')[0].timer = data.timer;
		$('#game-menu-meta-message')[0].innerHTML = 'night';
		$('#game-menu-meta img')[4].src = '/img/icon/moon.1.64px.png';
		$('#game-menu-meta img')[3].src = '/img/icon/sunset.0.64px.png';

		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				seat.timer = data.timer;
				if (seat.username == data.username) {
					seat.turnphase = 'sunset';
				} else {
					seat.turnphase = 'respond(sunset)';
				}
			});
		});
	};
	var eventPlay = function(data) {
		$('#game-menu-pass')[0].timer = data.timer;
		$('#game-menu-meta-stars img').each(function(img) {
			img.src = "/img/icon/stars.0.64px.png";
		});
		$('#game-menu-meta-stars').append($('<img class="disp-iblock" src="/img/icon/stars.1.64px.png"/>'));
		SE.gamecards.get(data.card.gcid, data.card.cardid).then(function(card) {
			SE.event.fire('play-react', data, card);
		});
		seatsready.then(function(seats) {
			$.each(seats, function(name, seat) {
				seat.timer = data.timer;
				seat.turnphase = 'response';
			});

			seats[data.username].hand = data.hand;

			SE.go(function() { // fix reload element duplicates
				seats[data.username].resetElements(data.elements);
			});
		});
	};
	var eventPass = function(data) {
		seatsready.then(function(seats) {
			seats[data.username].timer = 0;
		});

		$('#game-menu-meta-stars img').each(function(img) {
			console.log('pass event on images');
			img.src = "/img/icon/stars.0.64px.png";
		});

		var msg = (data.auto?'auto-pass':'pass')+(data.target?'('+data.target+')':'');

		newAlert({
			timeout:3000,
			class:'tip',
			username:data.username,
			message:msg
		});
	};
	var eventResolve = function(data) {
		console.log('event resolve', data);
	};
	var eventSpawn = function(data) {
		$('#play-menu').slideUp();
		$('#game-menu-pass')[0].timer = data.timer;
		seatsready.then(function(seats) {
			SE.widget.new('se-gc', data.card).then(function(gc) {
				seats[data.username].addActiveCard(gc);
			});
		});
	};
	var eventAttack = function(data) {
		$('#game-menu-pass')[0].timer = data.timer;
		$('#game-menu-meta-message')[0].innerHTML = 'attack';
		$('#game-menu-meta-stars').empty();
		$('#game-menu-meta img')[2].src = '/img/icon/combat.1.64px.png';
		$('#game-menu-meta img')[1].src = '/img/icon/sun.0.64px.png';
		canDeclareAttack = data.username == username;

		seatsready.then(function(seats) {
			$('.se-gc', seats[username]).each(function(i, gc) {
				if (data.attackoptions[gc.gcid]) {
					gc.showAttack();
				} else {
					gc.showClear();
				}
			});

			$.each(seats, function(i, seat) {
				if (seat.username == data.username) {
					seat.timer = data.timer;
					seat.turnphase = 'attack';
				} else {
					seat.timer = 0;
					seat.turnphase = 'wait';
				}
			});
		});
		$('#start-menu-timer').fadeOut();
		$('#game-menu').fadeIn();
	};
	var eventDefend = function(data) {
		$('#game-menu-pass')[0].timer = data.timer;
		$('#game-menu-meta-message')[0].innerHTML = 'defend';
		$('#game-menu-meta img')[3].src = '/img/icon/sunset.1.64px.png';
		$('#game-menu-meta img')[2].src = '/img/icon/combat.0.64px.png';

		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				if (seat.username == data.username) {
					seat.timer = data.timer;
					seat.turnphase = data.turnphase;
				} else {
					seat.timer = 0;
					seat.turnphase = 'wait';
				}
			});
		});
		$('#start-menu-timer').fadeOut();
		$('#game-menu').fadeIn();
	};
	var eventEnd = function(data) {
		console.warn('end', data);
		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				seat.timer = 0;
			});
			$.each(data.winners, function(i, name) {
				seats[name].turnphase = 'winner';
			});
			$.each(data.losers, function(i, name) {
				seats[name].turnphase = 'loser';
			});
		});
	};

	var websocketAlert = function(data) {
		SE.widget.new('se-alert', data).then(function(alert) {
			$(alert).css({display:'none'}); // warm up transition prop
			$('#alert-menu').append(alert);
			$(alert).slideDown();
		});
	};
	var websocketAnimate = function(data) {
		if (data.animate == 'not enough elements') {
			SE.widget.new('se-alert').then(function(alert) {
				$(alert).css({display:'none'});
				alert.message = 'not enough elements';
				alert.setMode('danger');
				alert.autoDismissSeconds(2);
				$('#alert-menu').append(alert);
				$(alert).slideDown();
			});
		} else if (data.animate == 'add card') {
			SE.gamecards.get(data.gcid, data.cardid).then(function(card) {
				SE.event.fire('add-card-to-hand', card);
			}, function(err) {
				console.error('websocket animate add card', err);
			});
		} else if (data.animate == 'add element') {
			seatsready.then(function(seats) {
				$.each(seats, function(i, seat) {
					if (seat.username != data.username) return true;
					seat.addElement(data.element, true);
					return false;
				});
			});
		} else {
			console.log('websocket animate not recognized', data.animate);
		};
	};
	var websocketGameDone = function(data) {
		$('#game-menu, #elements-menu, #start-menu').slideUp();
		$('#done-menu').slideDown();
		$.each(data.winners, function(i, name) {
			var span = $('<h3>'+name+'</h3>');
			console.log('list winner', name, span);
			$('#done-menu-winners').append(span);
		});
		$.each(data.losers, function(i, name) {
			console.log('list loser', name);
			var span = $('<h3>'+name+'</h3>');
			$('#done-menu-losers').append(span);
		});
	};
	var websocketTimer = function(data) {
		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				if (seat.username == data.username) {
					seat.timer = data.timer;
				}
			});
		});
	};
	var websocketAttack = function(data) {
		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				if (seat.username == data.username) {
					console.log('websocket attack found person', data.username);
					$.each($('.se-gameseat-active .se-card', seat), function(i, card) {
						if (card.gcid == data.gcid) {
							console.log('websocket attack found gcid', data.gcid);
							card.banner = 'Attack';
							return false;
						}
					});
					return false;
				}
			});
		});
	};
	SE.event.on('reset-cards-in-hand', function(data) {
		console.log('reset-cards-in-hand', data);

		$('#hand-menu .se-card-spinner')[0].empty();
		$.each(data, function(i, carddata) {
			SE.gamecards.get(carddata.gcid, carddata.cardid).then(function(card) {
				SE.event.fire('add-card-to-hand', card);
			});
		});
	});
	SE.event.on('add-card-to-hand', function(card) {
		$(card).click(function(e) {
			SE.event.fire('play-confirm', card);
		});
		$('#hand-menu .se-card-spinner')[0].append(card);
	});

	// play dialog
	SE.event.on('play-confirm', function(c) {
		SE.widget.new('se-card', c.cardid).then(function(card) {
			var playMenu = $('#play-menu');
			playMenu.css({opacity:0.5});
			playMenu.slideDown();
			playMenu.css({opacity:1});
			$('[data-ctrl="title"]', playMenu)[0].innerHTML = "Examine "+card.name;
			$('[data-ctrl="content"]', playMenu).empty();
			$('[data-ctrl="content"]', playMenu).append(card);

			$('button', playMenu).off('click');
			$('button', playMenu).on('click', function() {
				SE.event.fire('play-hide');
			});
			$('[data-ctrl="play"]', playMenu).slideDown();
			$('[data-ctrl="play"]', playMenu).click(function() {
				SE.websocket.send('play', {
					gameid: gameid,
					gcid: parseInt(c.gcid)
				});
			});
		});
	});
	SE.event.on('play-react', function(data, card) {
		console.log('play-react', data);
		var playMenu = $('#play-menu');
		$('[data-ctrl="content"]', playMenu).empty();
		playMenu.css({opacity:0.5});
		playMenu.slideDown();
		playMenu.css({opacity:1});
		$('[data-ctrl="title"]', playMenu)[0].innerHTML = 'Play: '+data.username;;
		$('[data-ctrl="content"]', playMenu).append(card);

		$('button', playMenu).off('click');
		$('button', playMenu).on('click', function() {
			SE.event.fire('play-hide');
		});
		$('[data-ctrl="play"]', playMenu).slideUp();
	});
	SE.event.on('play-hide', function() {
		$('#play-menu').css({opacity:0.5});
		$('#play-menu').slideUp();
	});
	// end play dialog

	SE.event.on('se-gc-click', function(gc) {
		if (canDeclareAttack && gc.username == username) {
			console.warn('se-gc-click attack!', gc);
			SE.websocket.send('attack', {
				gameid: gameid,
				gcid: gc.gcid
			});
		} else {
			console.warn('se-gc-click', gc);
		}
	});

	SE.event.on('websocket.message', function(name, data) {
		if (name == 'start') {
			eventStart(data);
		} else if (name == 'sunrise') {
			eventSunrise(data);
		} else if (name == 'game') {
			eventGame(data);
		} else if (name == 'hand') {
			eventHand(data);
		} else if (name == 'main') {
			eventMain(data);
		} else if (name == 'sunset') {
			eventSunset(data);
		} else if (name == 'play') {
			eventPlay(data);
		} else if (name == 'resolve') {
			eventResolve(data);
		} else if (name == 'spawn') {
			eventSpawn(data);
		} else if (name == 'pass') {
			eventPass(data);
		} else if (name == 'gamedone') {
			websocketGameDone(data);
		} else if (name == 'attack') {
			eventAttack(data);
		} else if (name == 'defend') {
			eventDefend(data);
		} else if (name == 'timer') {
			websocketTimer(data);
		} else if (name == 'attack') {
			websocketAttack(data);
		} else if (name == 'animate') {
			websocketAnimate(data);
		} else if (name == 'alert') {
			websocketAlert(data);
		} else if (name == 'end') {
			eventEnd(data);
		} else {
			console.warn('websocket event not recognized', name);
		}
	});

	// elements menu click triggers
	var makeSendElementTrigger = function(elementid) {
		return function() {
			SE.websocket.send('element', {
				gameid:gameid,
				elementid:elementid
			});
		}
	};
	for (var i=1;i<8;i++) {
		$('#elements-menu-'+i).click(makeSendElementTrigger(i));
	}
	$('#elements-menu button').click(function() {
		$('#elements-menu').fadeOut();
	});

	// start menu click triggers
	$('#start-button-keep').click(function(){
		SE.websocket.send('hand', {
			gameid:gameid,
			choice: 'keep'
		});
		$('#start-menu').fadeOut();
	});
	$('#start-button-mulligan').click(function(){
		SE.websocket.send('hand', {
			gameid:gameid,
			choice:'mulligan'
		});
		$('#start-menu').fadeOut();
	});

	// pass button
	$('#game-menu-pass').click(function() {
		SE.websocket.send('pass', {
			gameid:gameid
		});
		this.timer = 0;
	});

	// so let's go then
	vii.ping().then(function(data) {
		if (!data.username) return console.warn('login required');
		app.data = data;

		var search = window.location.search.substr(1).split('=');
		if (search[0] == 'gameid') {
			var gdat = data.games[search[1]];
			if (gdat) {
				SE.websocket.send('join', {
					gameid:search[1]
				});
			} else {
				console.warn("apparently you are trying to do spectate mode fool")
			}
		}
	});
});
