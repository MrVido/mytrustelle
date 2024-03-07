import React, { useState } from 'react';

interface NPSSurveyProps {
  onSubmit: (score: number) => void; // Callback function to handle the submission
}

const NPSSurvey: React.FC<NPSSurveyProps> = ({ onSubmit }) => {
  const [score, setScore] = useState<number | null>(null);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (score !== null) {
      onSubmit(score);
      alert(`Thank you for your feedback! You rated us a ${score}/10.`);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="nps-survey">
      <h2>How likely are you to recommend us to a friend or colleague?</h2>
      <div className="nps-options">
        {Array.from({ length: 11 }, (_, i) => (
          <label key={i}>
            {i}
            <input
              type="radio"
              value={i}
              name="nps-score"
              onChange={() => setScore(i)}
              checked={score === i}
            />
          </label>
        ))}
      </div>
      <button type="submit" disabled={score === null}>
        Submit
      </button>
    </form>
  );
};

export default NPSSurvey;
