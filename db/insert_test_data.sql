-- Вставка тестовых данных в таблицу suppliers
INSERT INTO suppliers (title, email, phone, created_at, is_deleted) VALUES
('ООО Манижа-СОГД', 'supplier1@example.com', '123-456-7890', CURRENT_TIMESTAMP, false),
('ООО Сафар', 'supplier2@example.com', '234-567-8901', CURRENT_TIMESTAMP, false),
('ООО ГлавПродукт', 'supplier3@example.com', '345-678-9012', CURRENT_TIMESTAMP, false),
('ООО Империя', 'supplier4@example.com', '456-789-0123', CURRENT_TIMESTAMP, false),
('ООО Продукт No.1', 'supplier5@example.com', '567-890-1234', CURRENT_TIMESTAMP, false),
('ООО Памир', 'supplier6@example.com', '678-901-2345', CURRENT_TIMESTAMP, false),
('ООО Памир-2009', 'supplier7@example.com', '789-012-3456', CURRENT_TIMESTAMP, false),
('ООО Сухроб', 'supplier8@example.com', '890-123-4567', CURRENT_TIMESTAMP, false),
('ООО Аниса', 'supplier9@example.com', '901-234-5678', CURRENT_TIMESTAMP, false),
('ООО Комрон-99', 'supplier10@example.com', '012-345-6789', CURRENT_TIMESTAMP, false);

-- Вставка тестовых данных в таблицу categories 
INSERT INTO categories (title, description, created_at, is_deleted) VALUES
('Фрукты', 'Категория для всех видов фруктов', CURRENT_TIMESTAMP, false),
('Овощи', 'Категория для всех видов овощей', CURRENT_TIMESTAMP, false),
('Молочные продукты', 'Категория для молочных продуктов', CURRENT_TIMESTAMP, false),
('Мясо', 'Категория для мясных продуктов', CURRENT_TIMESTAMP, false),
('Хлебобулочные изделия', 'Категория для хлебобулочных изделий', CURRENT_TIMESTAMP, false),
('Напитки', 'Категория для всех видов напитков', CURRENT_TIMESTAMP, false),
('Снеки', 'Категория для различных закусок', CURRENT_TIMESTAMP, false),
('Замороженные продукты', 'Категория для замороженных продуктов', CURRENT_TIMESTAMP, false),
('Консервы', 'Категория для консервированных продуктов', CURRENT_TIMESTAMP, false),
('Личная гигиена', 'Категория для товаров личной гигиены', CURRENT_TIMESTAMP, false),
('Кондитерские изделия', 'Категория для сладостей и кондитерских изделий', CURRENT_TIMESTAMP, false),
('Крупы и макароны', 'Категория для круп и макаронных изделий', CURRENT_TIMESTAMP, false),
('Масла и соусы', 'Категория для растительных масел и соусов', CURRENT_TIMESTAMP, false),
('Рыба и морепродукты', 'Категория для рыбы и морепродуктов', CURRENT_TIMESTAMP, false),
('Специи и приправы', 'Категория для специй и приправ', CURRENT_TIMESTAMP, false),
('Алкоголь', 'Категория для алкогольных напитков', CURRENT_TIMESTAMP, false),
('Табачные изделия', 'Категория для сигарет и табачных изделий', CURRENT_TIMESTAMP, false),
('Зерновые и хлопья', 'Категория для завтраков и хлопьев', CURRENT_TIMESTAMP, false),
('Бытовая химия', 'Категория для товаров бытовой химии', CURRENT_TIMESTAMP, false),
('Детское питание', 'Категория для продуктов детского питания', CURRENT_TIMESTAMP, false);

-- Вставка тестовых данных в таблицу products
INSERT INTO products (barcode, category_id, title, supplier_id, quantity, stock, supplier_price, retail_price, total_price, is_paid_to_supplier, is_vat_applicable, is_excise_applicable, unit, storage_location, created_at, is_deleted) VALUES
-- Фрукты
('123456789001', 1, 'Яблоки', 1, 100.0, 80.0, 50.0, 75.0, 5000.0, false, true, false, 'кг', 'Склад 1', CURRENT_TIMESTAMP, false),
('123456789002', 1, 'Груши', 2, 50.0, 40.0, 60.0, 90.0, 4500.0, false, true, false, 'кг', 'Склад 1', CURRENT_TIMESTAMP, false),
('123456789003', 1, 'Бананы', 1, 120.0, 100.0, 55.0, 80.0, 6600.0, false, true, false, 'кг', 'Склад 1', CURRENT_TIMESTAMP, false),
('123456789004', 1, 'Апельсины', 3, 70.0, 60.0, 65.0, 100.0, 7000.0, false, true, false, 'кг', 'Склад 1', CURRENT_TIMESTAMP, false),
('123456789005', 1, 'Киви', 4, 40.0, 30.0, 80.0, 120.0, 3200.0, false, true, false, 'кг', 'Склад 1', CURRENT_TIMESTAMP, false),

