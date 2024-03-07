import React, { useState } from 'react';

const AccessibilityOptions: React.FC = () => {
  const [fontSize, setFontSize] = useState('normal');
  const [contrast, setContrast] = useState('normal');

  const handleFontSizeChange = (size: string) => {
    setFontSize(size);
    document.documentElement.style.setProperty('--user-font-size', size);
  };

  const handleContrastChange = (contrastMode: string) => {
    setContrast(contrastMode);
    document.documentElement.setAttribute('data-contrast', contrastMode);
  };

  return (
    <div className="accessibility-options">
      <h2>Accessibility Options</h2>
      <div>
        <label>Font Size:</label>
        <button onClick={() => handleFontSizeChange('small')}>Small</button>
        <button onClick={() => handleFontSizeChange('normal')}>Normal</button>
        <button onClick={() => handleFontSizeChange('large')}>Large</button>
      </div>
      <div>
        <label>Contrast:</label>
        <button onClick={() => handleContrastChange('normal')}>Normal</button>
        <button onClick={() => handleContrastChange('high')}>High Contrast</button>
      </div>
    </div>
  );
};

export default AccessibilityOptions;

