import datetime
from random import randint

import uvicorn

from fastapi import FastAPI, Query
from pydantic import BaseModel, Field
from pydantic_settings import BaseSettings, SettingsConfigDict


class AppSettings(BaseSettings):
    port: int = 8081
    host: str = "0.0.0.0"
    reload: bool = True
    model_config = SettingsConfigDict(populate_by_name=True, from_attributes=True, env_file=".env", extra="ignore")


settings = AppSettings()
app = FastAPI()


class Response(BaseModel):
    value: float = Field(default_factory=lambda: randint(0, 30))
    unit: str = "celsius"
    timestamp: datetime.datetime = Field(default_factory=datetime.datetime.now)
    location: str
    status: str = "active"
    sensor_id: str
    sensor_type: str = "temperature"
    description: str = "A temperature sensor"



@app.get("/temperature", response_model=Response)
async def root(location: str | None = None, sensor_id: str | None = Query(default=None, alias="sensorId")) -> Response:
    sensors_db = {"1": "Living Room", "2": "Bedroom", "3": "Kitchen"}
    locations_db = {"Living Room": "1", "Bedroom": "2", "Kitchen": "3"}

    if location is None:
        location = sensors_db.get(sensor_id, "Unknown")

    if sensor_id is None:
        sensor_id = locations_db.get(location, "0")

    return Response(location=location, sensor_id=sensor_id)


if __name__ == '__main__':
    uvicorn.run("app:app", host=settings.host, port=settings.port, reload=settings.reload)

