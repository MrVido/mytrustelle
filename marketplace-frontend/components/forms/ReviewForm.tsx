import React, { useState, useEffect } from 'react';
import { Rating } from './Rating'; // Import Rating component (replace with actual path)

interface Review {
  // Define the structure of a review object (adapt based on your data)
  id: string;
  userId: string; // User who submitted the review
  productId: string; // Product being reviewed
  rating: number;
  content: string;
  createdAt: string; // Date and time of review creation
}

interface ReviewFormProps {
  productId: string; // ID of the product being reviewed
  onSubmit: (review: Review) => void; // Function to handle successful review submission
}

const ReviewForm: React.FC<ReviewFormProps> = ({ productId, onSubmit }) => {
  const [rating, setRating] = useState(0);
  const [content, setContent] = useState('');

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const review = {
      productId,
      rating,
      content,
      // Add other properties like userId or createdAt if needed (replace with logic)
      userId: '123', // Replace with actual user ID retrieval logic
      createdAt: new Date().toISOString(),
    };
    onSubmit(review);
    setRating(0); // Clear form after submission (optional)
    setContent('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Write a Review</h2>
      <Rating rating={rating} onRateChange={setRating} />
      <label htmlFor="content">
        Review Content:
        <textarea id="content" name="content" value={content} onChange={(e) => setContent(e.target.value)} required />
      </label>
      <button type="submit">Submit Review</button>
    </form>
  );
};

export default ReviewForm;
