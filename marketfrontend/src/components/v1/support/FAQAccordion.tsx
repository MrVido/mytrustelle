import React, { useState } from 'react';

interface FAQ {
  question: string;
  answer: string;
}

interface FAQAccordionProps {
  faqs: FAQ[];
}

const FAQAccordion: React.FC<FAQAccordionProps> = ({ faqs }) => {
  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  const toggleAccordion = (index: number) => {
    setActiveIndex(activeIndex === index ? null : index);
  };

  return (
    <div className="faq-accordion">
      {faqs.map((faq, index) => (
        <div key={index} className="faq-item">
          <button className="faq-question" onClick={() => toggleAccordion(index)}>
            {faq.question}
          </button>
          <div
            className={`faq-answer ${activeIndex === index ? 'active' : ''}`}
            aria-expanded={activeIndex === index}
          >
            {faq.answer}
          </div>
        </div>
      ))}
    </div>
  );
};

export default FAQAccordion;
