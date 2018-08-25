INSERT INTO cards (id, type, name, text, image)
	VALUES (1, 2, "Time Traveler", "", "/img/card/1.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (1, 1, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(1, 0, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (1, 1, "", "", 0, 1, "time-traveler", "<b>kill</b>: Create a new Element");

INSERT INTO cards (id, type, name, text, image)
	VALUES (2, 2, "Ifrit", "", "/img/card/2.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (2, 5, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(2, 1, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (2, 1, "", "bodyorplayer", 1, 0, "ifrit", "<b>use</b>: Deal 1 damage to Target Being or Player");

INSERT INTO cards (id, type, name, text, image)
	VALUES (3, 2, "Vine Spirit", "", "/img/card/3.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (3, 3, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(3, 1, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (3, 1, "", "", 1, 0, "vine-spirit", "<b>use</b>: Draw a card");

INSERT INTO cards (id, type, name, text, image)
	VALUES (4, 2, "Zealot", "So where is the party?", "/img/card/4.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (4, 4, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(4, 2, 2);

INSERT INTO cards (id, type, name, text, image)
	VALUES (5, 2, "Water Dancer", "", "/img/card/5.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (5, 2, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(5, 0, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (5, 1, "", "body", 1, 0, "water-dancer", "<b>use</b>: Target Being becomes Asleep");

INSERT INTO cards (id, type, name, text, image)
	VALUES (6, 2, "Pixie", "", "/img/card/6.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (6, 6, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(6, 0, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (6, 1, "", "bodyorplayer", 1, 0, "pixie", "<b>use</b>: Target Being or Player gets +1 life");

INSERT INTO cards (id, type, name, text, image)
	VALUES (7, 2, "Nightmare Ader", "", "/img/card/7.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (7, 7, 1);
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(7, 0, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (7, 1, "", "", 0, 1, "nightmare-ader", "<b>kill</b>: Create a Body which is a copy of a card in any players' Past");

INSERT INTO cards (id, type, name, text, image)
	VALUES (8, 1, "New Element", "Prepare for tomorrow", "/img/card/8.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (8, 1, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (8, 1, "play", "", 0, 0, "new-element", "Create a new Element");

INSERT INTO cards (id, type, name, text, image)
	VALUES (9, 1, "Burn", "Burn...", "/img/card/9.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (9, 5, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (9, 1, "play", "body", 0, 0, "burn", "Deal 2 damage to Target Being");

INSERT INTO cards (id, type, name, text, image)
	VALUES (10, 1, "Inspire Growth", "The best defense is a strong offense", "/img/card/10.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (10, 3, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (10, 1, "play", "", 0, 0, "inspire-growth", "Target Being gains +2 Attack");

INSERT INTO cards (id, type, name, text, image)
	VALUES (11, 1, "Energize", "Now is the time!", "img/card/11.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (11, 4, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (11, 1, "play", "", 0, 0, "energize", 'Add <app-icon src="/img/icon/element-4.png"></app-icon><app-icon src="/img/icon/element-4.png"></app-icon> to your Karma until Sunset');

INSERT INTO cards (id, type, name, text, image)
	VALUES (12, 1, "Bend Will", "It's a trap!", "/img/card/12.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (12, 2, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (12, 1, "play", "bodyoritem", 0, 0, "bend-will", "Choose one<br/>-Target Being or Item becomes Asleep<br/>-Target Being or Item becomes Awake");

INSERT INTO cards (id, type, name, text, image)
	VALUES (13, 1, "Grace", "Is that better?", "/img/card/13.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (13, 6, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (13, 1, "play", "body", 0, 0, "grace", "Target Being gets +2 Life");

INSERT INTO cards (id, type, name, text, image)
	VALUES (14, 1, "Memorialize", "Never forget", "/img/card/14.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (14, 7, 1);
INSERT INTO cards_powers (cardid, id, xtrigger,	target,	usesturn, useskill,	script, text)
	VALUES (14, 1, "play", "", 0, 0, "memorialize", "Create a Body which is a copy of a card in any players' Past");

INSERT INTO packs (id, name, size, cost, image)
	VALUES (1, "Alpha Beings Pack (1)", 1, 3, "/img/pack/1.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
	VALUES (1, 1, 1), (1, 2, 1), (1, 3, 1), (1, 4, 1), (1, 5, 1), (1, 6, 1), (1, 7, 1);
INSERT INTO packs (id, name, size, cost, image)
	VALUES (2, "Alpha Beings Pack (3)", 3, 5, "/img/pack/2.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
	VALUES (2, 1, 1), (2, 2, 1), (2, 3, 1), (2, 4, 1), (2, 5, 1), (2, 6, 1), (2, 7, 1);
INSERT INTO packs (id, name, size, cost, image)
	VALUES (3, "Alpha Beings Pack (5)", 5, 7, "/img/pack/3.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
	VALUES (3, 1, 1), (3, 2, 1), (3, 3, 1), (3, 4, 1), (3, 5, 1), (3, 6, 1), (3, 7, 1);
INSERT INTO packs (id, name, size, cost, image)
	VALUES (4, "Alpha Spells Pack (1)", 1, 3, "/img/pack/4.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
	VALUES (4, 8, 1), (4, 9, 1), (4, 10, 1), (4, 11, 1), (4, 12, 1), (4, 13, 1), (4, 14, 1);
INSERT INTO packs (id, name, size, cost, image)
	VALUES (5, "Alpha Spells Pack (3)", 3, 5, "/img/pack/5.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
	VALUES (5, 8, 1), (5, 9, 1), (5, 10, 1), (5, 11, 1), (5, 12, 1), (5, 13, 1), (5, 14, 1);
INSERT INTO packs (id, name, size, cost, image)
	VALUES (6, "Alpha Spells Pack (5)", 5, 7, "/img/pack/6.jpg");
INSERT INTO packs_cards (packid, cardid, weight)
	VALUES (6, 8, 1), (6, 9, 1), (6, 10, 1), (6, 11, 1), (6, 12, 1), (6, 13, 1), (6, 14, 1);

INSERT INTO decks (id, name, level, color)
	VALUES (1, "AllCards", 1, "black");
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (1, 1, 3), (1, 2, 3), (1, 3, 3), (1, 4, 3), (1, 5, 3), (1, 6, 3), (1, 7, 3), (1, 8, 3), (1, 9, 3), (1, 10, 3), (1, 11, 3), (1, 12, 3), (1, 13, 3), (1, 14, 3);

INSERT INTO decks (id, name, level, color)
	VALUES (2, "WREG", 2, "red");
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (2, 1, 3), (2, 2, 3), (2, 3, 3), (2, 4, 3), (2, 8, 3), (2, 9, 3), (2, 10, 3);

INSERT INTO decks (id, name, level, color)
	VALUES (3, "GBVK", 3, "fuchsia");
INSERT INTO decks_items (deckid, cardid, amount)
	VALUES (3, 4, 3), (3, 5, 3), (3, 6, 3), (3, 7, 3), (3, 12, 3), (3, 13, 3), (3, 14, 3);
