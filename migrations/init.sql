CREATE TABLE IF NOT EXISTS item (
    itemId SERIAL PRIMARY KEY,
    ean TEXT,
    title TEXT NOT NULL,
    brand TEXT default '',
    amount INTEGER default 1,
    note TEXT,
    expiredDate DATE default null
);

INSERT INTO item (ean, title, brand, amount, note, expiredDate) 
VALUES (
    '4901002133528',
    'Golden Curry Japanese Curry Sauce Roux Block Medium Hot 198g S & B Japan',
    'Golden Curry',
    1,
    'Mirin for cooking',
    '2024-02-25'
);