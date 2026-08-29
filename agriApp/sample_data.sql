INSERT INTO products (name, category, price, stock) VALUES
('Wheat Seeds', 'Seeds', 1200.00, 50),
('Urea Fertilizer', 'Fertilizer', 850.00, 100),
('Pesticide X', 'Pesticide', 650.00, 30)
ON CONFLICT (id) DO NOTHING;

INSERT INTO sales (product_name, quantity, buyer, date, total) VALUES
('Wheat Seeds', 2, 'Ramesh', '2026-08-30', 2400.00),
('Urea Fertilizer', 1, 'Suresh', '2026-08-30', 850.00)
ON CONFLICT (id) DO NOTHING;

INSERT INTO schedules (farmer_name, crop_name, land_area, task, date, status) VALUES
('Ramesh Yadav', 'Wheat', '2 Acre', 'Seeding', '2026-09-02', 'Pending'),
('Suresh Pal', 'Rice', '3 Acre', 'Fertilizer Spray', '2026-09-05', 'Completed')
ON CONFLICT (id) DO NOTHING;
