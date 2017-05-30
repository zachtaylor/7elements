SE.myaccount = {
	get: function() {
		return new Promise(function(resolve, reject) {
			$.getJSON('/api/myaccount.json').done(resolve).fail(reject);
		})
	}
};
