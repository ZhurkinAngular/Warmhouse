const express = require('express');
const app = express();
const port = process.env.PORT || 8081;

// Middleware to parse JSON
app.use(express.json());

// Helper function to generate random temperature between min and max
function getRandomTemperature(min = 15, max = 30) {
  return parseFloat((Math.random() * (max - min) + min).toFixed(1));
}

// Helper function to get current timestamp in ISO format
function getCurrentTimestamp() {
  return new Date().toISOString();
}

// GET /temperature endpoint
app.get('/temperature', (req, res) => {
  let location = req.query.location || '';
  let sensorID = '';
  
  // If no location is provided, use a default based on sensor ID
  if (location === '') {
    sensorID = '0';
    location = 'Unknown';
  } else {
    // If location is provided, generate sensor ID based on location
    switch (location) {
      case 'Living Room':
        sensorID = '1';
        break;
      case 'Bedroom':
        sensorID = '2';
        break;
      case 'Kitchen':
        sensorID = '3';
        break;
      default:
        sensorID = '0';
    }
  }
  
  // Generate random temperature data
  const temperature = getRandomTemperature();
  const status = temperature > 25 ? 'active' : 'inactive';
  
  // Return temperature data
  res.json({
    value: temperature,
    unit: '°C',
    timestamp: getCurrentTimestamp(),
    location: location,
    status: status,
    sensor_id: sensorID,
    sensor_type: 'temperature',
    description: `Temperature sensor in ${location}`
  });
});

// GET /temperature/:sensorID endpoint
app.get('/temperature/:sensorID', (req, res) => {
  const sensorID = req.params.sensorID;
  let location = '';
  
  // Determine location based on sensor ID
  switch (sensorID) {
    case '1':
      location = 'Living Room';
      break;
    case '2':
      location = 'Bedroom';
      break;
    case '3':
      location = 'Kitchen';
      break;
    default:
      location = 'Unknown';
  }
  
  // Generate random temperature data
  const temperature = getRandomTemperature();
  const status = temperature > 25 ? 'active' : 'inactive';
  
  // Return temperature data
  res.json({
    value: temperature,
    unit: '°C',
    timestamp: getCurrentTimestamp(),
    location: location,
    status: status,
    sensor_id: sensorID,
    sensor_type: 'temperature',
    description: `Temperature sensor in ${location}`
  });
});

// Health check endpoint
app.get('/health', (req, res) => {
  res.status(200).json({ status: 'ok' });
});

// Start the server
app.listen(port, () => {
  console.log(`Temperature API listening at http://localhost:${port}`);
});