$(function() {
	var seatsready = Promise.Fake();
	var username, gameid, deckid = 1, games = {}, hand = {}, decks = [], attacks = {};

	SE.widget.controlProperty($('#start-menu')[0], 'timer');
	SE.widget.controlProperty($('#game-menu-pass')[0], 'timer');

	var updatetimer = function(me) {
		me.updatetimer = function() {
			if (me.timer > 0) me.timer--;
			window.setTimeout(me.updatetimer, 1000);
		};
		me.updatetimer();
	};
	updatetimer($('#start-menu')[0]);
	updatetimer($('#game-menu-pass')[0]);
	var makeSendAttackTrigger = function(gcid) {
		return function() {
			console.log('send attack', gcid);
			SE.websocket.send('attack', {
				gcid: gcid
			});
		}
	};
	var makeSendDefendTrigger = function(gcid) {
		console.log('make send defend trigger', gcid);
		return function() {
			var defendButton = $('#game-menu-defend');
			console.log('fire defend ', gcid, defendButton);

			// SE.websocket.send('defend', {
			// 	gcid: gcid
			// });
		}
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
		} else if (data.animate == 'draw card') {
			SE.gamecards.get(data.card.gcid, data.card.cardid).then(function(card) {
				SE.event.fire('add-card-to-hand', card);
			}, function(err) {
				console.error('websocket animate draw card', err);
			})
		} else {
			console.log('websocket animate not recognized', data.animate);
		}
	};
	var websocketGameStart = function(data) {
		console.log('gamestart', data);
		gameid = data.gameid;
		$('#start-menu').slideDown();
		$('#start-menu-hand').empty();
		$.each(data.hand, function(i, carddata) {
			SE.widget.new('se-card', carddata.cardid).then(function(card) {
				$('#start-menu-hand').append(card);
			});
		});
		$('#start-menu')[0].timer = data.timer;
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
	var websocketGame = function(data) {
		var seats = $('#game-menu-seats')[0];

		username = data.username;
		gameid = data.gameid;

		$.each(data.seats, function(seatid, seatdata) {
			SE.widget.new('se-gameseat').then(function(seat) {
				if (!seats.children[seatid]) $(seats).append(seat);
				seat.update(seatdata);
				seat.timer = 0;
				seat.turnphase = 'wait';
			});
		});

		SE.event.fire('reset-cards-in-hand', data.hand);
		$('#start-menu').slideUp();
		$('#game-menu').slideDown();
		window.setTimeout(function() {
			seatsready.resolve($('#game-menu-seats')[0].children);
		}, 1000);
	};
	var websocketPlay = function(data) {
		$('#game-menu-pass')[0].timer = data.timer;
		SE.gamecards.get(data.gcid, data.cardid).then(function(card) {
			$.each($('#game-menu-seats')[0].children, function(i, seat) {
				seat.timer = data.timer;
				seat.turnphase = 'respond';
				if (seat.username == data.username) {
					$(card).detach().off('mouseover').off('mouseout');
					$(card).css({top:'0px',opacity:'0.33', display:'none'});
					$('.se-gameseat-active .se-card-spinner', seat)[0].append(card);
					$(card).slideDown();
				}
			});
		});
	};
	var websocketTurn = function(data) {
		if (data.username == username) {
			$('#game-menu-pass')[0].timer = data.timer;
			if (!data.element) $('#elements-menu').fadeIn();
		}

		seatsready.then(function(seats) {
			$.each(seats, function(i, seat) {
				if (seat.username == data.username) {
					$(seat).addClass('active');
					seat.timer = data.timer;
					seat.turnphase = data.turnphase;
				} else {
					$(seat).removeClass('active');
					seat.timer = 0;
					seat.turnphase = 'wait';
				}

				if (seat.turnphase == 'begin') {
					console.log('reactivating seat', seat);
					seat.reactivate();
				} else if (seat.turnphase == 'attack') {
					$('.se-card', seat).off('click');
					$.each($('.se-card', seat), function(i, card) {
						$(card).off('click').click(makeSendAttackTrigger(card.gcid));
					});
				} else if (seat.turnphase == 'defend') {
					$('.se-card', seat).off('click');
					$.each($('.se-card', seat), function(i, card) {
						$(card).off('click').click(makeSendDefendTrigger(card.gcid));
					});
				} else if (seat.turnphase == 'done') {
					$('.se-card').each(function(i, card) {
						card.banner = '';
					});
				}
			});
		});

		$('#start-menu-timer').fadeOut();
		$('#game-menu').fadeIn();
	};
	var websocketElements = function(data) {
		$.each($('#game-menu-seats')[0].children, function(i, seat) {
			if (seat.username == data.username) {
				seat.updateelements(data.elements);
				return false;
			}
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
		console.log('websocket attack');
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
	var websocketRespond = function(data) {
		$.each($('#game-menu-seats')[0].children, function(i, seat) {
			if (seat.username == data.username) {
				seat.turnphase = 'wait';
				seat.timer = 0
			}
		});
	};
	SE.event.on('update-games-menu', function() {
		console.log('update-games-menu', decks, deckid);

		$('#ready-menu-list').empty();
		$.each(games, function(gameid, gamedata) {
			console.log('found game data', gameid, gamedata, $('#ready-menu-list')[0]);
		});
	});
	SE.event.on('reset-cards-in-hand', function(data) {
		console.log('reset-cards-in-hand', data);

		$('#game-menu-hand .se-card-spinner')[0].empty();
		$.each(data, function(i, carddata) {
			SE.gamecards.get(carddata.gcid, carddata.cardid).then(function(card) {
				SE.event.fire('add-card-to-hand', card);
			});
		});
	});
	SE.event.on('add-card-to-hand', function(card) {
		console.log('add-card-to-hand gcid=', card.gcid, 'cardid=', card.cardid);
		$(card).mouseup(function() {
			SE.websocket.send('play', {
				gameid: gameid,
				gcid: parseInt(card.gcid)
			});
		});
		$(card).css({top:'0px'}); // warm up css transition property
		$(card).mouseover(function() {
			$(card).css({top:'-250px'});
		});
		$(card).mouseout(function() {
			$(card).css({top:'0px'});
		});
		$('#game-menu-hand .se-card-spinner')[0].append(card);
	});

	SE.event.on('websocket.message', function(name, data) {
		console.log('websocket message', name, data);
		if (name == 'gamestart') {
			websocketGameStart(data);
		} else if (name == 'gamedone') {
			websocketGameDone(data);
		} else if (name == 'game') {
			websocketGame(data);
		} else if (name == 'turn') {
			websocketTurn(data);
		} else if (name == 'play') {
			websocketPlay(data);
		} else if (name == 'elements') {
			websocketElements(data);
		} else if (name == 'respond') {
			websocketRespond(data);
		} else if (name == 'timer') {
			websocketTimer(data);
		} else if (name == 'attack') {
			websocketAttack(data);
		} else if (name == 'animate') {
			websocketAnimate(data);
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

	$('#start-button-keep').click(function(){
		SE.websocket.send('hand', {
			gameid:gameid,
			choice: 'keep'
		});
		$('#start-menu button').fadeOut();
	});
	$('#start-button-mulligan').click(function(){
		SE.websocket.send('hand', {
			gameid:gameid,
			choice:'mulligan'
		});
		$('#start-menu button').fadeOut();
	});
	$('#elements-menu button').click(function() {
		$('#elements-menu').fadeOut();
	});
	$('#game-menu-pass').click(function() {
		SE.websocket.send('pass', {
				gameid:gameid
			});
		this.timer = 0;
	});

	// lets go then
	SE.websocket.open().then(function() {
		console.log('websocket opened!');
		var search = window.location.search.substr(1).split('=');
		if (search[0] == 'gameid' && search[1] > 0) {
			SE.websocket.send('join', {
				gameid:parseInt(search[1])
			});
		}
	});
});
