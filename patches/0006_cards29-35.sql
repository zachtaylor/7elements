INSERT INTO cards (id, type, image)
	VALUES (29, 3, "/img/card/29.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (29, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (29, 1, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (29, "en-US", "Handrails", "<b>spend</b>: Unspend target item or body", "");

INSERT INTO cards (id, type, image)
	VALUES (30, 3, "/img/card/30.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (30, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (30, 2, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (30, "en-US", "Wand of Suppression", "<b>spend</b>: Spend target item or body", "");

INSERT INTO cards (id, type, image)
	VALUES (31, 2, "/img/card/31.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (31, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (31, 3, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (31, "en-US", "Sunfeeder Elf", "<b>spend</b>:Sunfeeder Elf gets +1 attack", "None can grow like those that feed on the sun");
INSERT INTO cards_bodies (cardid, attack, health)
	VALUES(31, 2, 3);

INSERT INTO cards (id, type, image)
	VALUES (32, 2, "/img/card/32.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (32, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (32, 4, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (32, "en-US", "Instigator", "<se-symbol icon='element-4'></se-symbol>:Instigator gets +1 attack and +1 health", "");
INSERT INTO cards_bodies(cardid, attack, health)
	VALUES(32, 2, 3);

INSERT INTO cards (id, type, image)
	VALUES (33, 1, "/img/card/33.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (33, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (33, 5, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (33, "en-US", "Lightning Strike", "Deal 3 damage to target body", "");

INSERT INTO cards (id, type, image)
	VALUES (34, 2, "/img/card/34.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (34, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (34, 6, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (34, "en-US", "Sprite", "<se-symbol icon='element-0'></se-symbol> + <se-symbol icon='element-6'></se-symbol> + <b>spend</b>: Create another Sprite", "");
INSERT INTO cards_bodies(cardid, attack, health)
	VALUES(34, 2, 3);

INSERT INTO cards (id, type, image)
	VALUES (35, 1, "/img/card/35.jpg");
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (35, 0, 1);
INSERT INTO cards_element_costs (cardid, element, count)
	VALUES (35, 7, 1);
INSERT INTO cards_text (cardid, language, name, description, flavor)
	VALUES (35, "en-US", "Grave Birth", "Move target body card from your spent pile back into play, spent, then you lose 2 life", "Thriller!");
