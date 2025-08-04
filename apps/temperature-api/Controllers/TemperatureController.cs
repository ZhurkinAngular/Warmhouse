using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Http.HttpResults;
using Microsoft.AspNetCore.Mvc;
using temperature_api.Context;
using temperature_api.Models;

namespace temperature_api.Controllers
{
    [ApiVersion("1")]
    [Route("api/v{api-version:apiVersion}/")]
    [ApiController]
    public class SensorsController : ControllerBase
    {
        [HttpGet]
        [Route("temperature")]
        public IActionResult GetLocationTemperature(string? sensorId=null, string? location = null)
        {
            if (string.IsNullOrEmpty(sensorId)
                && string.IsNullOrEmpty(location))
            {
                return BadRequest();
            }

            if (!string.IsNullOrEmpty(sensorId))
            {
                location = sensorId switch
                {
                    "1" => "Living Room",
                    "2" => "Bedroom",
                    "3" => "Kitchen",
                    _ => "Unknown"
                };
            }

            if (!string.IsNullOrEmpty(location))
            {
                sensorId = location switch
                {
                    "Living Room" => "1",
                    "Bedroom" => "2",
                    "Kitchen" => "3",
                    _ => "0"
                };
            }

            var rnd = new Random();
            var temp = (double)rnd.Next(0, 40);

            var data = new TemperatureData()
            {
                Value = temp,
                Location = location,
                SensorID = sensorId,
                Description = $"temperature sensor at {location}"
            };

            return Ok(data);
        }

        [HttpGet]
        [Route("getAllSensors")]
        public IActionResult GetAllSensors()
        {
            var result = new List<Sensor>();

            using (SmarthomeContext context = new SmarthomeContext())
            {
                result = context.Sensors.ToList();
            }

            return Ok(result);
        }

        [HttpPost]
        [Route("createSensor")]
        public IActionResult CreateSensor(Sensor sensor)
        {
            using (SmarthomeContext context = new SmarthomeContext())
            {
                 context.Sensors.Add(sensor);
            }

            return Created();
        }
    }

}
