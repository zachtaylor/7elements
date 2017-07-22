SE.widget = {
	cache: {},
	controllers: {},
	_promise: function(name) {
		return new Promise(function(resolve, reject) {
			$.get('/html/'+name+'.html').done(resolve).fail(reject);
		});
	},
	_builder: function(name, data) {
		return function WidgetFactory() {
			var widget = $(data)[0];
			$.each(SE.widget.controllers[name], function(i, f) {
				f.apply(widget);
			});
			return widget;
		};
	},
	get: function(name) {
		if (!SE.widget.cache[name]) SE.widget.cache[name] = SE.widget._promise(name);
		return new Promise(function(resolve, reject) {
			SE.widget.cache[name].then(function(data) {
				resolve(SE.widget._builder(name, data));
			}, reject);
		});
	},
	new: function(name) {
		return new Promise(function(resolve, reject) {
			SE.widget.get(name).then(function(widgetFactory) {
				resolve(widgetFactory());
			}, reject);
		});
	},
	provide: function(name, data) {
		SE.widget.cache[name] = Promise.resolve(data);
	},
	control: function(name, f) {
		SE.widget.controllers[name] = SE.widget.controllers[name] || [];
		SE.widget.controllers[name].push(f);
	},
	replace: function(name) {
		var script = document.currentScript;
		SE.widget.new(name).then(function(widget) {
			$(script).replaceWith(widget);
		}, function() {
			$(script).replaceWith('<span>widget not found: '+name+'</span>');
		});
	}
};
