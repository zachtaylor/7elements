SE.widget.control('se-app-chat', function() {
	var me = this;
	vii.ping();

	SE.event.on('chat-search-start', function() {
		var username = prompt("Please enter username", "");
		if (username) {
			SE.websocket.send('chat', {
				channel: username,
				message: 'hi',
			});
		}
	});

	me.channels = {};
	me.getChannel = function(name) {
		if (me.channels[name]) return me.channels[name]
		me.channels[name] = {
			tab: $('<li>'+name+'</li>')[0],
			history: $('<div class="pos-abs bottom0" style="max-height:100%;width:100%;overflow:auto;"></div>')[0],
		}
		$(me.$channels).append(me.channels[name].tab);
		return me.channels[name]
	};
	me.focusChannel = function(name) {
		me.channel = name;
		if (me.history) $(me.history).detach();
		me.history = me.getChannel(name).history;
		$(me.$channelhistory).append(me.history);
	};

	SE.event.on('/chat', function(data) {
		SE.widget.new('se-chat-line', data).then(function(html) {
			$(me.getChannel(data.channel).history).append(html);
			vii.sound.play('chat');
		});
	});

	var input = $('<input class="vii width100"/>')[0];
	$(input).keyup(function(e) {
		if (e.which == 13) {
			SE.websocket.send('chat', {
				channel: me.channel,
				message: $(input).val()
			});
			$(input).val('');
		};
	});
	me.footer = input;

	$(me).click(function(e) {
		$(input).focus();
	});

	// 'all' channel
	me.focusChannel('all');
});
