CREATE TABLE packages (
    package_name VARCHAR(50) PRIMARY KEY,
    price INT NOT NULL,
    startTime TIME,
    endTime TIME,
    created_at TIMESTAMP DEFAULT NOW()
)