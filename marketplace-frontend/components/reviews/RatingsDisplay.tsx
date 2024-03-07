import React from 'react';

interface RatingsDisplayProps {
  rating: number; // The average rating value
  reviewCount: number; // The total number of reviews
}

const RatingsDisplay: React.FC<RatingsDisplayProps> = ({ rating, reviewCount }) => {
  return (
    <div className="ratings-display">
      <span className="stars">{/* Render star icons based on rating */}</span>
      <span className="rating-value">{rating.toFixed(1)}</span>
      <span className="review-count">({reviewCount} reviews)</span>
    </div>
  );
};

export default RatingsDisplay;
