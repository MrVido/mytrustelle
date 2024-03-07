import React, { useState } from 'react';

interface PrioritySupportOptionProps {
  isEnabled: boolean;
  onTogglePrioritySupport: (isEnabled: boolean) => void;
}

const PrioritySupportOption: React.FC<PrioritySupportOptionProps> = ({ isEnabled, onTogglePrioritySupport }) => {
  const [isPrioritySupportEnabled, setIsPrioritySupportEnabled] = useState(isEnabled);

  const handleToggle = () => {
    const newStatus = !isPrioritySupportEnabled;
    setIsPrioritySupportEnabled(newStatus);
    onTogglePrioritySupport(newStatus);
  };

  return (
    <div className="priority-support-option">
      <h3>Priority Support</h3>
      <p>Get faster support responses by enabling Priority Support.</p>
      <button onClick={handleToggle}>
        {isPrioritySupportEnabled ? 'Disable Priority Support' : 'Enable Priority Support'}
      </button>
    </div>
  );
};

export default PrioritySupportOption;
