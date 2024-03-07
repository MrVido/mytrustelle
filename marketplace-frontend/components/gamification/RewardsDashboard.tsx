import React, { useState, useEffect } from 'react';
import  LoyaltyTierStatus  from './LoyaltyTierStatus'; // Import LoyaltyTierStatus component

interface Reward {
  // Interface for reward data (adapt based on your data structure)
  id: string;
  name: string;
  description: string;
  pointsRequired: number;
  userPoints: number; // User's current points for this reward (optional)
  imageUrl?: string; // Optional image URL for the reward
}

interface User {
  // Interface for user data (adapt based on your data structure)
  name: string;
  loyaltyTier: string; // Loyalty tier level (e.g., "bronze", "silver", "gold")
  points: number; // User's total loyalty points
}

interface RewardsDashboardProps {
  user: User;
  rewards: Reward[];
  fetchRewards?: () => void; // Optional function to fetch rewards (if not provided upfront)
}

const RewardsDashboard: React.FC<RewardsDashboardProps> = ({ user, rewards, fetchRewards }) => {
  const [isLoading, setIsLoading] = useState(false); // State for loading indicator (optional)

  useEffect(() => {
    if (!rewards && fetchRewards) {
      setIsLoading(true);
      fetchRewards()
        .then((data) => setRewards(data))
        .finally(() => setIsLoading(false));
    }
  }, [rewards, fetchRewards]);

  const [selectedReward, setSelectedReward] = useState<Reward | null>(null);

  const handleRewardSelect = (reward: Reward) => {
    setSelectedReward(reward);
  };

  return (
    <div className="rewards-dashboard">
      {isLoading && <p>Loading rewards...</p>}
      <LoyaltyTierStatus user={user} />
      <h2>Your Rewards</h2>
      {rewards?.length > 0 ? (
        <ul className="rewards-list">
          {rewards.map((reward) => (
            <li key={reward.id}>
              <button onClick={() => handleRewardSelect(reward)}>
                {reward.imageUrl && <img src={reward.imageUrl} alt={reward.name} />}
                <span>{reward.name}</span>
              </button>
              {reward.userPoints !== undefined && (
                <p>
                  Points: {reward.userPoints}/{reward.pointsRequired}
                </p>
              )}
            </li>
          ))}
        </ul>
      ) : (
        <p>You don't have any rewards yet.</p>
      )}
      {selectedReward && (
        <div className="reward-details">
          <h3>{selectedReward.name}</h3>
          <p>{selectedReward.description}</p>
          {selectedReward.pointsRequired && (
            <p>Points required: {selectedReward.pointsRequired}</p>
          )}
          {/* Add button or link to claim reward (optional) */}
        </div>
      )}
    </div>
  );
};

export default RewardsDashboard;
