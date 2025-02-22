CREATE TABLE computers_specs (
    computer_id SERIAL NOT NULL PRIMARY KEY,
    cpu VARCHAR(50) NOT NULL,
    gpu VARCHAR(50) NOT NULL,
    ram VARCHAR(50) NOT NULL,
    ssd VARCHAR(50) NOT NULL,
    hdd VARCHAR(50) NOT NULL,
    monitor VARCHAR(50) NOT NULL,
    keyboard VARCHAR(50) NOT NULL,
    headset VARCHAR(50) NOT NULL,
    mouse VARCHAR(50) NOT NULL,
    -- TODO: Add mousepad
    created_at TIMESTAMP DEFAULT NOW()
);



