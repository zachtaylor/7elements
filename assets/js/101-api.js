SE.api = {
	cache: {},
	_promise: function(eventName, path) {
		return new Promise(function(resolve, reject) {
			$.getJSON(path).done(function(data) {
				resolve(data);
				SE.event.fire(eventName, data);
			}).catch(false).fail(reject);
		});
	},
	get: function(name, f) {
		var eventName = 'data.'+name;
		var path = '/api/'+name+'.json';
		if (f) {
			SE.event.on(eventName, f);
			return SE.api.get(name);
		};
		SE.api.cache[name] = SE.api.cache[name] || SE.api._promise(eventName, path);
		return SE.api.cache[name];
	},
	uncacheGet: function(name, f) {
		SE.api.cache[name] = false;
		return SE.api.get(name, f);
	},
	post: function(name, data) {
		var path = '/api/'+name+'.json';
		return new Promise(function(resolve, reject) {
			$.post(path, data).done(resolve).fail(reject);
		});
	}
};
