from fastapi import FastAPI, Depends, HTTPException
from pydantic import BaseModel
from datetime import datetime
import random
from sqlalchemy.future import select
from models import SensorDB
from sqlalchemy.ext.asyncio import AsyncSession
from database import get_db

app = FastAPI()

class TemperatureResponse(BaseModel):
    value: float
    unit: str
    timestamp: str
    location: str
    status: str
    sensor_id: str
    sensor_type: str
    description: str

def generate_temperature_data(location: str, sensor_id: str):
    return TemperatureResponse(
        value=round(random.uniform(-50, 50), 2),
        unit="Celsius",
        timestamp=datetime.utcnow().isoformat() + "Z",
        location=location,
        status="active",
        sensor_id=sensor_id,
        sensor_type="DS18B20",
        description="Simulated temperature sensor"
    )

@app.get("/temperature")
async def get_temperature(location: str = "", db: AsyncSession = Depends(get_db)):
    result = await db.execute(select(SensorDB).where(SensorDB.location == location))
    sensors = result.scalars().all()

    if not sensors:
        raise HTTPException(status_code=404, detail="No sensors found at this location")

    sensor = sensors[0]
    return generate_temperature_data(location=sensor.location, sensor_id=str(sensor.id))


@app.get("/temperature/{sensor_id}")
async def get_temperature_by_id(sensor_id: int, db: AsyncSession = Depends(get_db)):
    result = await db.execute(select(SensorDB).where(SensorDB.id == sensor_id))
    sensor = result.scalar_one_or_none()

    if not sensor:
        raise HTTPException(status_code=404, detail="Sensor ID not found")

    return generate_temperature_data(location=sensor.location, sensor_id=str(sensor.id))