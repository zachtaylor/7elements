-- update deck tables

ALTER TABLE decks RENAME COLUMN cover TO cover_old;
ALTER TABLE decks DROP COLUMN level;
ALTER TABLE decks ADD user VARCHAR(24);
ALTER TABLE decks ADD cover INT;
UPDATE TABLE decks SET user='vii';

-- patch #5 is an executable
