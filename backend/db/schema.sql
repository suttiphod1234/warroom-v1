-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ROLES: Admin, MP (Member of Parliament/Candidate), Agent (Head), Voter (Member)
CREATE TYPE user_role AS ENUM ('admin', 'mp', 'agent', 'voter');

-- USERS Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    role user_role NOT NULL DEFAULT 'voter',
    phone_number VARCHAR(20),
    line_id VARCHAR(50),
    
    -- Hierarchy for MLM Tree
    sponsor_id UUID REFERENCES users(id), -- Who invited this user (Direct Sponsor)
    placement_id UUID REFERENCES users(id), -- Upliner in the binary/matrix tree
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_banned BOOLEAN DEFAULT FALSE,
    tier_level INT DEFAULT 1, -- 1=Bronze, 2=Silver, etc.
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- INDEX for Hierarchy Performance
CREATE INDEX idx_users_sponsor ON users(sponsor_id);
CREATE INDEX idx_users_placement ON users(placement_id);

-- WALLETS / POINTS (PV = Personal Volume, GV = Group Volume)
CREATE TABLE wallets (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    pv_balance DECIMAL(15, 2) DEFAULT 0.00, -- Cumulative Personal Points
    gv_balance DECIMAL(15, 2) DEFAULT 0.00, -- Group Points (Left/Right leg logic handled in app code)
    commission_balance DECIMAL(15, 2) DEFAULT 0.00, -- Withdrawable Cash
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- TRANSACTION HISTORY
CREATE TYPE transaction_type AS ENUM ('deposit', 'withdrawal', 'commission_payout', 'purchase_pv', 'transfer');

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    amount DECIMAL(15, 2) NOT NULL,
    transaction_type transaction_type NOT NULL,
    status VARCHAR(20) DEFAULT 'pending', -- pending, completed, failed
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- AUDIT LOG (For sensitive actions)
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    actor_id UUID REFERENCES users(id),
    action VARCHAR(50) NOT NULL,
    target_table VARCHAR(50),
    target_id UUID,
    details JSONB, -- Store previous values or specific changes
    ip_address VARCHAR(45),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
