INSERT INTO cards (id, type, image)
	VALUES (36, 3, "/img/card/36.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (36, 1, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (36, "en-US", "Banhammer", "Spend an unspent body you control, <se-symbol-tap></se-symbol-tap>: Hide target item", "Banhammer does not 'kill' bodies or affect the spent pile");

INSERT INTO cards (id, type, image)
	VALUES (37, 1, "/img/card/37.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (37, 2, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (37, "en-US", "Strangle Hold", "Spend up to 2 target items or bodies, and on your opponent(s) next unspend step(s), they do not unspend any items or bodies", "");

INSERT INTO cards (id, type, image)
	VALUES (38, 1, "/img/card/38.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (38, 3, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (38, "en-US", "Slaughter", "Until end of turn, bodys you control gain '<se-symbol-tap></se-symbol-tap>: Deal damage equal to this body's attack to target body or player'", "");

INSERT INTO cards (id, type, image)
	VALUES (39, 1, "/img/card/39.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (39, 4, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (39, "en-US", "Reset", "Destroy all items and bodies", "");

INSERT INTO cards (id, type, image)
	VALUES (40, 1, "/img/card/40.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (40, 5, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (40, "en-US", "All Shall Burn", "Deal 3 damage to all bodys your opponent(s) control", "");

INSERT INTO cards (id, type, image)
	VALUES (41, 1, "/img/card/41.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (41, 6, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (41, "en-US", "Saving Grace", "Creatures you control gain indestructible until end of turn", "");

INSERT INTO cards (id, type, image)
	VALUES (42, 1, "/img/card/42.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (42, 7, 3);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (42, "en-US", "RIP", "You gain life equal to target body's health, then destroy that body", "");
