import React from 'react';

interface AlertProps {
  type?: 'success' | 'error' | 'info' | 'warning'; // Define allowed alert types
  message: string;
  onClose?: () => void; // Optional callback function for closing the alert
}

const Alert: React.FC<AlertProps> = ({ type = 'info', message, onClose }) => {
  const alertClasses = `alert alert-${type}`; // Base class with type suffix

  return (
    <div className={alertClasses} role="alert">
      {message}
      {onClose && (
        <button type="button" className="close" onClick={onClose} aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      )}
    </div>
  );
};

export default Alert;
