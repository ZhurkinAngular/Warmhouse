using System;
using System.Collections.Generic;

namespace temperature_api.Models;

public partial class Sensor
{
    public int Id { get; set; }

    public string Name { get; set; } = null!;

    public string Type { get; set; } = null!;

    public string Location { get; set; } = null!;

    public double? Value { get; set; }

    public string? Unit { get; set; }

    public string Status { get; set; } = null!;

    public DateTime LastUpdated { get; set; }

    public DateTime CreatedAt { get; set; }
}
