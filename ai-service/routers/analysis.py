from fastapi import APIRouter, HTTPException
from pydantic import BaseModel
from scrapers.social_scraper import scrape_hashtag
from services.gemini_service import generate_summary

router = APIRouter(
    prefix="/analyze",
    tags=["analysis"]
)

class TrendRequest(BaseModel):
    hashtag: str
    platform: str = "all"

@router.post("/trend")
async def analyze_trend(req: TrendRequest):
    # 1. Scrape Data (Simulated)
    posts = scrape_hashtag(req.platform, req.hashtag)
    
    if not posts:
        raise HTTPException(status_code=404, detail="No posts found")
    
    # 2. Generate Summary (Gemini)
    summary = generate_summary(posts)
    
    return {
        "hashtag": req.hashtag,
        "post_count": len(posts),
        "summary": summary,
        "sample_posts": posts[:3]
    }
