from datetime import datetime
import random

def scrape_hashtag(platform: str, hashtag: str):
    """
    Simulates scraping social media posts.
    In a real scenario, this would use Selenium/Playwright/Official APIs.
    """
    results = []
    
    # Simulate 5-10 posts
    count = random.randint(5, 10)
    
    sentiments = ["positive", "negative", "neutral"]
    
    for i in range(count):
        sentiment = random.choice(sentiments)
        content = f"This is a simulated post #{i+1} about #{hashtag} on {platform}. The sentiment seems {sentiment}."
        
        post = {
            "platform": platform,
            "external_id": f"{platform}_{random.randint(10000, 99999)}",
            "content": content,
            "author": f"user_{random.randint(1, 100)}",
            "likes": random.randint(0, 500),
            "shares": random.randint(0, 100),
            "scraped_at": datetime.now().isoformat()
        }
        results.append(post)
        
    return results
