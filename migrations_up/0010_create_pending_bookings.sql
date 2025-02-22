CREATE TABLE pending_bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    computer_id INT NOT NULL REFERENCES computers_specs(computer_id) ON DELETE CASCADE,
    package_name VARCHAR(50) NOT NULL NOT NULL REFERENCES packages(package_name) ,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    total_price INT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'active')),
    created_at TIMESTAMP DEFAULT NOW()
);