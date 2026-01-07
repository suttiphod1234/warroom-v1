# 🚀 Quick Deploy to Vercel (GitHub Integration)

## ขั้นตอนการ Deploy Frontend ไปยัง Vercel

### 1. เตรียม Vercel Account
1. ไปที่ [https://vercel.com](https://vercel.com)
2. คลิก "Sign Up" หรือ "Login"
3. เลือก "Continue with GitHub"
4. อนุญาตให้ Vercel เข้าถึง GitHub account ของคุณ

### 2. Import Project จาก GitHub

1. ใน Vercel Dashboard คลิก **"Add New..."** → **"Project"**
2. คลิก **"Import Git Repository"**
3. เลือก repository: **`suttiphod1234/warroom-v1`**
4. คลิก **"Import"**

### 3. Configure Project Settings

**Framework Preset:** Next.js (Vercel จะตรวจจับอัตโนมัติ)

**Root Directory:** 
- คลิก "Edit" ข้าง Root Directory
- เลือก **`frontend`**

**Build Settings:**
- Build Command: `npm run build` (ค่าเริ่มต้น)
- Output Directory: `.next` (ค่าเริ่มต้น)
- Install Command: `npm install` (ค่าเริ่มต้น)

### 4. Environment Variables (สำคัญ!)

คลิก **"Environment Variables"** และเพิ่ม:

```
NEXT_PUBLIC_BACKEND_URL=https://your-backend-url.com
NEXT_PUBLIC_AI_URL=https://your-ai-service-url.com
```

**หมายเหตุ:** ตอนนี้ใส่ค่าชั่วคราวได้ เช่น:
```
NEXT_PUBLIC_BACKEND_URL=http://localhost:8080
NEXT_PUBLIC_AI_URL=http://localhost:8000
```

### 5. Deploy!

1. คลิก **"Deploy"**
2. รอ 2-3 นาที (Vercel จะ build และ deploy)
3. เมื่อเสร็จจะได้ URL เช่น: `https://warroom-v1.vercel.app`

---

## 🎉 ผลลัพธ์

เมื่อ Deploy สำเร็จ คุณจะได้:

✅ **Production URL**: `https://warroom-v1-xxx.vercel.app`
✅ **Auto-deploy**: ทุกครั้งที่ push ไป GitHub, Vercel จะ deploy อัตโนมัติ
✅ **Preview Deployments**: ทุก Pull Request จะได้ preview URL
✅ **HTTPS**: มี SSL Certificate ฟรี
✅ **CDN**: เว็บไซต์เร็วทั่วโลก

---

## 📱 ทดสอบเว็บไซต์

หลัง Deploy เสร็จ ให้ทดสอบ:

1. เปิด URL ที่ได้: `https://warroom-v1-xxx.vercel.app`
2. ควรเห็นหน้า Homepage ของ War Room
3. คลิก "เข้าสู่ระบบ" → ควรเห็นหน้า Login

**หมายเหตุ:** 
- ตอนนี้จะยัง Login ไม่ได้ เพราะ Backend ยังไม่ได้ deploy
- หน้าเว็บจะแสดงได้ แต่ฟังก์ชันที่ต้องใช้ Backend จะยังใช้ไม่ได้

---

## 🔧 Deploy Backend (ขั้นตอนถัดไป)

หลังจาก Frontend deploy แล้ว ให้ deploy Backend:

### Option 1: Railway (แนะนำ)
1. ไปที่ [https://railway.app](https://railway.app)
2. Sign in with GitHub
3. New Project → Deploy from GitHub
4. เลือก `warroom-v1`
5. เลือก `backend` directory
6. เพิ่ม PostgreSQL database
7. Deploy!

### Option 2: Render
1. ไปที่ [https://render.com](https://render.com)
2. New → Web Service
3. Connect GitHub → เลือก `warroom-v1`
4. Root Directory: `backend`
5. Build Command: `go build -o main .`
6. Start Command: `./main`
7. Deploy!

---

## 🔗 เชื่อมต่อ Frontend กับ Backend

เมื่อ Backend deploy เสร็จ (ได้ URL เช่น `https://warroom-backend.railway.app`):

1. กลับไปที่ Vercel Dashboard
2. เลือก Project "warroom-v1"
3. Settings → Environment Variables
4. แก้ไข `NEXT_PUBLIC_BACKEND_URL` เป็น URL จริงของ Backend
5. Redeploy (Deployments → คลิก ... → Redeploy)

---

## ✅ Checklist

- [ ] Vercel account พร้อม
- [ ] Import project จาก GitHub
- [ ] ตั้งค่า Root Directory = `frontend`
- [ ] เพิ่ม Environment Variables
- [ ] Deploy สำเร็จ
- [ ] ทดสอบเว็บไซต์ที่ Production URL
- [ ] (Optional) Deploy Backend
- [ ] (Optional) เชื่อมต่อ Frontend กับ Backend

---

## 🆘 Troubleshooting

### Build Failed
- ตรวจสอบว่า Root Directory ตั้งเป็น `frontend`
- ดู Build Logs ใน Vercel
- ตรวจสอบว่า `package.json` ถูกต้อง

### Page Not Found (404)
- ตรวจสอบว่า Output Directory เป็น `.next`
- ตรวจสอบว่า Framework Preset เป็น Next.js

### Environment Variables ไม่ทำงาน
- ตรวจสอบว่าขึ้นต้นด้วย `NEXT_PUBLIC_`
- Redeploy หลังจากเพิ่ม Environment Variables

---

## 📞 Support

หากมีปัญหา:
1. ดู Vercel Build Logs
2. ตรวจสอบ `DEPLOYMENT.md` สำหรับรายละเอียดเพิ่มเติม
3. ดู [Vercel Documentation](https://vercel.com/docs)

---

**พร้อม Deploy แล้ว!** 🚀 เริ่มต้นที่ [Vercel.com](https://vercel.com) เลย!
