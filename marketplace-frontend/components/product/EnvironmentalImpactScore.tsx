import React from 'react';

interface EnvironmentalImpactScoreProps {
  score: number; // Assume a score out of 100
  maxScore?: number;
}

const EnvironmentalImpactScore: React.FC<EnvironmentalImpactScoreProps> = ({ score, maxScore = 100 }) => {
  const scorePercentage = (score / maxScore) * 100;
  return (
    <div className="environmental-impact-score">
      <label>Environmental Impact Score:</label>
      <progress value={score} max={maxScore}></progress>
      <span>{scorePercentage}%</span>
    </div>
  );
};

export default EnvironmentalImpactScore;
