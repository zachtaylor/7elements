SE.shadow = {
	cache: {},
	attach: function(el, name) {
		if (!SE.shadow.cache[name]) {
			SE.shadow.cache[name] = new Promise(function(resolve, reject) {
				$.get('/html/'+name+'.html').done(function(data) {
					var shadow = el.attachShadow({mode:'open'});
					shadow.innerHTML = data;
					resolve(shadow);
				}).fail(reject);
			});
		};
		return SE.shadow.cache[name];
	},
	get: function(name) {
		return SE.shadow.cache[name];
	}
};
