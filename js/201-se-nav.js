customElements.define('se-nav', class extends HTMLElement {
	constructor() {
		super();
		SE.event.on('data.myaccount', this.updater())
	}

	connectedCallback() {
		var updater = this.updater();

		SE.shadow.attach(this, 'se-nav').then(function(shadow) {
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

			SE.req.getmyaccount().then(updater, function() {
				var myAccountLink = shadow.querySelector('#myaccount-link');
				myAccountLink.innerHTML = 'Login';
				myAccountLink.href = '/login/';
			});
		});
	}

	update(data) {
		if (!this.shadowRoot) return;
		this.shadowRoot.querySelector('#data-cards').innerHTML = data.cards;
		this.shadowRoot.querySelector('#data-packs').innerHTML = data.packs;
		this.shadowRoot.querySelector('#data-username').innerHTML = '('+data.username+')';
	}

	updater() {
		var nav = this;
		return function(data) {
			nav.update(data);
		};
	}
});
