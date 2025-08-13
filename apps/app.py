from flask import Flask, request, jsonify
import psycopg2
from psycopg2.extras import RealDictCursor
import random

app = Flask(__name__)

# Настройки БД
DB_HOST = "postgres"  # имя сервиса в docker-compose
DB_NAME = "smarthome"
DB_USER = "smarthome"
DB_PASS = "smarthome"

def get_connection():
    return psycopg2.connect(
        host=DB_HOST,
        database=DB_NAME,
        user=DB_USER,
        password=DB_PASS
    )

def update_temperature(sensor_id, temperature):
    print(f"[LOG] Updating sensor {sensor_id} to {temperature}")
    conn = get_connection()
    cur = conn.cursor()
    cur.execute("""
        UPDATE sensors
        SET temperature = %s, updated_at = NOW()
        WHERE sensor_id = %s
    """, (temperature, str(sensor_id)))  # str() для гарантии совпадения с типом sensor_id в БД
    conn.commit()
    cur.close()
    conn.close()

def get_sensor(sensor_id):
    conn = get_connection()
    cur = conn.cursor(cursor_factory=RealDictCursor)
    cur.execute("SELECT * FROM sensors WHERE sensor_id = %s", (str(sensor_id),))
    row = cur.fetchone()
    cur.close()
    conn.close()
    return row

@app.route("/temperature")
def temperature():
    sensor_id = request.args.get("sensorId")
    if not sensor_id:
        return jsonify({"error": "sensorId is required"}), 400

    # Генерация случайной температуры для теста
    temp = round(random.uniform(19.0, 26.0), 1)

    # Обновление в базе
    update_temperature(sensor_id, temp)

    # Возврат текущего состояния сенсора
    sensor = get_sensor(sensor_id)
    return jsonify({
        "sensorId": sensor["sensor_id"],
        "location": sensor["location"],
        "temperature": sensor["temperature"]
    })

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081)
