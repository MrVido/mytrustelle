import React, { useState, useEffect } from 'react';

interface Tutorial {
  id: string;
  title: string;
  description: string;
  accessUrl: string;
}

interface TutorialAccessProps {
  tutorials: Tutorial[]; // Array of tutorials available
}

const TutorialAccess: React.FC<TutorialAccessProps> = ({ tutorials }) => {
  const [selectedTutorialId, setSelectedTutorialId] = useState<string | null>(null);

  useEffect(() => {
    if (tutorials.length > 0) {
      // Automatically select the first tutorial by default
      setSelectedTutorialId(tutorials[0].id);
    }
  }, [tutorials]);

  const handleAccessTutorial = (tutorialId: string) => {
    const tutorial = tutorials.find((t) => t.id === tutorialId);
    if (tutorial) {
      window.open(tutorial.accessUrl, '_blank');
    }
  };

  return (
    <div className="tutorial-access">
      <h2>Access Tutorials</h2>
      {tutorials.length > 0 ? (
        <select
          value={selectedTutorialId || ''}
          onChange={(e) => setSelectedTutorialId(e.target.value)}
        >
          {tutorials.map((tutorial) => (
            <option key={tutorial.id} value={tutorial.id}>
              {tutorial.title}
            </option>
          ))}
        </select>
      ) : (
        <p>No tutorials available at the moment.</p>
      )}
      <button
        disabled={!selectedTutorialId}
        onClick={() => selectedTutorialId && handleAccessTutorial(selectedTutorialId)}
      >
        Access Selected Tutorial
      </button>
    </div>
  );
};

export default TutorialAccess;
