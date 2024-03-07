import React, { useState, useEffect } from 'react';

interface SurveyFormProps {
  title?: string;
  questions: SurveyQuestion[];
  onSubmit?: (answers: SurveyAnswer[]) => void;
}

interface SurveyQuestion {
  type: 'text' | 'radio' | 'checkbox' | 'select';
  label: string;
  options?: string[];
  required?: boolean;
}

interface SurveyAnswer {
  questionId: number;
  answer: string | string[];
}

const SurveyForm: React.FC<SurveyFormProps> = ({ title = '', questions, onSubmit }) => {
  const [answers, setAnswers] = useState<SurveyAnswer[]>([]);

  // Initialize answers based on questions
  useEffect(() => {
    const initialAnswers = questions.map((question, index) => ({
      questionId: index, // Assuming index as questionId for simplicity
      answer: question.type === 'checkbox' ? [] : '', // Initialize checkboxes with empty array, others with empty string
    }));
    setAnswers(initialAnswers);
  }, [questions]);

  const handleAnswerChange = (questionId: number, newAnswer: string | string[]) => {
    setAnswers((prevAnswers) =>
      prevAnswers.map((answer) =>
        answer.questionId === questionId ? { ...answer, answer: newAnswer } : answer
      )
    );
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    onSubmit && onSubmit(answers);
  };

  return (
    <form className="survey-form" onSubmit={handleSubmit}>
      {title && <h2>{title}</h2>}
      {questions.map((question, index) => (
        <div key={index} className="survey-question">
          <label>{question.label}</label>
          {question.type === 'text' && (
            <input
              type="text"
              onChange={(e) => handleAnswerChange(index, e.target.value)}
              required={question.required}
            />
          )}
          {question.type === 'radio' &&
            question.options &&
            question.options.map((option, optionIndex) => (
              <label key={optionIndex}>
                <input
                  type="radio"
                  name={`question-${index}`}
                  value={option}
                  onChange={() => handleAnswerChange(index, option)}
                  required={question.required}
                />
                {option}
              </label>
            ))}
          {/* Implement checkbox and select handling similar to the text and radio */}
        </div>
      ))}
      <button type="submit">Submit Survey</button>
    </form>
  );
};

export default SurveyForm;
