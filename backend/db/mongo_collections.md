# MongoDB Collections Design

## Overview
MongoDB will be used to store unstructured data, specifically high-volume social media scrapings and system logs that don't require strict relational integrity.

## Collections

### 1. `social_raw`
Stores raw data fetched from various social media platforms (Facebook, X, TikTok).
```json
{
  "_id": "ObjectId",
  "platform": "facebook", // facebook, x, tiktok
  "post_id": "string_unique_id_from_platform",
  "content": "Full text content of the post...",
  "author": {
    "name": "User Name",
    "profile_url": "..."
  },
  "metrics": {
    "likes": 100,
    "shares": 20,
    "comments": 5
  },
  "media_urls": ["url1", "url2"],
  "posted_at": "ISODate",
  "fetched_at": "ISODate",
  "sentiment_score": 0.5, // Added by NLP service later
  "keywords_matched": ["election", "policy"]
}
```

### 2. `social_trends`
Aggregated trends or daily summaries analyzed by AI.
```json
{
  "_id": "ObjectId",
  "date": "YYYY-MM-DD",
  "platform": "all",
  "top_keywords": [
    {"keyword": "tax", "count": 500},
    {"keyword": "road", "count": 300}
  ],
  "sentiment_summary": {
    "positive": 40,
    "neutral": 50,
    "negative": 10
  }
}
```

### 3. `system_logs`
General application logs for debugging and monitoring.
```json
{
  "_id": "ObjectId",
  "service": "backend-core",
  "level": "ERROR", // INFO, WARN, ERROR
  "message": "Connection timeout",
  "timestamp": "ISODate",
  "context": {
    "request_id": "...",
    "user_id": "..."
  }
}
```
