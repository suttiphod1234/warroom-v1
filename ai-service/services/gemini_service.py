import os
import google.generativeai as genai

# Configure Gemini (Mock if no key)
GOOGLE_API_KEY = os.getenv('GOOGLE_API_KEY')
if GOOGLE_API_KEY:
    genai.configure(api_key=GOOGLE_API_KEY)

def analyze_sentiment(text: str):
    """
    Uses Gemini to analyze sentiment.
    """
    if not GOOGLE_API_KEY:
        # Mock response if no key provided
        return {"score": 0.5, "label": "neutral (mock)"}
    
    try:
        model = genai.GenerativeModel('gemini-pro')
        response = model.generate_content(f"Analyze the sentiment of this text (positive/negative/neutral): '{text}'")
        return {"result": response.text}
    except Exception as e:
        return {"error": str(e)}

def generate_summary(posts: list):
    """
    Generates a summary of the provided posts.
    """
    if not GOOGLE_API_KEY:
        return "Simulated summary: The public sentiment is mixed regarding this topic."

    try:
        text_content = "\n".join([p['content'] for p in posts])
        model = genai.GenerativeModel('gemini-pro')
        prompt = f"Summarize the general sentiment and key points from these social media posts:\n\n{text_content}"
        response = model.generate_content(prompt)
        return response.text
    except Exception as e:
        return f"Error generating summary: {str(e)}"
