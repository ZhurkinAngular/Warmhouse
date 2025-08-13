from flask import Flask, request, jsonify
import random
import psycopg2
from psycopg2.extras import RealDictCursor

app = Flask(__name__)

# Настройки БД
DB_HOST = "postgres"
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

@app.route("/temperature", methods=["GET"])
def get_temperature():
    sensor_id = request.args.get("sensorId", "")
    location = request.args.get("location", "")

    # Определяем defaults
    if location == "":
        location = {"1": "Living Room", "2": "Bedroom", "3": "Kitchen"}.get(sensor_id, "Unknown")
    if sensor_id == "":
        sensor_id = {"Living Room": "1", "Bedroom": "2", "Kitchen": "3"}.get(location, "0")

    temperature = round(random.uniform(18.0, 28.0), 1)

    # Сохраняем в БД
    try:
        conn = get_connection()
        cur = conn.cursor()
        cur.execute("""
            INSERT INTO sensors (sensor_id, location, temperature)
            VALUES (%s, %s, %s)
            ON CONFLICT (sensor_id) DO UPDATE 
            SET temperature = EXCLUDED.temperature,
                updated_at = NOW()
        """, (sensor_id, location, temperature))
        conn.commit()
        cur.close()
        conn.close()
    except Exception as e:
        print("DB error:", e)

    return jsonify({
        "sensorId": sensor_id,
        "location": location,
        "temperature": temperature
    })

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081)
