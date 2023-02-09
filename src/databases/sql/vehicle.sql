-- TODO CREATE TABLE VEHICLE
CREATE TABLE IF NOT EXISTS vehicle (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    location VARCHAR NOT NULL,
    description TEXT,
    status VARCHAR NOT NULL,
    stock INT NOT NULL,
    category_id INT,
    price INT,
    rating INT,
    image VARCHAR,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT fk_category
        FOREIGN KEY (category_id) 
            REFERENCES category (id)
);

-- TODO DELETE TABLE VEHICLE
-- DROP TABLE IF EXISTS vehicle;

--TODO SHOW ALL VEHICLE
-- SELECT * FROM vehicle;

-- TODO RESTART ID COUNT
-- ALTER SEQUENCE vehicle_id_seq RESTART WITH 1

-- TODO DELETE
-- DELETE FROM vehicle