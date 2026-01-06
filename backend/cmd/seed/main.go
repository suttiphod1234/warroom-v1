package main

import (
	"fmt"
	"log"
	"warroom-backend/db"
	"warroom-backend/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Connect to database
	db.Connect()
	
	// Run migrations
	db.Migrate()

	fmt.Println("🌱 Starting database seeding...")

	// Create Admin User
	adminID := createUser("admin", "password123", "Admin User", models.RoleAdmin, nil, nil)
	fmt.Println("✅ Created Admin:", adminID)

	// Create MP (Member of Parliament)
	mpID := createUser("mp001", "password123", "MP Somchai", models.RoleMP, nil, nil)
	fmt.Println("✅ Created MP:", mpID)

	// Create Agent (Head of network) under MP
	agent1ID := createUser("agent001", "password123", "Agent Somsak", models.RoleAgent, &mpID, &mpID)
	fmt.Println("✅ Created Agent 1:", agent1ID)

	agent2ID := createUser("agent002", "password123", "Agent Somying", models.RoleAgent, &mpID, &mpID)
	fmt.Println("✅ Created Agent 2:", agent2ID)

	// Create Voters (Members) under Agent 1
	voter1ID := createUser("voter001", "password123", "Voter Nok", models.RoleVoter, &agent1ID, &agent1ID)
	fmt.Println("✅ Created Voter 1:", voter1ID)

	voter2ID := createUser("voter002", "password123", "Voter Ploy", models.RoleVoter, &agent1ID, &agent1ID)
	fmt.Println("✅ Created Voter 2:", voter2ID)

	voter3ID := createUser("voter003", "password123", "Voter Mint", models.RoleVoter, &agent1ID, &agent1ID)
	fmt.Println("✅ Created Voter 3:", voter3ID)

	// Create Voters under Agent 2
	voter4ID := createUser("voter004", "password123", "Voter Pim", models.RoleVoter, &agent2ID, &agent2ID)
	fmt.Println("✅ Created Voter 4:", voter4ID)

	voter5ID := createUser("voter005", "password123", "Voter Fon", models.RoleVoter, &agent2ID, &agent2ID)
	fmt.Println("✅ Created Voter 5:", voter5ID)

	// Create sub-level under Voter 1
	subVoter1ID := createUser("voter006", "password123", "Voter Beam", models.RoleVoter, &voter1ID, &voter1ID)
	fmt.Println("✅ Created Sub-Voter 1:", subVoter1ID)

	fmt.Println("\n🎉 Database seeding completed!")
	fmt.Println("\n📊 Hierarchy Structure:")
	fmt.Println("Admin (admin)")
	fmt.Println("MP (mp001)")
	fmt.Println("  ├─ Agent 1 (agent001)")
	fmt.Println("  │   ├─ Voter 1 (voter001)")
	fmt.Println("  │   │   └─ Voter 6 (voter006)")
	fmt.Println("  │   ├─ Voter 2 (voter002)")
	fmt.Println("  │   └─ Voter 3 (voter003)")
	fmt.Println("  └─ Agent 2 (agent002)")
	fmt.Println("      ├─ Voter 4 (voter004)")
	fmt.Println("      └─ Voter 5 (voter005)")
	fmt.Println("\n🔑 All passwords: password123")
}

func createUser(username, password, fullName string, role models.UserRole, sponsorID, placementID *uuid.UUID) uuid.UUID {
	// Check if user exists
	var existing models.User
	if err := db.DB.Where("username = ?", username).First(&existing).Error; err == nil {
		fmt.Printf("⚠️  User %s already exists, skipping...\n", username)
		return existing.ID
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	// Create user
	user := models.User{
		Username:     username,
		PasswordHash: string(hash),
		FullName:     fullName,
		Role:         role,
		SponsorID:    sponsorID,
		PlacementID:  placementID,
		IsActive:     true,
		IsBanned:     false,
		TierLevel:    1,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Fatal("Failed to create user:", err)
	}

	// Create wallet for user
	wallet := models.Wallet{
		UserID:            user.ID,
		PVBalance:         0,
		GVBalance:         0,
		CommissionBalance: 0,
	}

	if err := db.DB.Create(&wallet).Error; err != nil {
		log.Fatal("Failed to create wallet:", err)
	}

	return user.ID
}
