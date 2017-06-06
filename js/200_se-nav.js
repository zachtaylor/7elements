customElements.define('se-nav', class extends HTMLElement {
	constructor() {
		super();
		SE.event.on('data.myaccount', this.updateMyAccount);
	}

	connectedCallback() {
		SE.shadow.attach(this, 'se-nav').then(function(shadow) {
			SE.req.getmyaccount().then(function() {}, function() {
				var myAccountLink = shadow.querySelector('#myaccount-link');
				myAccountLink.innerHTML = 'Login';
				myAccountLink.href = '/login/';
			});

			var pathname = window.location.pathname;
			if (pathname == '/cards/') {
				shadow.querySelector('#cards-link').classList.add('active');
			} else if (pathname == '/packs/') {
				shadow.querySelector('#packs-link').classList.add('active');
			} else if (pathname == '/login/') {
				shadow.querySelector('#myaccount-link').classList.add('active');
			} else if (pathname == '/myaccount/') {
				shadow.querySelector('#myaccount-link').classList.add('active');
			}
		});
	}

	updateMyAccount(data) {
		SE.shadow.get('se-nav').then(function(shadow) {
			shadow.querySelector('#data-cards').innerHTML = data.cards;
			shadow.querySelector('#data-packs').innerHTML = data.packs;
			shadow.querySelector('#data-username').innerHTML = '('+data.username+')';
		});
	}
});
