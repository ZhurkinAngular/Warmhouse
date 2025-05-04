from sqlalchemy import Column, Integer, String, Float, DateTime, func
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class SensorDB(Base):
    __tablename__ = 'sensors'

    id = Column(Integer, primary_key=True)
    name = Column(String(100), nullable=False)
    type = Column(String(50), nullable=False)
    location = Column(String(100), nullable=False)
    value = Column(Float, default=0)
    unit = Column(String(20))
    status = Column(String(20), nullable=False, default='inactive')
    last_updated = Column(DateTime(timezone=True), nullable=False, default=func.now())
    created_at = Column(DateTime(timezone=True), nullable=False, default=func.now())