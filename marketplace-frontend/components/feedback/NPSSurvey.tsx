import React, { useState, useEffect } from 'react';

interface NPSSurveyProps {
  // Optional props for customization (e.g., initial score, question text)
}

const NPSSurvey: React.FC<NPSSurveyProps> = ({ initialScore = 0 } = {}) => {
  const [score, setScore] = useState(initialScore);
  const [feedback, setFeedback] = useState('');

  const handleScoreChange = (newScore: number) => {
    setScore(newScore);
  };

  const handleFeedbackChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setFeedback(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Handle form submission logic here (e.g., send data to server)
    console.log('Submitted NPS survey:', { score, feedback }); // Placeholder for submission logic
  };

  return (
    <form className="nps-survey" onSubmit={handleSubmit}>
      <h2>NPS Survey</h2>
      <p>On a scale of 0-10, how likely are you to recommend us to a friend or colleague?</p>
      <div className="score-selection">
        {[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((n) => (
          <button key={n} type="button" onClick={() => handleScoreChange(n)} className={n === score ? 'active' : ''}>
            {n}
          </button>
        ))}
      </div>
      <label htmlFor="feedback">Optional: Provide any feedback you have for us</label>
      <textarea id="feedback" name="feedback" value={feedback} onChange={handleFeedbackChange} />
      <button type="submit">Submit</button>
    </form>
  );
};

export default NPSSurvey;
