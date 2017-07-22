SE.api = {
	cache: {},
	_promise: function(eventName, path) {
		return new Promise(function(resolve, reject) {
			$.getJSON(path).done(function(data) {
				resolve(data);
				SE.event.fire(eventName, data);
			}).fail(reject);
		});
	},
	get: function(name, f) {
		var eventName = 'data.'+name;
		var path = '/api/'+name+'.json';
		if (f) {
			SE.event.on(eventName, f);
			return SE.api.get(name);
		};
		if (!SE.api.cache[name]) SE.api.cache[name] = SE.api._promise(eventName, path);
		return SE.api.cache[name];
	},
	post: function(name, data) {
		var path = '/api/'+name+'.json';
		return new Promise(function(resolve, reject) {
			console.log('POST ', path, data);
			$.post(path, data).done(resolve).fail(reject);
		});
	}
};
