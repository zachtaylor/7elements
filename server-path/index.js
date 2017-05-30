$(function() {
	SE.myaccount.get().then(function(data) {
		console.log(data);
		$('#myaccount').html(data.username + ' (' + data.cards + ' cards)');
	}, function(error) {
		if (error.responseText == 'session missing') {
			$('#myaccount').html('not logged in')[0].href = "/login/";
			return;
		};

		console.error(error.responseText);
	});
});
