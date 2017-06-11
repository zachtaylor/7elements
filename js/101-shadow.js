SE.shadow = {
	cache: {},
	attach: function(el, name) {
		if (!SE.shadow.cache[name]) {
			SE.shadow.cache[name] = new Promise(function(resolve, reject) {
				$.get('/html/'+name+'.html').done(function(data) {
					resolve(data);
				}).fail(reject);
			});
		}

		return new Promise(function(resolve, reject) {
			SE.shadow.cache[name].then(function(data) {
				var shadow = el.shadowRoot;
				if (!shadow) shadow = el.attachShadow({mode:'open'});
				shadow.innerHTML = data;
				resolve(shadow);
			}, reject);
		});
	}
};
