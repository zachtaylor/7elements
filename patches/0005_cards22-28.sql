INSERT INTO cards (id, type, image)
	VALUES (22, 1, "/img/card/22.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (22, 1, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (22, "en-US", "Cast Out", "Hide target item or body", "Cast Out does not 'kill' bodies or affect the spent pile");

INSERT INTO cards (id, type, image)
	VALUES (23, 1, "/img/card/23.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (23, 2, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (23, "en-US", "Cancel", "Return target play, body, or item to it's owner's hand", "BOOOOOOOOOOOOOOOOMMMMMM");

INSERT INTO cards (id, type, image)
	VALUES (24, 1, "/img/card/24.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (24, 3, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (24, "en-US", "Savage Strike", "Spend an unspent body you control, then deal 'x' damage to target body, where 'x' is 2 times your body's attack", "");

INSERT INTO cards (id, type, image)
	VALUES (25, 1, "/img/card/25.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (25, 4, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (25, "en-US", "Hard Bargain", "Destroy target nonelement nonbody item", "You got the stuff, and I got the money...");

INSERT INTO cards (id, type, image)
	VALUES (26, 1, "/img/card/26.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (26, 5, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (26, "en-US", "Channel Lightning", "Deal 3 damage divided how you choose among target body(s) and/or target player(s)", "");

INSERT INTO cards (id, type, image)
	VALUES (27, 1, "/img/card/27.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (27, 6, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (27, "en-US", "Divine Shield", "Target body gains unkillable until end of turn, then you gain 2 life", "");

INSERT INTO cards (id, type, image)
	VALUES (28, 1, "/img/card/28.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (28, 7, 2);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (28, "en-US", "Painful Memories", "Return up to 2 target body cards from your graveyard to your hand, then you lose 2 life", "");
