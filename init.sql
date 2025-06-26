-- Initialize the library database with sample data

-- Create extensions if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Insert sample books (will be created after GORM migrations)
-- This file ensures the database is properly initialized

-- Sample data insertion (optional - can be removed if not needed)
INSERT INTO books (title, author, isbn, published_year, genre, available_copies, created_at, updated_at) 
VALUES 
    ('The Go Programming Language', 'Alan Donovan & Brian Kernighan', '978-0134190440', 2015, 'Programming', 5, NOW(), NOW()),
    ('Clean Code', 'Robert C. Martin', '978-0132350884', 2008, 'Programming', 3, NOW(), NOW()),
    ('Design Patterns', 'Gang of Four', '978-0201633612', 1994, 'Programming', 2, NOW(), NOW()),
    ('The Pragmatic Programmer', 'David Thomas & Andrew Hunt', '978-0135957059', 2019, 'Programming', 4, NOW(), NOW()),
    ('Effective Go', 'Various Authors', '978-1234567890', 2020, 'Programming', 6, NOW(), NOW())
ON CONFLICT (isbn) DO NOTHING;