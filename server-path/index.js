$(function() {
	SE.myaccount.get().then(function(data) {
		if (data.packs > 0) {
			$('#data-myaccount-packs').html(data.packs);
		};
		$('#data-myaccount-username').html(' ('+data.username+') ');
		$('#data-myaccount-cards').html(data.cards);
	}, function(error) {
		if (error.responseText == 'session missing') {
			$('#nav-myaccount-link')[0].href = '/login/';
			$('#data-myaccount-username').html(' (login) ');
			return;
		};

		console.error(error.responseText);
	});
});
