-- Insert data into users table
INSERT INTO users (username, email, password_hash, full_name) VALUES
('user1', 'user1@example.com', 'hashedpassword1', 'User One'),
('user2', 'user2@example.com', 'hashedpassword2', 'User Two'),
('user3', 'user3@example.com', 'hashedpassword3', 'User Three');

-- Insert data into categories table
INSERT INTO categories (category_name) VALUES
('Desserts'),
('Main Courses'),
('Appetizers');

-- Insert recipes without featured_image_id
INSERT INTO recipes (user_id, category_id, title, description, preparation_time) VALUES
((SELECT id FROM users WHERE username = 'user1'), (SELECT id FROM categories WHERE category_name = 'Desserts'), 'Chocolate Cake', 'A rich chocolate cake', '1 hour'),
((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM categories WHERE category_name = 'Main Courses'), 'Spaghetti Carbonara', 'Classic Italian pasta dish', '30 minutes'),
((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM categories WHERE category_name = 'Appetizers'), 'Bruschetta', 'Grilled bread with tomato and basil', '15 minutes');

-- Insert recipe images
INSERT INTO recipe_images (recipe_id, image_url, is_featured) VALUES
((SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 'https://example.com/images/chocolate_cake.jpg', TRUE),
((SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 'https://example.com/images/spaghetti_carbonara.jpg', TRUE),
((SELECT id FROM recipes WHERE title = 'Bruschetta'), 'https://example.com/images/bruschetta.jpg', TRUE);

-- Update recipes with featured_image_id
UPDATE recipes SET featured_image_id = (SELECT id FROM recipe_images WHERE recipe_id = recipes.id AND is_featured = TRUE);

-- Insert data into steps table
INSERT INTO steps (recipe_id, step_number, instruction) VALUES
((SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 1, 'Preheat the oven to 350°F (175°C).'),
((SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 2, 'Mix flour, sugar, cocoa powder, baking powder, and salt.'),
((SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 1, 'Boil pasta in salted water until al dente.'),
((SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 2, 'Cook pancetta in a pan until crispy.'),
((SELECT id FROM recipes WHERE title = 'Bruschetta'), 1, 'Grill the bread slices until golden brown.'),
((SELECT id FROM recipes WHERE title = 'Bruschetta'), 2, 'Rub garlic on the grilled bread.');

-- Insert data into ingredients table
INSERT INTO ingredients (recipe_id, ingredient_name, quantity) VALUES
((SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 'Flour', '2 cups'),
((SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 'Cocoa powder', '1/2 cup'),
((SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 'Spaghetti', '200g'),
((SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 'Pancetta', '100g'),
((SELECT id FROM recipes WHERE title = 'Bruschetta'), 'Tomatoes', '3 medium'),
((SELECT id FROM recipes WHERE title = 'Bruschetta'), 'Basil leaves', 'A handful');

-- Insert data into likes table
INSERT INTO likes (user_id, recipe_id, reaction) VALUES
((SELECT id FROM users WHERE username = 'user1'), (SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), TRUE),
((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM recipes WHERE title = 'Bruschetta'), TRUE),
((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM recipes WHERE title = 'Chocolate Cake'), FALSE);

-- Insert data into bookmarks table
INSERT INTO bookmarks (user_id, recipe_id) VALUES
((SELECT id FROM users WHERE username = 'user1'), (SELECT id FROM recipes WHERE title = 'Bruschetta')),
((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM recipes WHERE title = 'Chocolate Cake')),
((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'));

-- Insert data into comments table
INSERT INTO comments (user_id, recipe_id, content) VALUES
((SELECT id FROM users WHERE username = 'user1'), (SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 'Delicious recipe!'),
((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM recipes WHERE title = 'Bruschetta'), 'Perfect for a quick snack!'),
((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 'Rich and decadent.');

-- Insert data into ratings table
INSERT INTO ratings (user_id, recipe_id, rating) VALUES
((SELECT id FROM users WHERE username = 'user1'), (SELECT id FROM recipes WHERE title = 'Chocolate Cake'), 5),
((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM recipes WHERE title = 'Spaghetti Carbonara'), 4),
((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM recipes WHERE title = 'Bruschetta'), 3);
