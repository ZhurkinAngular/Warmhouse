import express, { Request, Response } from 'express';

const app = express();
const port = 3000;

type TTemperatureData = {
  Value: number;
  Unit: string;
  Timestamp: string;
  Location: string;
  Status: string;
  SensorID: string;
  SensorType: string;
  Description: string;
}

const fetchTemperatureData = async (location: string, sensorId: string): Promise<TTemperatureData> => {
  // If no location is provided, use a default based on sensor ID
  if (location === "") {
    switch (sensorId) {
      case "1":
        location = "Living Room"
      case "2":
        location = "Bedroom"
      case "3":
        location = "Kitchen"
      default:
        location = "Unknown"
    }
  }

  // If no sensor ID is provided, generate one based on location
  if (sensorId == "") {
    switch (location) {
      case "Living Room":
        sensorId = "1"
      case "Bedroom":
        sensorId = "2"
      case "Kitchen":
        sensorId = "3"
      default:
        sensorId = "0"
    }
  }
  const value = Math.floor(Math.random() * 131) - 30;
  return {
    Value: value,
    Unit: '°C',
    Timestamp: new Date().toISOString(),
    Location: location,
    Status: 'active',
    SensorID: sensorId,
    SensorType: 'temperature',
    Description: `Temperature sensor in ${location}`,
  };
};

app.get('/temperature/:sensorId', async (req: Request, res: Response) => {
  const location = req.query.location;
  const { sensorId } = req.params;

  if (!sensorId && !location) {
    return res.status(400).json({ error: 'Sensor id or location required' });
  }

  res.json(await fetchTemperatureData(location as string || '', sensorId || ''));
});

app.get('/health', async (req: Request, res: Response) => {
  res.json({ status: 'ok' });
});


app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
