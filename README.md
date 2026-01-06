# 🗳️ War Room V1 - Political MLM & AI Monitoring System

Welcome to the **War Room**, a comprehensive system designed for political strategy, member management (MLM style), and AI-driven social sentiment analysis.

## 🏗️ Architecture
The system is built as a Monorepo with the following microservices:

| Service | Tech Stack | Description |
|---------|------------|-------------|
| **Backend** | Go (Fiber), GORM | Core logic, Auth, MLM Engine, Database Management. |
| **AI Service** | Python (FastAPI), Gemini | Social Media Scraper (Mock), Sentiment Analysis, Trends. |
| **Frontend** | Next.js 14, Tailwind | Web Dashboard, War Room UI, Genealogy. |
| **Mobile** | Flutter | Member App for check-ins and wallet viewing. |
| **Database** | PostgreSQL, MongoDB, Redis | Hybrid storage for relational and unstructured data. |

## 🚀 Getting Started

### Prerequisites
- Docker & Docker Compose (Recommended)
- OR manually install:
  - Node.js 18+
  - Go 1.21+
  - Python 3.10+
  - Flutter SDK 3.16+

### Quick Start with Docker (Recommended)

1. **Clone the repository**
```bash
git clone https://github.com/suttiphod1234/warroom-v1.git
cd warroom-v1
```

2. **Set up environment variables**
```bash
cp .env.example .env
# Edit .env with your settings (optional for local dev)
```

3. **Start all services**
```bash
docker-compose up -d --build
```

This will start:
- PostgreSQL on port 5432
- MongoDB on port 27017
- Redis on port 6379
- Backend API on port 8080
- AI Service on port 8000

4. **Access the services**
- Backend API: http://localhost:8080
- AI Service: http://localhost:8000
- API Docs (AI): http://localhost:8000/docs

### Manual Development Setup

#### Backend (Go)
```bash
cd backend
go mod tidy
go run main.go
# API running at http://localhost:8080
```

#### AI Service (Python)
```bash
cd ai-service
pip install -r requirements.txt
uvicorn main:app --reload
# API running at http://localhost:8000
```

#### Frontend (Next.js)
```bash
cd frontend
npm install
npm run dev
# Web App running at http://localhost:3000
```

#### Mobile (Flutter)
```bash
cd mobile
flutter pub get
flutter run
```

## 🔑 API Usage

### 1. Register a User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123",
    "full_name": "Admin User",
    "role": "admin"
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123"
  }'
```

### 3. Get Profile (requires token)
```bash
curl http://localhost:8080/api/auth/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Analyze Social Trends (AI Service)
```bash
curl -X POST http://localhost:8000/analyze/trend \
  -H "Content-Type: application/json" \
  -d '{
    "hashtag": "election2024",
    "platform": "all"
  }'
```

## 🧪 Testing

### Run Backend Tests
```bash
cd backend
go test ./...
```

### Run AI Service Tests
```bash
cd ai-service
pytest
```

### Run Frontend Build
```bash
cd frontend
npm run build
```

### Run Flutter Analysis
```bash
cd mobile
flutter analyze
```

## 📱 Features

### Web Dashboard
- **Login Page**: `/login` - Secure authentication
- **Dashboard**: `/dashboard` - View PV, GV, Commission, Tier Level
- **War Room**: `/warroom` - AI-powered social monitoring
- **Meeting Room**: `/meeting` - Jitsi video conferencing

### Mobile App
- Login with JWT authentication
- View wallet (PV/GV/Commission)
- Join secret meetings via Jitsi

### Backend APIs
- `/api/auth/register` - User registration
- `/api/auth/login` - User login
- `/api/auth/profile` - Get user profile
- `/api/mlm/purchase` - Simulate purchase (adds PV, distributes GV)

### AI Service APIs
- `/analyze/trend` - Analyze social media trends
- `/health` - Health check

## 🐛 Troubleshooting

### Docker Issues
```bash
# Stop all containers
docker-compose down

# Remove volumes and restart
docker-compose down -v
docker-compose up -d --build
```

### Database Connection Issues
- Check if PostgreSQL is running: `docker ps`
- Verify credentials in `.env` or `docker-compose.yml`
- Check logs: `docker-compose logs postgres`

### Port Already in Use
```bash
# Find process using port 8080
lsof -i :8080
# Kill the process
kill -9 <PID>
```

## 📝 Project Structure
```
warroom-v1/
├── backend/          # Go backend service
├── ai-service/       # Python AI service
├── frontend/         # Next.js web app
├── mobile/           # Flutter mobile app
├── docker/           # Docker configurations
├── .github/          # GitHub Actions CI/CD
└── docker-compose.yml
```

## 🤝 Contributing
This is a private project. For access, contact the repository owner.

## 📄 License
Proprietary - All rights reserved
