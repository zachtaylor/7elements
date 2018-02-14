// # vii.ping
// returns a promise to find ping data >= 1 time
vii.ping = function() {
	if (!vii.ping.then) {
		vii.ping.chain = [];
		vii.ping.then = function(f) {
			vii.ping.chain.push(f);
			if (vii.ping.data) f(vii.ping.data);
		};
		vii.ping.trigger = function() {
			$.each(vii.ping.chain, function(i, f) {
				f(vii.ping.data);
			});
		};
		SE.event.on('websocket.message', function(name, data) {
			vii.ping.data = data;
			if (name == '/ping') vii.ping.trigger();
		});
		SE.websocket.send('ping');
	}
	return vii.ping;
};
