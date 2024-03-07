import React, { useState } from 'react';

interface UserFeedbackProps {
  // Props to control modal visibility and handle submissions
  isOpen?: boolean; // Flag indicating if the modal is open.
  onClose?: () => void; // Function to handle modal closure.
  onSubmit?: (feedback: UserFeedback) => void; // Function to handle feedback submission.
}

interface UserFeedback {
  // Structure for user feedback data
  message?: string; // Detailed feedback message (from text input)
  rating?: number; // User's rating (e.g., 1-5 stars)
  category?: string; // Feedback category (e.g., bug report, suggestion)
}

const UserFeedbackModal: React.FC<UserFeedbackProps> = ({ isOpen = false, onClose, onSubmit }) => {
  const [feedback, setFeedback] = useState<UserFeedback>({}); // State for user feedback data

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFeedback({ ...feedback, [event.target.name]: event.target.value });
  };

  const handleRatingChange = (newRating: number) => {
    setFeedback({ ...feedback, rating: newRating });
  };

  const handleCategoryChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setFeedback({ ...feedback, category: event.target.value });
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (onSubmit) {
      onSubmit(feedback); // Call onSubmit function with collected feedback data
    }
    onClose?.(); // Close the modal after submission (if onClose provided)
  };

  return (
    <div className="user-feedback-modal" style={{ display: isOpen ? 'block' : 'none' }}>
      <h2>Provide Feedback</h2>
      <form onSubmit={handleSubmit}>
        <label htmlFor="message">Message:</label>
        <textarea id="message" name="message" onChange={handleInputChange} />
        {/* Rating component (implementation depends on chosen library) */}
        <label>Rating:</label>
        <div className="rating-options">
          {/* Render rating options (e.g., stars) and handle selection through handleRatingChange */}
        </div>
        <label htmlFor="category">Category:</label>
        <select id="category" name="category" onChange={handleCategoryChange}>
          <option value="">Select Category</option>
          <option value="bug">Bug Report</option>
          <option value="suggestion">Suggestion</option>
          {/* Add more options as needed */}
        </select>
        <button type="submit">Submit Feedback</button>
      </form>
      <button onClick={onClose}>Close</button>
    </div>
  );
};

export default UserFeedbackModal;
