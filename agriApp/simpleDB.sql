INSERT INTO farmers (name, village, mobile, land_area) VALUES
('Ramesh Yadav', 'Kanpur Dehat', '9876543210', '2 Acre'),
('Suresh Pal', 'Bilhaur', '9876543211', '3 Acre');

INSERT INTO products (name, category, price, stock) VALUES
('Wheat Seeds', 'Seeds', 1200.00, 50),
('Urea Fertilizer', 'Fertilizer', 850.00, 100),
('Pesticide X', 'Pesticide', 650.00, 30);

INSERT INTO sales (farmer_id, product_id, quantity, total, sale_date) VALUES
(1, 1, 2, 2400.00, '2026-08-30'),
(2, 2, 1, 850.00, '2026-08-30');

INSERT INTO farming_schedules (farmer_id, crop_name, season, task, schedule_date, status) VALUES
(1, 'Wheat', 'Rabi', 'Seed Sowing', '2026-09-02', 'Pending'),
(1, 'Wheat', 'Rabi', 'Irrigation', '2026-09-10', 'Pending'),
(2, 'Rice', 'Kharif', 'Fertilizer Spray', '2026-09-05', 'Completed');