-- Овощи
('123456789006', 2, 'Картофель', 2, 200.0, 150.0, 20.0, 30.0, 6000.0, false, true, false, 'кг', 'Склад 2', CURRENT_TIMESTAMP, false),
('123456789007', 2, 'Морковь', 3, 180.0, 140.0, 25.0, 35.0, 6300.0, false, true, false, 'кг', 'Склад 2', CURRENT_TIMESTAMP, false),
('123456789008', 2, 'Лук', 1, 160.0, 130.0, 15.0, 25.0, 4000.0, false, true, false, 'кг', 'Склад 2', CURRENT_TIMESTAMP, false),
('123456789009', 2, 'Капуста', 4, 100.0, 80.0, 18.0, 28.0, 2800.0, false, true, false, 'кг', 'Склад 2', CURRENT_TIMESTAMP, false),
('123456789010', 2, 'Свекла', 5, 120.0, 90.0, 22.0, 32.0, 3840.0, false, true, false, 'кг', 'Склад 2', CURRENT_TIMESTAMP, false),

-- Молочные продукты
('123456789011', 3, 'Молоко', 2, 200.0, 160.0, 40.0, 60.0, 9600.0, false, true, false, 'литр', 'Холодильник 1', CURRENT_TIMESTAMP, false),
('123456789012', 3, 'Йогурт', 3, 150.0, 120.0, 50.0, 75.0, 7500.0, false, true, false, 'литр', 'Холодильник 1', CURRENT_TIMESTAMP, false),
('123456789013', 3, 'Кефир', 4, 180.0, 150.0, 45.0, 65.0, 8100.0, false, true, false, 'литр', 'Холодильник 1', CURRENT_TIMESTAMP, false),
('123456789014', 3, 'Сметана', 1, 130.0, 100.0, 55.0, 80.0, 7150.0, false, true, false, 'литр', 'Холодильник 1', CURRENT_TIMESTAMP, false),
('123456789015', 3, 'Творог', 5, 100.0, 70.0, 60.0, 90.0, 6000.0, false, true, false, 'кг', 'Холодильник 1', CURRENT_TIMESTAMP, false),

-- Мясо
('123456789016', 4, 'Говядина', 2, 80.0, 70.0, 120.0, 150.0, 9600.0, false, true, false, 'кг', 'Холодильник 2', CURRENT_TIMESTAMP, false),
('123456789017', 4, 'Утка', 3, 90.0, 75.0, 140.0, 180.0, 12600.0, false, true, false, 'кг', 'Холодильник 2', CURRENT_TIMESTAMP, false),
('123456789018', 4, 'Курица', 1, 150.0, 130.0, 100.0, 120.0, 15000.0, false, true, false, 'кг', 'Холодильник 2', CURRENT_TIMESTAMP, false),
('123456789019', 4, 'Баранина', 4, 60.0, 50.0, 150.0, 190.0, 9000.0, false, true, false, 'кг', 'Холодильник 2', CURRENT_TIMESTAMP, false),
('123456789020', 4, 'Индейка', 5, 70.0, 60.0, 130.0, 160.0, 9100.0, false, true, false, 'кг', 'Холодильник 2', CURRENT_TIMESTAMP, false),

