using System;
using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;
using temperature_api.Models;

namespace temperature_api.Context;

public partial class SmarthomeContext : DbContext
{
    public SmarthomeContext()
    {
    }

    public SmarthomeContext(DbContextOptions<SmarthomeContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Sensor> Sensors { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        => optionsBuilder.UseNpgsql("Host=smarthome-postgres;Port=5432;Database=smarthome;Username=postgres;Password=postgres");

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Sensor>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("sensors_pkey");

            entity.ToTable("sensors");

            entity.HasIndex(e => e.Location, "idx_sensors_location");

            entity.HasIndex(e => e.Status, "idx_sensors_status");

            entity.HasIndex(e => e.Type, "idx_sensors_type");

            entity.Property(e => e.Id).HasColumnName("id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnName("created_at");
            entity.Property(e => e.LastUpdated)
                .HasDefaultValueSql("now()")
                .HasColumnName("last_updated");
            entity.Property(e => e.Location)
                .HasMaxLength(100)
                .HasColumnName("location");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
            entity.Property(e => e.Status)
                .HasMaxLength(20)
                .HasDefaultValueSql("'inactive'::character varying")
                .HasColumnName("status");
            entity.Property(e => e.Type)
                .HasMaxLength(50)
                .HasColumnName("type");
            entity.Property(e => e.Unit)
                .HasMaxLength(20)
                .HasColumnName("unit");
            entity.Property(e => e.Value)
                .HasDefaultValueSql("0")
                .HasColumnName("value");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
