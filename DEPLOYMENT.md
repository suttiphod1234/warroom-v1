# 🚀 War Room V1 - Deployment Guide

## Overview
This guide will help you deploy the War Room system to production using free/affordable cloud services.

## Architecture
```
Frontend (Vercel) → Backend (Railway) → Database (Railway PostgreSQL + MongoDB Atlas)
                  ↓
              AI Service (Railway)
```

---

## Option 1: Quick Deploy (Recommended for Demo)

### 1. Deploy Frontend to Vercel

**Step 1:** Push your code to GitHub (already done ✅)

**Step 2:** Go to [Vercel](https://vercel.com)
- Sign in with GitHub
- Click "New Project"
- Import `suttiphod1234/warroom-v1`
- Framework Preset: **Next.js**
- Root Directory: `frontend`
- Click "Deploy"

**Step 3:** Set Environment Variables in Vercel
```
NEXT_PUBLIC_BACKEND_URL=https://your-backend.railway.app
NEXT_PUBLIC_AI_URL=https://your-ai-service.railway.app
```

**Result:** Your frontend will be live at `https://warroom-v1.vercel.app`

---

### 2. Deploy Backend to Railway

**Step 1:** Go to [Railway.app](https://railway.app)
- Sign in with GitHub
- Click "New Project" → "Deploy from GitHub repo"
- Select `suttiphod1234/warroom-v1`

**Step 2:** Add PostgreSQL Database
- Click "New" → "Database" → "Add PostgreSQL"
- Railway will automatically create a database

**Step 3:** Configure Backend Service
- Click "New" → "GitHub Repo" → Select your repo
- Root Directory: `backend`
- Add Environment Variables:
```bash
DB_HOST=${{Postgres.PGHOST}}
DB_USER=${{Postgres.PGUSER}}
DB_PASS=${{Postgres.PGPASSWORD}}
DB_NAME=${{Postgres.PGDATABASE}}
DB_PORT=${{Postgres.PGPORT}}
JWT_SECRET=your_super_secret_key_change_this
PORT=8080
```

**Step 4:** Add Start Command
- Settings → Deploy → Start Command: `./main`
- Settings → Deploy → Build Command: `go build -o main .`

**Result:** Backend will be live at `https://warroom-backend-xxx.railway.app`

---

### 3. Deploy AI Service to Railway

**Step 1:** In Railway, click "New" → "GitHub Repo"
- Root Directory: `ai-service`

**Step 2:** Add MongoDB Atlas (Free Tier)
- Go to [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
- Create free cluster
- Get connection string

**Step 3:** Configure AI Service
- Add Environment Variables:
```bash
MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/warroom
GOOGLE_API_KEY=your_gemini_api_key
PORT=8000
```

**Step 4:** Add Start Command
- Start Command: `uvicorn main:app --host 0.0.0.0 --port 8000`

**Result:** AI Service will be live at `https://warroom-ai-xxx.railway.app`

---

## Option 2: Deploy with Docker (Advanced)

### Using Railway with Docker

**Step 1:** Railway will auto-detect `Dockerfile` in each service

**Step 2:** Set environment variables as above

**Step 3:** Railway will build and deploy automatically

---

## Option 3: Self-Hosted (VPS)

### Requirements
- Ubuntu 22.04 VPS (DigitalOcean, Linode, etc.)
- Domain name (optional)

### Installation Steps

```bash
# 1. Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# 2. Clone repository
git clone https://github.com/suttiphod1234/warroom-v1.git
cd warroom-v1

# 3. Set environment variables
cp .env.example .env
nano .env  # Edit with your settings

# 4. Start services
docker-compose up -d --build

# 5. Check status
docker-compose ps
```

### Setup Nginx Reverse Proxy

```nginx
# /etc/nginx/sites-available/warroom
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:3000;  # Frontend
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }

    location /api {
        proxy_pass http://localhost:8080;  # Backend
    }

    location /ai {
        proxy_pass http://localhost:8000;  # AI Service
    }
}
```

```bash
# Enable site
sudo ln -s /etc/nginx/sites-available/warroom /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx

# Setup SSL with Let's Encrypt
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com
```

---

## Environment Variables Reference

### Frontend (.env.local)
```bash
NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
NEXT_PUBLIC_AI_URL=http://localhost:8000
```

### Backend (.env)
```bash
DB_HOST=localhost
DB_USER=admin
DB_PASS=password
DB_NAME=warroom_db
DB_PORT=5432
JWT_SECRET=super_secret_key
REDIS_HOST=localhost
REDIS_PORT=6379
PORT=8080
```

### AI Service (.env)
```bash
MONGO_URI=mongodb://admin:password@localhost:27017/warroom
GOOGLE_API_KEY=your_gemini_api_key
PORT=8000
```

---

## Post-Deployment Checklist

- [ ] Frontend loads correctly
- [ ] Can register new user
- [ ] Can login and get JWT token
- [ ] Dashboard displays user data
- [ ] MLM purchase works (PV/GV distribution)
- [ ] AI analysis endpoint responds
- [ ] Meeting room loads Jitsi
- [ ] Mobile app can connect to APIs

---

## Monitoring & Maintenance

### Check Logs
```bash
# Railway: Click on service → Logs tab

# Docker:
docker-compose logs -f backend
docker-compose logs -f ai-service
```

### Database Backup
```bash
# PostgreSQL
docker exec warroom_postgres pg_dump -U admin warroom_db > backup.sql

# MongoDB
docker exec warroom_mongo mongodump --out /backup
```

### Update Deployment
```bash
# Pull latest code
git pull origin main

# Rebuild and restart
docker-compose down
docker-compose up -d --build
```

---

## Cost Estimation

### Free Tier (Recommended for Testing)
- **Vercel**: Free (Frontend)
- **Railway**: $5/month credit (Backend + DB)
- **MongoDB Atlas**: Free 512MB (AI Service DB)
- **Total**: ~$0-5/month

### Production (Recommended)
- **Vercel Pro**: $20/month (Frontend)
- **Railway**: ~$20/month (Backend + PostgreSQL)
- **MongoDB Atlas**: $9/month (Shared cluster)
- **Total**: ~$49/month

### Self-Hosted VPS
- **DigitalOcean Droplet**: $6-12/month (2GB RAM)
- **Domain**: $10-15/year
- **Total**: ~$6-12/month

---

## Troubleshooting

### Frontend can't connect to Backend
- Check CORS settings in backend
- Verify `NEXT_PUBLIC_BACKEND_URL` is correct
- Check backend is running: `curl https://your-backend.railway.app/health`

### Database connection failed
- Verify database credentials
- Check if database service is running
- Test connection: `psql -h host -U user -d database`

### AI Service not responding
- Check if `GOOGLE_API_KEY` is set
- Verify MongoDB connection
- Check logs for errors

---

## Support

For deployment issues:
1. Check Railway/Vercel logs
2. Review `TESTING.md` for API testing
3. Open an issue on GitHub

---

## Quick Links

- **Vercel Dashboard**: https://vercel.com/dashboard
- **Railway Dashboard**: https://railway.app/dashboard
- **MongoDB Atlas**: https://cloud.mongodb.com
- **GitHub Repo**: https://github.com/suttiphod1234/warroom-v1

---

**Ready to deploy?** Start with Option 1 (Vercel + Railway) for the fastest setup! 🚀