-- Хлебобулочные изделия
('123456789031', 5, 'Белый хлеб', 1, 100.0, 90.0, 30.0, 50.0, 3000.0, false, true, false, 'шт', 'Хлебный склад', CURRENT_TIMESTAMP, false),
('123456789032', 5, 'Чёрный хлеб', 2, 80.0, 70.0, 35.0, 55.0, 2800.0, false, true, false, 'шт', 'Хлебный склад', CURRENT_TIMESTAMP, false),
('123456789033', 5, 'Батон', 3, 90.0, 80.0, 25.0, 40.0, 2250.0, false, true, false, 'шт', 'Хлебный склад', CURRENT_TIMESTAMP, false),
('123456789034', 5, 'Булочки с маком', 4, 50.0, 45.0, 20.0, 30.0, 1500.0, false, true, false, 'шт', 'Хлебный склад', CURRENT_TIMESTAMP, false),
('123456789035', 5, 'Ржаной хлеб', 5, 60.0, 50.0, 40.0, 60.0, 2400.0, false, true, false, 'шт', 'Хлебный склад', CURRENT_TIMESTAMP, false),

-- Напитки
('123456789041', 6, 'Кока-Кола', 1, 100.0, 90.0, 50.0, 70.0, 5000.0, false, true, false, 'литр', 'Склад напитков', CURRENT_TIMESTAMP, false),
('123456789042', 6, 'Пепси', 2, 100.0, 85.0, 50.0, 70.0, 4250.0, false, true, false, 'литр', 'Склад напитков', CURRENT_TIMESTAMP, false),
('123456789043', 6, 'Минеральная вода', 3, 120.0, 100.0, 25.0, 40.0, 3000.0, false, true, false, 'литр', 'Склад напитков', CURRENT_TIMESTAMP, false),
('123456789044', 6, 'Сок яблочный', 4, 80.0, 70.0, 40.0, 60.0, 3200.0, false, true, false, 'литр', 'Склад напитков', CURRENT_TIMESTAMP, false),
('123456789045', 6, 'Сок апельсиновый', 5, 70.0, 60.0, 45.0, 65.0, 2925.0, false, true, false, 'литр', 'Склад напитков', CURRENT_TIMESTAMP, false),

-- Снеки
('123456789066', 7, 'Чипсы', 1, 200.0, 180.0, 40.0, 60.0, 7200.0, false, true, false, 'упаковка', 'Склад снеков', CURRENT_TIMESTAMP, false),
('123456789067', 7, 'Сухарики', 2, 150.0, 130.0, 25.0, 45.0, 5850.0, false, true, false, 'упаковка', 'Склад снеков', CURRENT_TIMESTAMP, false),
('123456789068', 7, 'Орешки', 3, 100.0, 90.0, 50.0, 80.0, 8000.0, false, true, false, 'кг', 'Склад снеков', CURRENT_TIMESTAMP, false),
('123456789069', 7, 'Попкорн', 4, 180.0, 160.0, 35.0, 55.0, 9900.0, false, true, false, 'упаковка', 'Склад снеков', CURRENT_TIMESTAMP, false),
('123456789070', 7, 'Крекеры', 5, 130.0, 110.0, 30.0, 50.0, 6500.0, false, true, false, 'упаковка', 'Склад снеков', CURRENT_TIMESTAMP, false),

-- Замороженные продукты
('123456789071', 8, 'Замороженные овощи', 1, 200.0, 180.0, 60.0, 90.0, 10800.0, false, true, false, 'кг', 'Склад замороженных', CURRENT_TIMESTAMP, false),
('123456789072', 8, 'Замороженная пицца', 2, 150.0, 130.0, 120.0, 180.0, 19500.0, false, true, false, 'шт', 'Склад замороженных', CURRENT_TIMESTAMP, false),
('123456789073', 8, 'Замороженная рыба', 3, 100.0, 90.0, 200.0, 250.0, 22500.0, false, true, false, 'кг', 'Склад замороженных', CURRENT_TIMESTAMP, false),
('123456789074', 8, 'Мороженое', 4, 300.0, 270.0, 40.0, 60.0, 16200.0, false, true, false, 'шт', 'Склад замороженных', CURRENT_TIMESTAMP, false),
('123456789075', 8, 'Замороженные ягоды', 5, 120.0, 100.0, 70.0, 100.0, 12000.0, false, true, false, 'кг', 'Склад замороженных', CURRENT_TIMESTAMP, false),

