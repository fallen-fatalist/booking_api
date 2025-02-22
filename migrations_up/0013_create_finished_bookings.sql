CREATE TABLE finished_bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id),
    computer_id INT NOT NULL REFERENCES computers_specs(computer_id),
    package_name VARCHAR(50) NOT NULL REFERENCES packages(package_name),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    total_price INT NOT NULL,
    status VARCHAR(20) DEFAULT 'finished' CHECK (status IN ('finished', 'canceled')),
    created_at TIMESTAMP NOT NULL
);
