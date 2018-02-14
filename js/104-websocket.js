SE.websocket = {
	websocket: null,
	promise: null,
	open: function() {
		if (!SE.websocket.promise)
			SE.websocket.promise = new Promise(function(resolve, reject) {
				SE.websocket.websocket = new WebSocket(window.location.protocol.replace('http', 'ws')+window.location.host+'/api/websocket');
				SE.websocket.websocket.onopen = function() {
					console.debug('websocket opened');
					resolve(SE.websocket.websocket);
					SE.event.fire('websocket.open');
				}
				SE.websocket.websocket.onmessage = function(e) {
					var data = JSON.parse(e.data);
					console.debug('websocket.message', data.uri, data.data);
					SE.event.fire('websocket.message', data.uri, data.data);
				};
				SE.websocket.onclose = function(e) {
					console.warn('websocket connection lost', e);
					SE.event.fire('websocket.close');
				};
			});
		return SE.websocket.promise
	},
	send: function(name, data) {
		var message = {'uri':'/'+name};
		if (data) message.data = data;
		else message.data = {};

		SE.websocket.open().then(function(websocket) {
			console.debug('websocket.send', message.uri, message.data);
			websocket.send(JSON.stringify(message));
		});
	}
};
