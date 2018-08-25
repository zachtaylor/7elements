INSERT INTO cards (id, type, image)
	VALUES (15, 3, "/img/card/15.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (15, 1, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (15, "en-US", "The TARBIS", "<b>spend</b>: Hide an item or body inside The TARBIS<br/><se-symbol icon='element-1'></se-symbol>: Move all cards hidden in The TARBIS back into play", "The TARBIS does not 'kill' bodies or affect the spent pile");

INSERT INTO cards (id, type, image)
	VALUES (16, 3, "/img/card/16.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (16, 2, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (16, "en-US", "Pensieve", "<b>spend</b>: Peek at your next 2 cards, you may put either or both into your spent pile<br/><b>When you play a spell</b>: Untap Pensieve", "");

INSERT INTO cards (id, type, image)
	VALUES (17, 3, "/img/card/17.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (17, 3, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (17, "en-US", "Personal Trainer", "<b>spend</b>: Target body gets +2 attack<br/><b>When you play a body</b>: Unspend Personal Trainer", "");

INSERT INTO cards (id, type, image)
	VALUES (18, 2, "/img/card/18.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (18, 4, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (18, "en-US", "Bouncer", "", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(18, 3, 3);

INSERT INTO cards (id, type, image)
	VALUES (19, 3, "/img/card/19.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (19, 5, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (19, "en-US", "Burning Rage", "<b>spend</b>: Deal 2 damage to target body<br/><b>When you lose life</b>: Untap Burning Rage", "");

INSERT INTO cards (id, type, image)
	VALUES (20, 3, "/img/card/20.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (20, 6, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (20, "en-US", "Fountain of Life", "<b>spend</b>: Target body gets +2 life<br/><b>When you gain life</b>: Untap Foutain of Life", "");

INSERT INTO cards (id, type, image)
	VALUES (21, 3, "/img/card/21.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (21, 7, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (21, "en-US", "Swamp Preserve", "<se-symbol-tap></se-symbol-tap> + sacrifice <se-symbol icon='element-7'></se-symbol>: Draw a card<br/><b>:When a body of yours dies</b>: create <se-symbol icon='element-7'></se-symbol>", "");
