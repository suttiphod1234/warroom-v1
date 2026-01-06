from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(title="War Room AI Service", version="1.0.0")

class HealthCheck(BaseModel):
    status: str
    service: str

@app.get("/", tags=["Root"])
async def read_root():
    return {"message": "War Room AI Service is Running"}

@app.get("/health", response_model=HealthCheck, tags=["Health"])
async def health_check():
    return HealthCheck(status="ok", service="ai-service")
