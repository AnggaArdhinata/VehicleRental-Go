-- TODO CREATE TABLE USERS
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    birthday DATE,
    gender VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    image VARCHAR,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
    
)

-- TODO DELETE TABLE USERS
-- DROP TABLE IF EXISTS users;

--TODO SHOW ALL USERS
-- SELECT * FROM users;

--TODO INSERT USERS
-- INSERT INTO users (name, email, password, birthday, gender, address, phone, image)
-- VALUES ('Ardi', 'ardi@mail.com', 'ardi123', '2023-01-07', 'male', 'semarang', '0845345343', 'picture')

-- TODO DELETE ALL USERS
-- DELETE FROM users;

-- TODO RESTART ID USERS TO 1
-- ALTER SEQUENCE users_id_seq RESTART WITH 1