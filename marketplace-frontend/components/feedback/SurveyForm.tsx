import React, { useState, useEffect } from 'react';

interface SurveyFormProps {
  // Props to define survey content and functionality
  title?: string; // Optional title for the survey
  questions: SurveyQuestion[]; // Array of questions for the survey
  onSubmit?: (answers: SurveyAnswer[]) => void; // Function to handle form submission
}

interface SurveyQuestion {
  // Interface to define the structure of a survey question
  type: 'text' | 'radio' | 'checkbox' | 'select'; // Question type (e.g., text input, multiple choice)
  label: string; // Question text
  options?: string[]; // Optional options for multiple choice questions
  required?: boolean; // Optional flag indicating if the question is required
}

interface SurveyAnswer {
  // Interface to define the structure of a survey answer
  questionId: number; // Unique identifier for the question
  answer: string | string[]; // Answer value (text for text input, array of selected options for multiple choice)
}

const SurveyForm: React.FC<SurveyFormProps> = ({ title = '', questions, onSubmit }) => {
  const [answers, setAnswers] = useState<SurveyAnswer[]>([]);

  const handleAnswerChange = (questionId: number, newAnswer: string | string[]) => {
    setAnswers((prevAnswers) =>
      prevAnswers.map((answer) => (answer.questionId === questionId ? { ...answer, answer: newAnswer } : answer))
    );
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (onSubmit) {
      onSubmit(answers); // Call onSubmit function with collected answers
    } else {
      console.log('Submitted survey:', answers); // Placeholder if no onSubmit provided
    }
  };

  return (
    <form className="survey-form" onSubmit={handleSubmit}>
      {title && <h2>{title}</h2>}
      {questions.map((question, index) => (
        <div key={index} className="survey-question">
          <label htmlFor={`question-${index}`}>{question.label}</label>
          {question.type === 'text' && (
            <input
              id={`question-${index}`}
              type="text"
              onChange={(e) => handleAnswerChange(question.questionId, e.target.value)}
              required={question.required}
            />
          )}
          {question.type === 'radio' && question.options && (
            <div className="radio-options">
              {question.options.map((option, optionIndex) => (
                <label key={optionIndex}>
                  <input
                    type="radio"
                    id={`question-${index}-${optionIndex}`} // Unique id for each radio button
                    name={`question-${index}`} // Ensures radio buttons within a question act as a group
                    value={option}
                    onChange={() => handleAnswerChange(question.questionId, option)}
                    required={question.required}
                  />{' '}
                  {option}
                </label>
              ))}
            </div>
          )}
          {/* ... rest of the code for checkbox and select remains the same */}
        </div>
      ))}
      <button type="submit">Submit Survey</button>
    </form>
  );
};

export default SurveyForm;
