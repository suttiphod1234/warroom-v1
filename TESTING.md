# 🧪 Testing Guide - War Room V1

## Quick Start Testing

### 1. Start the System
```bash
# Start all services with Docker
docker-compose up -d --build

# Wait for services to be ready (about 30 seconds)
docker-compose logs -f
```

### 2. Seed Test Data
```bash
# Run the seed script to create test users
cd backend
go run cmd/seed/main.go
```

This creates the following test accounts:

| Username | Password | Role | Description |
|----------|----------|------|-------------|
| admin | password123 | Admin | System administrator |
| mp001 | password123 | MP | Member of Parliament |
| agent001 | password123 | Agent | Network head under MP |
| agent002 | password123 | Agent | Network head under MP |
| voter001-006 | password123 | Voter | Regular members |

**MLM Hierarchy:**
```
MP (mp001)
  ├─ Agent 1 (agent001)
  │   ├─ Voter 1 (voter001)
  │   │   └─ Voter 6 (voter006) ← Deepest level
  │   ├─ Voter 2 (voter002)
  │   └─ Voter 3 (voter003)
  └─ Agent 2 (agent002)
      ├─ Voter 4 (voter004)
      └─ Voter 5 (voter005)
```

### 3. Run Automated Tests
```bash
# From project root
./test-api.sh
```

This script will:
- ✅ Test user registration
- ✅ Test login/authentication
- ✅ Test profile retrieval
- ✅ Test MLM purchase (PV distribution)
- ✅ Verify GV flows to upline
- ✅ Test AI trend analysis
- ✅ Check all health endpoints

## Manual Testing Scenarios

### Scenario 1: Test MLM Point Distribution

**Goal:** Verify that when a user makes a purchase, points flow up the hierarchy.

```bash
# 1. Login as voter006 (deepest level)
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "voter006", "password": "password123"}'
# Save the token

# 2. Make a purchase (1000 PV)
curl -X POST http://localhost:8080/api/mlm/purchase \
  -H "Content-Type: application/json" \
  -d '{"user_id": "VOTER006_ID", "amount": 1000}'

# 3. Check voter006 wallet
# Expected: PVBalance = 1000

# 4. Check voter001 wallet (upline)
# Expected: GVBalance = 1000 (received from downline)

# 5. Check agent001 wallet (2 levels up)
# Expected: GVBalance = 1000 (cascaded up)
```

### Scenario 2: Test Tier Promotion

```bash
# Make multiple purchases to trigger tier upgrade
# Tier 1 (Bronze): Default
# Tier 2 (Silver): PV >= 1000
# Tier 3 (Gold): PV >= 5000 OR GV >= 20000

# Purchase 5000 PV
curl -X POST http://localhost:8080/api/mlm/purchase \
  -H "Content-Type: application/json" \
  -d '{"user_id": "USER_ID", "amount": 5000}'

# Check profile - TierLevel should be 3
```

### Scenario 3: Test Frontend Login

1. Open browser: http://localhost:3000/login
2. Login with: `admin` / `password123`
3. You should see the Command Center dashboard
4. Check the stats:
   - PV Balance
   - GV Balance
   - Commission
   - Tier Level

### Scenario 4: Test War Room (AI Analysis)

1. Navigate to: http://localhost:3000/warroom
2. Enter hashtag: `election2024`
3. Click "INITIATE SCAN"
4. You should see:
   - AI-generated summary
   - Sample social media posts
   - Sentiment analysis

### Scenario 5: Test Video Meeting

1. Go to: http://localhost:3000/meeting
2. The Jitsi Meet interface should load
3. You can join the room: `WarRoom-Secret-HQ-001`
4. Test with multiple browser tabs to simulate multiple users

## API Endpoint Reference

### Authentication
```bash
# Register
POST /api/auth/register
Body: {"username": "user", "password": "pass", "full_name": "Name", "role": "voter"}

# Login
POST /api/auth/login
Body: {"username": "user", "password": "pass"}
Returns: {"token": "JWT_TOKEN"}

# Get Profile (requires auth)
GET /api/auth/profile
Header: Authorization: Bearer JWT_TOKEN
```

### MLM Operations
```bash
# Make Purchase (triggers PV/GV distribution)
POST /api/mlm/purchase
Body: {"user_id": "UUID", "amount": 1000}
```

### AI Service
```bash
# Analyze Social Trend
POST http://localhost:8000/analyze/trend
Body: {"hashtag": "keyword", "platform": "facebook"}
```

## Verification Checklist

- [ ] All services start successfully
- [ ] Seed data creates 10+ users
- [ ] Can login with test accounts
- [ ] MLM purchase adds PV to buyer
- [ ] MLM purchase adds GV to upline (recursive)
- [ ] Tier promotion works at thresholds
- [ ] Frontend login page loads
- [ ] Dashboard shows correct stats
- [ ] War Room AI analysis works
- [ ] Meeting room loads Jitsi
- [ ] Mobile app can login (if testing Flutter)

## Troubleshooting

### "Connection refused" errors
```bash
# Check if services are running
docker-compose ps

# Check logs
docker-compose logs backend
docker-compose logs ai-service
```

### "User already exists" during seeding
```bash
# Reset database
docker-compose down -v
docker-compose up -d
# Wait 30 seconds, then run seed again
```

### Frontend can't connect to backend
- Check if backend is on port 8080: `curl http://localhost:8080/health`
- Check CORS settings in backend
- Check browser console for errors

## Performance Testing

### Load Test (optional)
```bash
# Install Apache Bench
brew install httpd  # macOS

# Test login endpoint (100 requests, 10 concurrent)
ab -n 100 -c 10 -p login.json -T application/json \
  http://localhost:8080/api/auth/login
```

## Expected Results

After running all tests, you should see:
- ✅ 10+ users in database
- ✅ Hierarchical relationships established
- ✅ Wallets created for all users
- ✅ PV/GV calculations working
- ✅ Tier promotions triggering correctly
- ✅ AI service responding
- ✅ Frontend accessible and functional

🎉 **If all checks pass, your War Room system is fully operational!**
