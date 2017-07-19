SE.req = {
	cache: {},
	// proto
	get: function(name, subscriber) {
		var eventName = 'data.'+name;
		var path = '/api/'+name+'.json';

		if (subscriber) {
			SE.event.on(eventName, subscriber);
			return SE.req.get(name);
		}
		if (!SE.req.cache[name]) {
			SE.req.cache[name] = SE.req.promiseData(eventName, path)
		}
	},
	promiseData: function(eventName, path) {
		return new Promise(function(resolve, reject) {
			console.log('GET', path);
			$.getJSON(path).done(function(data) {
				resolve(data);
				SE.event.fire(eventName, data);
			}).fail(reject);
		});
	},
	post: function(name, data) {
		var path = '/api/'+name+'.json';
		return new Promise(function(resolve, reject) {
			console.log('POST ', path, data);
			$.post(path, data).done(resolve).fail(reject);
		});
	},
	// proto
	getmyaccount: function(subscriber) {
		if (subscriber) {
			SE.event.on('data.myaccount', subscriber);
			return SE.req.getmyaccount();
		}

		if (!SE.req.cache.myaccount) {
			console.log('request /api/myaccount.json');
			SE.req.cache.myaccount = new Promise(function(resolve, reject) {
				$.getJSON('/api/myaccount.json').done(function(data) {
					resolve(data);
					SE.event.fire('data.myaccount', data);
				}).fail(reject);
			});
		};
		return SE.req.cache.myaccount;
	},
	getmycards: function(subscriber) {
		if (subscriber) {
			SE.event.on('data.mycards', subscriber);
			return SE.req.getmycards();
		}

		if (!SE.req.cache.mycards) {
			console.log('request /api/mycards.json');
			SE.req.cache.mycards = new Promise(function(resolve, reject) {
				$.getJSON('/api/mycards.json').done(function(data) {
					resolve(data);
					SE.event.fire('data.mycards', data);
				}).fail(reject);
			});
		};
		return SE.req.cache.mycards;
	},
	getcards: function(subscriber) {
		if (subscriber) {
			SE.event.on('data.cards', subscriber);
			return SE.req.getcards();
		}

		if (!SE.req.cache.cards) {
			console.log('request /api/cards.json');
			SE.req.cache.cards = new Promise(function(resolve, reject) {
				$.getJSON('/api/cards.json').done(function(data) {
					resolve(data);
					SE.event.fire('data.cards', data);
				}).fail(reject);
			});
		};
		return SE.req.cache.cards;
	},
	getmydecks: function(subscriber) {
		if (subscriber) {
			SE.event.on('data.mydecks', subscriber);
			return SE.req.getmydecks();
		}

		if (!SE.req.cache.mydecks) {
			console.log('request /api/decks.json');
			SE.req.cache.mydecks = new Promise(function(resolve, reject) {
				$.getJSON('/api/mydecks.json').done(function(data) {
					resolve(data);
					SE.event.fire('data.mydecks', data);
				}).fail(reject);
			});
		};
		return SE.req.cache.mydecks;
	},
	openpack: function() {
		return new Promise(function(resolve, reject) {
			$.getJSON('/api/openpack.json').done(resolve).fail(reject);
		});
	}
};
