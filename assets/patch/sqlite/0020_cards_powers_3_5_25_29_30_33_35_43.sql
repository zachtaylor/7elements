ALTER TABLE cards_powers RENAME TO _cards_powers;
CREATE TABLE cards_powers (
	cardid INTEGER,
	id INTEGER,
	trigger TEXT,
	target TEXT,
	usesturn INTEGER,
	script TEXT
);
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
  SELECT cardid, id, "instant", target, usesturn, script
  FROM _cards_powers;
DROP TABLE _cards_powers;
UPDATE cards_powers SET trigger="play" WHERE id=0;
UPDATE cards_powers SET id=0;
UPDATE cards_powers_costs SET powerid=0;
UPDATE cards_powers_texts SET powerid=0;
UPDATE cards_powers SET script="boen", target="self" WHERE cardid=3;
DELETE FROM cards_powers_costs WHERE cardid=3;
UPDATE cards_powers_texts SET description=
	"<img class='se-symbol' src='/img/icon/timer.20px.png'>:'Boen' gets +1 <img class='se-symbol' src='/img/icon/life.20px.png'>, and then draw a card"
	WHERE cardid=3;
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(5, 0, "instant", "bodyorplayer", 1, "ifrit");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (5, 0, "en-US",
	"<img class='se-symbol' src='/img/icon/timer.20px.png'>: Deal 1 damage to target <b>Body</b> or <b>Player</b>");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(11, 0, "play", "bodyoritem", 0, "energize");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (11, 0, "en-US", "Wake target <b>Body</b> or <b>Item</b>");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(25, 0, "play", "item", 0, "hard-bargain");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (25, 0, "en-US", "Destroy target <b>Item</b>");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(29, 0, "instant", "body", 1, "handrails");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (29, 0, "en-US",
	"<img class='se-symbol' src='/img/icon/timer.20px.png'>: Wake target <b>Body</b>");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(30, 0, "instant", "body", 1, "wand-of-suppression");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (30, 0, "en-US", "<img class='se-symbol' src='/img/icon/timer.20px.png'>: Sleep target <b>Body</b>");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(33, 0, "play", "body", 0, "lightning-strike");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (33, 0, "en-US", "Deal 3 damage to target <b>Body</b>");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(35, 0, "play", "", 0, "grave-birth");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (35, 0, "en-US", "Choose target <b>Body</b> in any player's <b>Past</b>, and create a new copy of it");
INSERT INTO cards_powers (cardid, id, trigger, target, usesturn, script)
	VALUES(43, 0, "instant", "", 1, "summoners-portal");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (43, 0, "en-US",
	"<img class='se-symbol' src='/img/icon/timer.20px.png'>: <span style='font-size:12px;'>Reveal the top card of your <b>Future</b><br/>If it is a <b>Body</b> or <b>Item</b>, then you may play it without paying any <b>Costs</b><br/>Otherwise, put it in your <b>Past</b>");
UPDATE cards_powers SET target="body", script="cloning-pool" WHERE cardid=48;
