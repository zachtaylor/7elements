INSERT INTO cards (id, type, image)
	VALUES (43, 3, "/img/card/43.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (43, 1, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (43, "en-US", "Summoner's Portal", "<se-symbol-tap></se-symbol-tap>: Show the top card of your deck, if it is a spell, put it in the spent pile, otherwise put it on your side spent", "");

INSERT INTO cards (id, type, image)
	VALUES (44, 1, "/img/card/44.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (44, 2, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (44, "en-US", "De ja Vu", "Skip your opponent(s) next turn(s)", "");

INSERT INTO cards (id, type, image)
	VALUES (45, 1, "/img/card/45.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (45, 3, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (45, "en-US", "Overgrowth", "Creatures you control gain +'x'/+0 where 'x' is the number of bodys you control", "");

INSERT INTO cards (id, type, image)
	VALUES (46, 1, "/img/card/46.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (46, 4, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (46, "en-US", "Call the Banners", "Create 3 bodys (3/3)", "");

INSERT INTO cards (id, type, image)
	VALUES (47, 1, "/img/card/47.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (47, 5, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (47, "en-US", "Death from Above", "Destory all bodies, then create <se-symbol icon='element-5'></se-symbol>", "");

INSERT INTO cards (id, type, image)
	VALUES (48, 3, "/img/card/48.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (48, 6, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (48, "en-US", "Cloning Pool", "<b>spend</b>: Create a spent copy of target item or body", "");

INSERT INTO cards (id, type, image)
	VALUES (49, 3, "/img/card/49.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (49, 7, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (49, "en-US", "Necromaster's Tome", "<b>spend</b>: Move a body card from a spent pile to your side, spent, then you lose 2 life", "");
