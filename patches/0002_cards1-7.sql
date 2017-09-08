INSERT INTO cards (id, type, image)
	VALUES (1, 2, "/img/card/1.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (1, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (1, "en-US", "Time Dealer", "<b>spend</b>: Create a new element", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(1, 0, 1);

INSERT INTO cards (id, type, image)
	VALUES (2, 2, "/img/card/2.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (2, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (2, "en-US", "Novice Seer", "<b>spend</b>: Peek at your next card, you may put it in the spent pile<br/><b>sacrifice</b>: Draw a card", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(2, 1, 1);

INSERT INTO cards (id, type, image)
	VALUES (3, 2, "/img/card/3.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (3, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (3, "en-US", "Elven Child", "<b>spend</b>: Target body gets +1 attack", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(3, 1, 2);

INSERT INTO cards (id, type, image)
	VALUES (4, 2, "/img/card/4.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (4, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (4, "en-US", "Zealot", "", "So where is the party?");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(4, 2, 2);

INSERT INTO cards (id, type, image)
	VALUES (5, 2, "/img/card/5.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (5, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (5, "en-US", "Ifrit", "<b>When Ifrit dies</b>: deal damage equal to Ifrit's attack to target body or player", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(5, 2, 1);

INSERT INTO cards (id, type, image)
	VALUES (6, 2, "/img/card/6.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (6, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (6, "en-US", "Pixie", "<b>spend</b>: Target body gets +1 life<br/><b>sacrifice</b>: You gain 1 life", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(6, 1, 1);

INSERT INTO cards (id, type, image)
	VALUES (7, 2, "/img/card/7.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (7, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (7, "en-US", "Nightmare Ader", "<b>When Nightmare Ader dies</b>: destroy target body", "");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(7, 0, 1);
