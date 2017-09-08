Promise.Fake = function() {
	var resolve, reject;
	var promise = new Promise(function(rs, rj) {
		resolve = rs;
		reject = rj;
	});

	promise.resolve = resolve;
	promise.reject = reject;
	promise.catch(() => {}).then(() => {promise.done = true;});
	return promise;
};
