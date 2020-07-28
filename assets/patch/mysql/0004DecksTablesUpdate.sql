-- update deck tables

ALTER TABLE accounts_decks DROP COLUMN register;
ALTER TABLE accounts_decks DROP COLUMN wins;
ALTER TABLE decks RENAME COLUMN cover TO cover_old;
ALTER TABLE decks DROP COLUMN level;
ALTER TABLE decks ADD cover INT;
