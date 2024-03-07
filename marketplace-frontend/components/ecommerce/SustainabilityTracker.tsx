import React, { useState, useEffect } from 'react';

interface SustainabilityAction {
  id: string;
  name: string;
  description: string;
  impact: string; // Environmental impact of the action (e.g., "Reduces carbon footprint")
  frequency: string; // How often the user performs the action (e.g., "Daily", "Weekly")
}

const SustainabilityTracker: React.FC = () => {
  // State variables
  const [actions, setActions] = useState<SustainabilityAction[]>([]);
  const [selectedAction, setSelectedAction] = useState<SustainabilityAction | null>(null);

  // Fetch sustainability actions from an API or local data (replace with your logic)
  useEffect(() => {
    const fetchActions = async () => {
      const response = await fetch('https://api.example.com/sustainability-actions'); // Replace with your API endpoint
      const data = await response.json();
      setActions(data);
    };

    fetchActions();
  }, []);

  // Function to handle selecting an action from the list
  const handleActionSelect = (action: SustainabilityAction) => {
    setSelectedAction(action);
  };

  return (
    <div className="sustainability-tracker">
      <h2>Sustainability Tracker</h2>
      {/* List of available actions */}
      <ul className="actions-list">
        {actions.map((action) => (
          <li key={action.id}>
            <button onClick={() => handleActionSelect(action)}>
              {action.name} ({action.frequency})
            </button>
          </li>
        ))}
      </ul>
      {/* Details of the selected action (if any) */}
      {selectedAction && (
        <div className="selected-action">
          <h3>{selectedAction.name}</h3>
          <p>{selectedAction.description}</p>
          <p>Impact: {selectedAction.impact}</p>
        </div>
      )}
    </div>
  );
};

export default SustainabilityTracker;
