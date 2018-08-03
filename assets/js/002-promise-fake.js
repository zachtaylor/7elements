Promise.Fake = function() {
	var resolve, reject;
	var promise = new Promise(function(rs, rj) {
		resolve = rs;
		reject = rj;
	});

	promise.resolve = function(v) {
		promise.val = v;
		resolve(v);
	};
	promise.reject = reject;
	// promise.catch(() => {}).then(() => {promise.done = true;});
	return promise;
};
