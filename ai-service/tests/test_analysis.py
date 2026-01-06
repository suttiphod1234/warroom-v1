from fastapi.testclient import TestClient
from main import app
from unittest.mock import patch

client = TestClient(app)

def test_analyze_trend():
    # Mock the Gemini service to avoid real API calls
    with patch("services.gemini_service.generate_summary") as mock_summary:
        mock_summary.return_value = "Mocked Summary of the trend."
        
        response = client.post("/analyze/trend", json={"hashtag": "election", "platform": "facebook"})
        
        assert response.status_code == 200
        data = response.json()
        assert data["hashtag"] == "election"
        assert "summary" in data
        assert data["summary"] == "Mocked Summary of the trend."
        assert len(data["sample_posts"]) > 0
