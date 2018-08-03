UPDATE cards SET image="/img/cards/time-walker-0.jpg" WHERE id=1;
UPDATE cards_text SET name="Time Walker" WHERE name="Time Dealer";
INSERT INTO cards_powers (cardid, id, usesturn, target, script)
	VALUES(1, 1, 1, "", "time-walker");
INSERT INTO cards_powers_texts (cardid, powerid, language, description)
	VALUES (1, 1, "en-US", "<img class='se-symbol' src='/img/icon/timer.20px.png'>: Create a new element");
