# 🚂 Deploy Backend to Railway - Quick Guide

## ขั้นตอนการ Deploy Backend (Go) บน Railway

### 1. เตรียม Railway Account
1. ไปที่ [https://railway.app](https://railway.app)
2. คลิก "Login" → "Login with GitHub"
3. อนุญาตให้ Railway เข้าถึง GitHub

### 2. สร้าง Project ใหม่
1. คลิก **"New Project"**
2. เลือก **"Deploy from GitHub repo"**
3. เลือก repository: **`suttiphod1234/warroom-v1`**
4. Railway จะสแกนหา Dockerfile

### 3. ตั้งค่า Backend Service

**Service Name:** `warroom-backend`

**Root Directory:** `backend`

**Environment Variables:**
```bash
PORT=8080
JWT_SECRET=your_super_secret_jwt_key_change_this_in_production
DB_HOST=${{Postgres.PGHOST}}
DB_PORT=${{Postgres.PGPORT}}
DB_USER=${{Postgres.PGUSER}}
DB_PASS=${{Postgres.PGPASSWORD}}
DB_NAME=${{Postgres.PGDATABASE}}
```

### 4. เพิ่ม PostgreSQL Database

1. ใน Railway Project คลิก **"New"**
2. เลือก **"Database"** → **"Add PostgreSQL"**
3. Railway จะสร้าง database และ set environment variables อัตโนมัติ

### 5. Deploy!

1. Railway จะ detect `Dockerfile` ใน `backend/` folder
2. คลิก **"Deploy"**
3. รอ 2-3 นาที
4. เสร็จแล้ว! Backend จะได้ URL เช่น: `https://warroom-backend-production.up.railway.app`

### 6. ทดสอบ Backend

```bash
# Test health endpoint
curl https://your-backend-url.railway.app/health

# Test register
curl -X POST https://your-backend-url.railway.app/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test123","full_name":"Test User","role":"voter"}'
```

---

## 🐍 Deploy AI Service to Railway

### 1. เพิ่ม Service ใหม่

1. ใน Railway Project คลิก **"New"**
2. เลือก **"GitHub Repo"** → เลือก `warroom-v1`
3. **Root Directory:** `ai-service`

### 2. เพิ่ม MongoDB (ใช้ MongoDB Atlas)

1. ไปที่ [https://www.mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
2. สร้าง Free Cluster (M0 - 512MB)
3. สร้าง Database User
4. เพิ่ม IP Address: `0.0.0.0/0` (allow all)
5. คัดลอก Connection String

### 3. ตั้งค่า AI Service

**Environment Variables:**
```bash
PORT=8000
MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/warroom
GOOGLE_API_KEY=your_gemini_api_key_here
```

### 4. Deploy AI Service

1. Railway จะ detect `Dockerfile` ใน `ai-service/`
2. คลิก **"Deploy"**
3. รอ 2-3 นาที
4. เสร็จแล้ว! AI Service จะได้ URL เช่น: `https://warroom-ai-production.up.railway.app`

---

## 🔗 เชื่อมต่อ Frontend กับ Backend

### 1. อัพเดท Vercel Environment Variables

1. ไปที่ Vercel Dashboard
2. เลือก Project "warroom-v1"
3. ไปที่ **Settings** → **Environment Variables**
4. เพิ่ม/แก้ไข:

```bash
NEXT_PUBLIC_BACKEND_URL=https://warroom-backend-production.up.railway.app
NEXT_PUBLIC_AI_URL=https://warroom-ai-production.up.railway.app
```

### 2. Redeploy Frontend

1. ไปที่ **Deployments** tab
2. คลิก **"..."** ข้าง deployment ล่าสุด
3. เลือก **"Redeploy"**
4. เสร็จแล้ว! ระบบพร้อมใช้งาน 🎉

---

## ✅ Checklist

- [ ] Railway account พร้อม
- [ ] Deploy Backend service
- [ ] เพิ่ม PostgreSQL database
- [ ] Deploy AI Service
- [ ] สร้าง MongoDB Atlas cluster
- [ ] ตั้งค่า Environment Variables
- [ ] เชื่อมต่อ Frontend กับ Backend
- [ ] ทดสอบ API endpoints
- [ ] ทดสอบ Login/Register
- [ ] ทดสอบ MLM purchase
- [ ] ทดสอบ AI analysis

---

## 💰 ค่าใช้จ่าย

### Railway (Backend + Database)
- **Free Trial**: $5 credit
- **Hobby Plan**: $5/month (ถ้าใช้เกิน free credit)
- **PostgreSQL**: รวมใน plan

### MongoDB Atlas
- **Free Tier (M0)**: 512MB - ฟรีตลอดไป
- **Shared Cluster (M2)**: $9/month (ถ้าต้องการมากกว่า)

### Vercel (Frontend)
- **Hobby**: ฟรี
- **Pro**: $20/month (ถ้าต้องการ custom domain + analytics)

**รวมทั้งหมด:** $0-5/เดือน (ใช้ free tiers)

---

## 🆘 Troubleshooting

### Backend ไม่ start
- ตรวจสอบ Dockerfile ใน `backend/`
- ดู Railway Logs
- ตรวจสอบว่า PORT environment variable ถูกต้อง

### Database connection failed
- ตรวจสอบว่า PostgreSQL service running
- ตรวจสอบ environment variables
- ดู Railway Logs สำหรับ connection errors

### Frontend ไม่เชื่อมต่อ Backend
- ตรวจสอบ CORS settings ใน Backend
- ตรวจสอบ `NEXT_PUBLIC_BACKEND_URL` ใน Vercel
- ดู Browser Console สำหรับ errors

---

**พร้อม Deploy แล้ว!** 🚀 เริ่มที่ [Railway.app](https://railway.app)