-- Консервы
('123456789061', 9, 'Тушенка', 1, 200.0, 180.0, 100.0, 150.0, 30000.0, false, true, false, 'шт', 'Склад консервов', CURRENT_TIMESTAMP, false),
('123456789062', 9, 'Огурцы маринованные', 2, 300.0, 250.0, 50.0, 70.0, 17500.0, false, true, false, 'шт', 'Склад консервов', CURRENT_TIMESTAMP, false),
('123456789063', 9, 'Томатная паста', 3, 500.0, 450.0, 20.0, 35.0, 22500.0, false, true, false, 'шт', 'Склад консервов', CURRENT_TIMESTAMP, false),
('123456789064', 9, 'Грибы консервированные', 4, 150.0, 130.0, 40.0, 60.0, 7800.0, false, true, false, 'шт', 'Склад консервов', CURRENT_TIMESTAMP, false),
('123456789065', 9, 'Фасоль консервированная', 5, 180.0, 160.0, 30.0, 50.0, 9000.0, false, true, false, 'шт', 'Склад консервов', CURRENT_TIMESTAMP, false),

-- Личная гигиена
('123456789081', 10, 'Шампунь', 1, 150.0, 130.0, 70.0, 100.0, 10500.0, false, true, false, 'шт', 'Склад гигиены', CURRENT_TIMESTAMP, false),
('123456789082', 10, 'Мыло', 2, 200.0, 180.0, 30.0, 50.0, 9000.0, false, true, false, 'шт', 'Склад гигиены', CURRENT_TIMESTAMP, false),
('123456789083', 10, 'Зубная паста', 3, 100.0, 90.0, 40.0, 65.0, 6500.0, false, true, false, 'шт', 'Склад гигиены', CURRENT_TIMESTAMP, false),
('123456789084', 10, 'Зубная щетка', 4, 120.0, 100.0, 25.0, 40.0, 4800.0, false, true, false, 'шт', 'Склад гигиены', CURRENT_TIMESTAMP, false),
('123456789085', 10, 'Гель для душа', 5, 130.0, 110.0, 50.0, 75.0, 9750.0, false, true, false, 'шт', 'Склад гигиены', CURRENT_TIMESTAMP, false),

-- Кондитерсекие изделия
('123456789036', 11, 'Шоколад', 1, 200.0, 180.0, 80.0, 120.0, 16000.0, false, true, false, 'шт', 'Склад сладостей', CURRENT_TIMESTAMP, false),
('123456789037', 11, 'Печенье', 2, 150.0, 130.0, 50.0, 70.0, 10500.0, false, true, false, 'шт', 'Склад сладостей', CURRENT_TIMESTAMP, false),
('123456789038', 11, 'Конфеты', 3, 300.0, 270.0, 100.0, 150.0, 30000.0, false, true, false, 'кг', 'Склад сладостей', CURRENT_TIMESTAMP, false),
('123456789039', 11, 'Торт', 4, 30.0, 25.0, 250.0, 350.0, 8750.0, false, true, false, 'шт', 'Склад сладостей', CURRENT_TIMESTAMP, false),
('123456789040', 11, 'Пирожное', 5, 50.0, 45.0, 120.0, 180.0, 9000.0, false, true, false, 'шт', 'Склад сладостей', CURRENT_TIMESTAMP, false),

-- Крупы и макароны
('123456789046', 12, 'Рис', 1, 200.0, 180.0, 40.0, 60.0, 7200.0, false, true, false, 'кг', 'Склад круп', CURRENT_TIMESTAMP, false),
('123456789047', 12, 'Гречка', 2, 150.0, 130.0, 45.0, 65.0, 7150.0, false, true, false, 'кг', 'Склад круп', CURRENT_TIMESTAMP, false),
('123456789048', 12, 'Овсянка', 3, 170.0, 160.0, 30.0, 50.0, 5100.0, false, true, false, 'кг', 'Склад круп', CURRENT_TIMESTAMP, false),
('123456789049', 12, 'Пшено', 4, 140.0, 130.0, 35.0, 55.0, 4550.0, false, true, false, 'кг', 'Склад круп', CURRENT_TIMESTAMP, false),
('123456789050', 12, 'Перловка', 5, 160.0, 140.0, 25.0, 45.0, 5600.0, false, true, false, 'кг', 'Склад круп', CURRENT_TIMESTAMP, false),

