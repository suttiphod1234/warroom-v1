# 🗳️ War Room V1 - Political MLM & AI Monitoring System

Welcome to the **War Room**, a comprehensive system designed for political strategy, member management (MLM style), and AI-driven social sentiment analysis.

## 🏗️ Architecture
The system is built as a Monorepo with the following microservices:

| Service | Tech Stack | Description |
|---------|------------|-------------|
| **Backend** | Go (Fiber), GORM | Core logic, Auth, MLM Engine, Database Management. |
| **AI Service** | Python (FastAPI), Gemini | Social Media Scraper (Mock), Sentiment Analysis, Trends. |
| **Frontend** | Next.js 14, Tailwind | Web Dashboard, War Room UI, Geneology. |
| **Mobile** | Flutter | Member App for check-ins and wallet viewing. |
| **Database** | PostgreSQL, MongoDB, Redis | Hybrid storage for relational and unstructured data. |

## 🚀 Getting Started

### Prerequisites
- Docker & Docker Compose
- Node.js 18+ (for Frontend dev)
- Go 1.21+ (for Backend dev)
- Python 3.10+ (for AI dev)
- Flutter SDK (for Mobile dev)

### 1-Click Start (Docker)
The easiest way to run the backend infrastructure is via Docker:

```bash
# Start Databases (Postgres, Mongo, Redis) and Core Services
docker-compose up -d --build
```

### Manual Development

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

## 🔑 Default Credentials
- **Admin User**: Created automatically on first run (check `backend/db/seed.go` if implemented, otherwise register via API).
- **Test Login**:
    - **POST** `/api/auth/register` First to create a user.
    - **POST** `/api/auth/login` To get a JWT Token.

## 🧪 Testing
- **Backend**: `cd backend && go test ./...`
- **AI**: `cd ai-service && pytest`
- **GitHub Actions**: Tests run automatically on push.
