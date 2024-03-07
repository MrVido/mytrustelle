import React, { useState, useEffect } from 'react';
import LoyaltyTierStatus from './LoyaltyTierStatus'; // Import LoyaltyTierStatus component

interface Reward {
  id: string;
  name: string;
  description: string;
  pointsRequired: number;
  userPoints: number; // User's current points for this reward (optional)
  imageUrl?: string; // Optional image URL for the reward
}

interface User {
  name: string;
  loyaltyTier: string; // Loyalty tier level (e.g., "bronze", "silver", "gold")
  points: number; // User's total loyalty points
}

interface RewardsDashboardProps {
  user: User;
  fetchRewards: () => Promise<Reward[]>; // Updated to be mandatory and return a Promise of Reward array
}

const RewardsDashboard: React.FC<RewardsDashboardProps> = ({ user, fetchRewards }) => {
  const [isLoading, setIsLoading] = useState(false);
  const [rewards, setRewards] = useState<Reward[]>([]); // Define rewards as internal state
  const [selectedReward, setSelectedReward] = useState<Reward | null>(null);

  useEffect(() => {
    setIsLoading(true);
    fetchRewards()
      .then((data) => setRewards(data))
      .finally(() => setIsLoading(false));
  }, [fetchRewards]);

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
