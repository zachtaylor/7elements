SE.req = {
	cache: {},
	getmyaccount: function() {
		if (!SE.req.cache.myaccount) {
			SE.req.cache.myaccount = new Promise(function(resolve, reject) {
				$.getJSON('/api/myaccount.json').done(function(data) {
					resolve(data);
					SE.event.fire('data.myaccount', data);
				}).fail(reject);
			});
		};
		return SE.req.cache.myaccount;
	},
	getcards: function() {
		if (!SE.req.cache.cards) {
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
