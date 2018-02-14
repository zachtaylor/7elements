CREATE TABLE cards_powers (
	cardid INTEGER,
	id INTEGER,
	usesturn INTEGER,
	script TEXT
);
CREATE TABLE cards_powers_costs (
	cardid INTEGER,
	powerid INTEGER,
	element INTEGER,
	count INTEGER
);
CREATE TABLE cards_powers_texts (
	cardid INTEGER,
	powerid INTEGER,
	language TEXT,
	description TEXT
);

UPDATE cards_text SET name="Boen" where cardid=3;
INSERT INTO cards_powers (cardid, id, usesturn, script)
	VALUES (3, 1, 1, "boen-0");
INSERT INTO cards_powers_costs (cardid, powerid, element, count)
	VALUES (3, 1, 3, 1);
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (3, 1, "en-US", "<b><img class='se-symbol' src='/img/icon/element-3.png'>+<img class='se-symbol' src='/img/icon/timer.png'>:</b>'boen' gets +1 <img src='/img/icon/attack.20px.png'> and +1 <img src='/img/icon/life.20px.png'>");

INSERT INTO cards_powers (cardid, id, usesturn, script)
	VALUES (48, 1, 1, "cloning-pool-0");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (48, 1, "en-US", "<b><img class='se-symbol' src='/img/icon/timer.png'>:</b> Create a clone of target body, then you gain 1 life");

INSERT INTO cards_powers (cardid, id, usesturn, script)
	VALUES(46, 0, 0, "call-the-banners");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (46, 0, "en-US", "Create 3 lives<br/>Each life created has 2 <img src='/img/icon/attack.20px.png'> and 2 <img src='/img/icon/life.20px.png'>");

INSERT INTO cards_powers (cardid, id, usesturn, script)
	VALUES(50, 0, 0, "7-elements");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (50, 0, "en-US", "Win the game right now");
