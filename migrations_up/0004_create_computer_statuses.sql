CREATE TABLE computers_statuses (
    computer_id SERIAL NOT NULL PRIMARY KEY,
    status VARCHAR(100) NOT NULL CHECK(status in ('busy', 'available', 'pending', 'not working')),
    updated_at TIMESTAMP DEFAULT NOW()
);