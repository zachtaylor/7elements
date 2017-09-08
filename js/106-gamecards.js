SE.gamecards = {
	cache: {},
	_promise: function(gcid, cardid) {
		return new Promise(function(resolve, reject) {
			SE.widget.new('se-card', cardid).then(function(card) {
				card.gcid = gcid;
				resolve(card);
			}, reject);
		});
	},
	get: function(gcid, cardid) {
		if (!SE.gamecards.cache[gcid]) {
			if (!cardid) {
				console.error('gcid not usable', gcid);
				var promise = Promise.Fake();
				promise.reject();
				return promise;
			}
			SE.gamecards.cache[gcid] = SE.gamecards._promise(gcid, cardid);
		}
		return SE.gamecards.cache[gcid];
	}
};
