-- Insert one author
INSERT INTO authors (id, name) VALUES ('123e4567-e89b-12d3-a456-426655440001', 'John Smith');

-- Insert 100 books by the author with random titles
INSERT INTO books (id, title, author_id, content, created_at, updated_at)
SELECT uuid_generate_v4(), 'Book ' || i, '123e4567-e89b-12d3-a456-426655440001', 'Content for book ' || i, NOW(), NOW()
FROM generate_series(1, 100) i;