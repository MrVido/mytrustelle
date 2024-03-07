import React from 'react';

interface User {
  // Interface for user data (adapt based on your data structure)
  name: string;
  loyaltyTier: string; // Loyalty tier level (e.g., "bronze", "silver", "gold")
}

interface LoyaltyTierStatusProps {
  user: User;
}

const LoyaltyTierStatus: React.FC<LoyaltyTierStatusProps> = ({ user }) => {
  const getTierIcon = (tier: string) => {
    // Implement logic to return an icon based on the loyalty tier (replace with your icons)
    switch (tier) {
      case 'bronze':
        return <span className="bronze-tier"></span>;
      case 'silver':
        return <span className="silver-tier"></span>;
      case 'gold':
        return <span className="gold-tier"></span>;
      default:
        return null;
    }
  };

  return (
    <div className="loyalty-tier-status">
      {getTierIcon(user.loyaltyTier)}
      <p>Welcome, {user.name}! You are a {user.loyaltyTier} member.</p>
    </div>
  );
};

export default LoyaltyTierStatus;
