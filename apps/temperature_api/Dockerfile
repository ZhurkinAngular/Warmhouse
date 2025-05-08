FROM python:3.10-slim

WORKDIR /app

COPY temperature_api.py .

RUN pip install flask

EXPOSE 8081

CMD ["python", "temperature_api.py"]
