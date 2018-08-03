vii.sound = {
	cache: {},
	play: function(name) {
		SE.go(function() {
			vii.sound._buildAudio(name).play();
		});
	},
	_buildAudio: function(name) {
		// https://stackoverflow.com/questions/9847580/how-to-detect-safari-chrome-ie-firefox-and-opera-browser
		var isFirefox = typeof InstallTrigger !== 'undefined';
		if (isFirefox) return new Audio('/sound/'+name+'.ogg');
		else return new Audio('/sound/'+name+'.mp3');
	}
};
