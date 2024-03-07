import React, { useState, useContext } from 'react';
import { IContext } from './Context'; // Assuming you have a context for language

const LanguageSelector: React.FC = () => {
  const { language, setLanguage } = useContext(IContext); // Access language and setLanguage from context

  const handleLanguageChange = (newLang: string) => {
    setLanguage(newLang); // Update language in context
  };

  return (
    <div className="language-selector">
      <select value={language} onChange={(e) => handleLanguageChange(e.target.value)}>
        <option value="en">English</option>
        <option value="fr">French</option>
      </select>
    </div>
  );
};

export default LanguageSelector;