-- Масла и соусы
('123456789091', 13, 'Оливковое масло', 1, 100.0, 90.0, 120.0, 160.0, 14400.0, false, true, false, 'л', 'Склад масла и соусов', CURRENT_TIMESTAMP, false),
('123456789092', 13, 'Подсолнечное масло', 2, 150.0, 130.0, 80.0, 110.0, 14300.0, false, true, false, 'л', 'Склад масла и соусов', CURRENT_TIMESTAMP, false),
('123456789093', 13, 'Соевый соус', 3, 80.0, 70.0, 60.0, 90.0, 5400.0, false, true, false, 'л', 'Склад масла и соусов', CURRENT_TIMESTAMP, false),
('123456789094', 13, 'Кетчуп', 4, 120.0, 100.0, 50.0, 75.0, 9000.0, false, true, false, 'л', 'Склад масла и соусов', CURRENT_TIMESTAMP, false),
('123456789095', 13, 'Майонез', 5, 130.0, 110.0, 55.0, 85.0, 11050.0, false, true, false, 'л', 'Склад масла и соусов', CURRENT_TIMESTAMP, false),
-- Рыба и морепродукты
('123456789051', 14, 'Лосось', 1, 50.0, 45.0, 300.0, 400.0, 18000.0, false, true, false, 'кг', 'Холодильник 3', CURRENT_TIMESTAMP, false),
('123456789052', 14, 'Тунец', 2, 40.0, 35.0, 350.0, 450.0, 15750.0, false, true, false, 'кг', 'Холодильник 3', CURRENT_TIMESTAMP, false),
('123456789053', 14, 'Сельдь', 3, 100.0, 85.0, 100.0, 150.0, 12750.0, false, true, false, 'кг', 'Холодильник 3', CURRENT_TIMESTAMP, false),
('123456789054', 14, 'Креветки', 4, 80.0, 70.0, 250.0, 350.0, 24500.0, false, true, false, 'кг', 'Холодильник 3', CURRENT_TIMESTAMP, false),
('123456789055', 14, 'Кальмары', 5, 60.0, 50.0, 200.0, 300.0, 15000.0, false, true, false, 'кг', 'Холодильник 3', CURRENT_TIMESTAMP, false),

-- Специи и приправы
('123456789101', 15, 'Черный перец', 1, 50.0, 45.0, 40.0, 60.0, 2700.0, false, true, false, 'кг', 'Склад специй', CURRENT_TIMESTAMP, false),
('123456789102', 15, 'Корица', 2, 40.0, 35.0, 55.0, 80.0, 3200.0, false, true, false, 'кг', 'Склад специй', CURRENT_TIMESTAMP, false),
('123456789103', 15, 'Карри', 3, 60.0, 50.0, 45.0, 70.0, 3150.0, false, true, false, 'кг', 'Склад специй', CURRENT_TIMESTAMP, false),
('123456789104', 15, 'Базилик', 4, 70.0, 60.0, 50.0, 75.0, 3750.0, false, true, false, 'кг', 'Склад специй', CURRENT_TIMESTAMP, false),
('123456789105', 15, 'Чесночный порошок', 5, 80.0, 70.0, 60.0, 90.0, 5400.0, false, true, false, 'кг', 'Склад специй', CURRENT_TIMESTAMP, false),

-- Алкоголь
('123456789021', 16, 'Водка', 1, 50.0, 40.0, 200.0, 250.0, 10000.0, false, true, true, 'литр', 'Алкосклад 1', CURRENT_TIMESTAMP, false),
('123456789022', 16, 'Виски', 2, 30.0, 25.0, 400.0, 500.0, 12000.0, false, true, true, 'литр', 'Алкосклад 1', CURRENT_TIMESTAMP, false),
('123456789023', 16, 'Коньяк', 3, 40.0, 35.0, 350.0, 450.0, 14000.0, false, true, true, 'литр', 'Алкосклад 1', CURRENT_TIMESTAMP, false),
('123456789024', 16, 'Ром', 4, 20.0, 15.0, 300.0, 380.0, 7600.0, false, true, true, 'литр', 'Алкосклад 1', CURRENT_TIMESTAMP, false),
('123456789025', 16, 'Шампанское', 5, 60.0, 50.0, 150.0, 200.0, 12000.0, false, true, true, 'литр', 'Алкосклад 1', CURRENT_TIMESTAMP, false),

