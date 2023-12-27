CREATE TABLE invoice (
    id SERIAL PRIMARY KEY,
    customer_id INT,
    date DATE
);

CREATE TABLE invoice_details (
    id SERIAL PRIMARY KEY,
    invoice_id INT,
    product_id INT,
    quantity INT
);

-- Generowanie 1000 faktur
INSERT INTO invoice (customer_id, date)
SELECT
    generate_series(1, 1000), -- generowanie customer_id od 1 do 1000
    CURRENT_DATE
;

-- Generowanie od 5 do 10 szczegółów faktury dla każdej faktury
DO
$$
DECLARE
    invoice_id INT;
BEGIN
    FOR invoice_id IN SELECT id FROM invoice
    LOOP
        INSERT INTO invoice_details (invoice_id, product_id, quantity)
        SELECT
            invoice_id,
            generate_series(1, 10), -- generowanie product_id od 1 do 10
            generate_series(1, 10)  -- generowanie quantity od 1 do 10
        LIMIT (random() * 6 + 5)::INT; -- generowanie losowej liczby od 5 do 10
    END LOOP;
END
$$;