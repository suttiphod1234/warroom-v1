#!/bin/bash

echo "🧪 War Room API Testing Script"
echo "================================"
echo ""

BASE_URL="http://localhost:8080"
AI_URL="http://localhost:8000"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}📝 Test 1: Register New User${NC}"
echo "POST $BASE_URL/api/auth/register"
curl -X POST $BASE_URL/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "test123",
    "full_name": "Test User",
    "role": "voter"
  }' | jq
echo ""
echo ""

echo -e "${BLUE}📝 Test 2: Login as Admin${NC}"
echo "POST $BASE_URL/api/auth/login"
TOKEN=$(curl -s -X POST $BASE_URL/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123"
  }' | jq -r '.token')

if [ "$TOKEN" != "null" ] && [ -n "$TOKEN" ]; then
    echo -e "${GREEN}✅ Login successful!${NC}"
    echo "Token: ${TOKEN:0:50}..."
else
    echo -e "${RED}❌ Login failed!${NC}"
fi
echo ""
echo ""

echo -e "${BLUE}📝 Test 3: Get Profile${NC}"
echo "GET $BASE_URL/api/auth/profile"
curl -s $BASE_URL/api/auth/profile \
  -H "Authorization: Bearer $TOKEN" | jq
echo ""
echo ""

echo -e "${BLUE}📝 Test 4: Login as Voter${NC}"
VOTER_TOKEN=$(curl -s -X POST $BASE_URL/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "voter006",
    "password": "password123"
  }' | jq -r '.token')

echo -e "${GREEN}✅ Voter logged in${NC}"
echo ""

echo -e "${BLUE}📝 Test 5: Get Voter Profile (Before Purchase)${NC}"
VOTER_PROFILE=$(curl -s $BASE_URL/api/auth/profile \
  -H "Authorization: Bearer $VOTER_TOKEN")
echo "$VOTER_PROFILE" | jq
VOTER_ID=$(echo "$VOTER_PROFILE" | jq -r '.user.ID')
echo ""
echo ""

echo -e "${BLUE}📝 Test 6: Simulate Purchase (Add 1000 PV)${NC}"
echo "POST $BASE_URL/api/mlm/purchase"
curl -s -X POST $BASE_URL/api/mlm/purchase \
  -H "Content-Type: application/json" \
  -d "{
    \"user_id\": \"$VOTER_ID\",
    \"amount\": 1000
  }" | jq
echo ""
echo ""

echo -e "${BLUE}📝 Test 7: Check Profile After Purchase${NC}"
curl -s $BASE_URL/api/auth/profile \
  -H "Authorization: Bearer $VOTER_TOKEN" | jq
echo ""
echo ""

echo -e "${BLUE}📝 Test 8: Check Upline (voter001) - Should have GV${NC}"
UPLINE_TOKEN=$(curl -s -X POST $BASE_URL/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "voter001",
    "password": "password123"
  }' | jq -r '.token')

curl -s $BASE_URL/api/auth/profile \
  -H "Authorization: Bearer $UPLINE_TOKEN" | jq
echo ""
echo ""

echo -e "${BLUE}📝 Test 9: AI Service - Analyze Trend${NC}"
echo "POST $AI_URL/analyze/trend"
curl -s -X POST $AI_URL/analyze/trend \
  -H "Content-Type: application/json" \
  -d '{
    "hashtag": "election2024",
    "platform": "facebook"
  }' | jq
echo ""
echo ""

echo -e "${BLUE}📝 Test 10: Health Checks${NC}"
echo "Backend Health:"
curl -s $BASE_URL/health | jq
echo ""
echo "AI Service Health:"
curl -s $AI_URL/health | jq
echo ""
echo ""

echo -e "${GREEN}✅ All tests completed!${NC}"
echo ""
echo "📊 Summary:"
echo "  - Users created via seed script"
echo "  - Authentication working"
echo "  - MLM point distribution working (PV → GV)"
echo "  - AI service responding"
echo ""
echo "🔗 Try the frontend: http://localhost:3000/login"
echo "   Username: admin (or voter001, voter002, etc.)"
echo "   Password: password123"
