SE.req = {
	cache: {},
	getmyaccount: function(subscriber) {
		if (subscriber) {
			SE.event.on('data.myaccount', subscriber);
			return SE.req.getmyaccount().then(subscriber);
		}

		if (!SE.req.cache.myaccount) {
			console.log('request \'/api/myaccount.json\'');
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
			return SE.req.getmycards().then(subscriber);
		}

		if (!SE.req.cache.mycards) {
			console.log('request \'/api/mycards.json\'');
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
			return SE.req.getcards().then(subscriber);
		}

		if (!SE.req.cache.cards) {
			console.log('request \'/api/cards.json\'');
			SE.req.cache.cards = new Promise(function(resolve, reject) {
				$.getJSON('/api/cards.json').done(function(data) {
					resolve(data);
					SE.event.fire('data.cards', data);
				}).fail(reject);
			});
		};
		return SE.req.cache.cards;
	},
	openpack: function() {
		return new Promise(function(resolve, reject) {
			$.getJSON('/api/openpack.json').done(resolve).fail(reject);
		});
	}
};
