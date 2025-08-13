-- Создаём таблицу датчиков
CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL PRIMARY KEY,
    sensor_id VARCHAR(10) NOT NULL,
    location VARCHAR(50) NOT NULL,
    temperature NUMERIC(5,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Вставляем несколько тестовых датчиков
INSERT INTO sensors (sensor_id, location, temperature)
VALUES
('1', 'Living Room', 22.5),
('2', 'Bedroom', 21.0),
('3', 'Kitchen', 23.3);
