import random
from datetime import datetime, timezone

from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/temperature/', methods=['GET'])
def get_temperature_by_location():
    temperature = round(random.uniform(20.0, 30.0))
    return jsonify({
        'value': temperature,
        'status': 'active',
        'timestamp': datetime.now(timezone.utc).isoformat()
    })


@app.route('/temperature/<id>', methods=['GET'])
def get_temperature_by_id(id):
    temperature = round(random.uniform(20.0, 30.0))
    return jsonify({
        'value': temperature,
        'status': 'active',
        'timestamp': datetime.now(timezone.utc).isoformat()
    })


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081, debug=True)