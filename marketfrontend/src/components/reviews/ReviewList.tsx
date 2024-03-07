import React from 'react';

interface Review {
  id: string;
  reviewerName: string;
  text: string;
  rating: number;
  date: Date;
}

interface ReviewListProps {
  reviews: Review[];
}

const ReviewList: React.FC<ReviewListProps> = ({ reviews }) => {
  return (
    <div className="review-list">
      {reviews.map((review) => (
        <div key={review.id} className="review">
          <div className="review-header">
            <span className="reviewer-name">{review.reviewerName}</span>
            <span className="review-date">{review.date.toLocaleDateString()}</span>
          </div>
          <div className="review-rating">
            {/* Implement rating display, e.g., stars */}
          </div>
          <div className="review-text">{review.text}</div>
        </div>
      ))}
    </div>
  );
};

export default ReviewList;
