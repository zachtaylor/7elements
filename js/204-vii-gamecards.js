vii.gamecard ={
	cache: {},
	get: function(id) {
		if (!vii.gamecard.cache[id]) {
			vii.gamecard.cache[id] = Promise.Fake();
		}
		return vii.gamecard.cache[id];
	},
	new: function(data) {
		return new Promise(function(resolve, reject) {
			SE.widget.new('se-gc', data).then(function(card) {
				resolve(card);
			}, reject);
		});
	},
	set: function(data) {
		if (!data.gcid) {
			console.warn('vii.gamecard.set data missing gcid', data);
		}
		var p = vii.gamecard.get(data.gcid);
		return new Promise(function(resolve, reject) {
			if (p.val) {
				resolve(p.val);
			}
			else {
				p.then(resolve, reject);
				vii.gamecard.new(data).then(p.resolve, p.reject);
			}
		});
	}
};
