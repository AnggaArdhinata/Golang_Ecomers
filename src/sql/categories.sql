--DROP TABLE CATEGORIES
-- DROP TABLE IF EXISTS categories;

--INSERT
-- INSERT INTO categories ("name", created_at, updated_at) VALUES('electronic', 'now()', 'now()');

--SELECT
-- SELECT * FROM categories

--UPDATE CATEGORY
UPDATE categories SET name='food', updated_at='now()' WHERE id= 1