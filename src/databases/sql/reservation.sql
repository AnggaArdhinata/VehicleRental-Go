-- TODO CREATE TABLE RESERVATION
CREATE TABLE IF NOT EXISTS reservation (
    id BIGSERIAL PRIMARY KEY,
    user_id INT,
    vehicle_id INT,
    qty INT,
    start_date DATE,
    return_date DATE,
    total_payment INT,
    isbooked BOOLEAN,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
        CONSTRAINT fk_user_vehicle
            FOREIGN KEY (user_id) REFERENCES users(id),
            FOREIGN KEY (vehicle_id) REFERENCES vehicle(id)
)

-- TODO DELETE TABLE RESERVATION
-- DROP TABLE reservation;

-- TODO SHOW DATA RESERVATION
-- SELECT * FROM reservation;