-- Табачные изделия
('123456789026', 17, 'Сигареты Marlboro', 1, 100.0, 90.0, 120.0, 150.0, 12000.0, false, true, true, 'пачка', 'Табачный склад', CURRENT_TIMESTAMP, false),
('123456789027', 17, 'Сигареты Parliament', 2, 80.0, 70.0, 130.0, 160.0, 10400.0, false, true, true, 'пачка', 'Табачный склад', CURRENT_TIMESTAMP, false),
('123456789028', 17, 'Сигары Cohiba', 3, 40.0, 30.0, 500.0, 700.0, 20000.0, false, true, true, 'пачка', 'Табачный склад', CURRENT_TIMESTAMP, false),
('123456789029', 17, 'Сигареты Winston', 4, 90.0, 80.0, 110.0, 140.0, 9900.0, false, true, true, 'пачка', 'Табачный склад', CURRENT_TIMESTAMP, false),
('123456789030', 17, 'Сигареты Camel', 5, 60.0, 50.0, 100.0, 130.0, 7800.0, false, true, true, 'пачка', 'Табачный склад', CURRENT_TIMESTAMP, false),

-- Зерновые и хлопья
('123456789201', 18, 'Кукурузные хлопья', 1, 150.0, 130.0, 30.0, 50.0, 7500.0, false, true, false, 'кг', 'Склад хлопьев', CURRENT_TIMESTAMP, false),
('123456789202', 18, 'Овсяные хлопья', 2, 200.0, 180.0, 25.0, 40.0, 7200.0, false, true, false, 'кг', 'Склад хлопьев', CURRENT_TIMESTAMP, false),
('123456789203', 18, 'Шоколадные шарики', 3, 100.0, 90.0, 35.0, 55.0, 4950.0, false, true, false, 'кг', 'Склад хлопьев', CURRENT_TIMESTAMP, false),
('123456789204', 18, 'Пшеничные хлопья', 4, 120.0, 100.0, 20.0, 35.0, 4200.0, false, true, false, 'кг', 'Склад хлопьев', CURRENT_TIMESTAMP, false),
('123456789205', 18, 'Мюсли с фруктами', 5, 140.0, 120.0, 40.0, 60.0, 8400.0, false, true, false, 'кг', 'Склад хлопьев', CURRENT_TIMESTAMP, false),

-- Бытовая химия
('123456789056', 19, 'Стиральный порошок', 1, 50.0, 40.0, 120.0, 150.0, 7500.0, false, true, false, 'кг', 'Склад бытовой химии', CURRENT_TIMESTAMP, false),
('123456789057', 19, 'Моющее средство', 2, 60.0, 50.0, 80.0, 120.0, 6000.0, false, true, false, 'л', 'Склад бытовой химии', CURRENT_TIMESTAMP, false),
('123456789058', 19, 'Отбеливатель', 3, 70.0, 60.0, 90.0, 130.0, 7800.0, false, true, false, 'л', 'Склад бытовой химии', CURRENT_TIMESTAMP, false),
('123456789059', 19, 'Чистящее средство', 4, 100.0, 90.0, 60.0, 80.0, 7200.0, false, true, false, 'л', 'Склад бытовой химии', CURRENT_TIMESTAMP, false),
('123456789060', 19, 'Освежитель воздуха', 5, 120.0, 100.0, 50.0, 75.0, 7500.0, false, true, false, 'л', 'Склад бытовой химии', CURRENT_TIMESTAMP, false),

-- Детское питание
('223456789101', 20, 'Молочная смесь Nestle', 1, 100.0, 90.0, 150.0, 200.0, 15000.0, false, true, false, 'шт', 'Склад детского питания', CURRENT_TIMESTAMP, false),
('223456789102', 20, 'Пюре фруктовое Gerber', 2, 200.0, 180.0, 50.0, 70.0, 10000.0, false, true, false, 'шт', 'Склад детского питания', CURRENT_TIMESTAMP, false),
('223456789103', 20, 'Каша Heinz', 3, 150.0, 130.0, 100.0, 140.0, 15000.0, false, true, false, 'шт', 'Склад детского питания', CURRENT_TIMESTAMP, false),
('223456789104', 20, 'Печенье HIPP', 4, 120.0, 100.0, 80.0, 120.0, 9600.0, false, true, false, 'шт', 'Склад детского питания', CURRENT_TIMESTAMP, false),
('223456789105', 20, 'Сок Агуша', 5, 180.0, 160.0, 30.0, 50.0, 9000.0, false, true, false, 'литр', 'Склад детского питания', CURRENT_TIMESTAMP, false);

