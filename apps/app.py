from flask import Flask, request, jsonify
import random
import psycopg2
from psycopg2.extras import RealDictCursor
import os

app = Flask(__name__)

# Конфигурация PostgreSQL
DB_HOST = os.getenv("POSTGRES_HOST", "localhost")
DB_NAME = os.getenv("POSTGRES_DB", "smarthome")
DB_USER = os.getenv("POSTGRES_USER", "smarthome")
DB_PASSWORD = os.getenv("POSTGRES_PASSWORD", "smarthome")
DB_PORT = os.getenv("POSTGRES_PORT", 5432)

# Подключение к базе
conn = psycopg2.connect(
    host=DB_HOST,
    dbname=DB_NAME,
    user=DB_USER,
    password=DB_PASSWORD,
    port=DB_PORT
)
conn.autocommit = True
cur = conn.cursor(cursor_factory=RealDictCursor)

@app.route("/temperature", methods=["GET"])
def get_temperature():
    sensor_id = request.args.get("sensorId", "")
    location = request.args.get("location", "")

    # Определяем location по sensor_id
    if location == "":
        if sensor_id == "1":
            location = "Living Room"
        elif sensor_id == "2":
            location = "Bedroom"
        elif sensor_id == "3":
            location = "Kitchen"
        else:
            location = "Unknown"

    # Определяем sensor_id по location
    if sensor_id == "":
        if location == "Living Room":
            sensor_id = "1"
        elif location == "Bedroom":
            sensor_id = "2"
        elif location == "Kitchen":
            sensor_id = "3"
        else:
            sensor_id = "0"

    # Генерация случайной температуры
    temperature = round(random.uniform(18.0, 28.0), 1)

    # Сохраняем или обновляем данные в PostgreSQL
    cur.execute("""
        INSERT INTO sensors (sensor_id, location, temperature)
        VALUES (%s, %s, %s)
        ON CONFLICT (sensor_id) DO UPDATE
        SET temperature = EXCLUDED.temperature,
            updated_at = NOW()
    """, (sensor_id, location, temperature))

    return jsonify({
        "sensorId": sensor_id,
        "location": location,
        "temperature": temperature
    })

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081)
