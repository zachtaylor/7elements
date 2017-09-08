SE.websocket = {
	websocket: null,
	promise: null,
	open: function() {
		if (!SE.websocket.promise)
			SE.websocket.promise = new Promise(function(resolve, reject) {
				SE.websocket.websocket = new WebSocket(window.location.protocol.replace('http', 'ws')+window.location.host+'/api/websocket');
				SE.websocket.websocket.onopen = function() {
					resolve(SE.websocket.websocket);
					SE.event.fire('websocket.open');
				}
				SE.websocket.websocket.onmessage = function(e) {
					var data = JSON.parse(e.data);
					SE.event.fire('websocket.message', data.name, data.data);
				};
				SE.websocket.onclose = function(e) {
					console.warn('websocket connection lost', e);
				};
			});
		return SE.websocket.promise
	},
	send: function(name, data) {
		var message = {'name':name};
		if (data) message.data = data;

		SE.websocket.open().then(function(websocket) {
			websocket.send(JSON.stringify(message));
		});
	}
};