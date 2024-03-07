import React, { useState } from 'react';

const Modal: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const openModal = () => setIsOpen(true);
  const closeModal = () => setIsOpen(false);

  return (
    <>
      <button onClick={openModal}>Open Modal</button>
      {isOpen && (
        <div className="modal">
          <div className="modal-content">
            {/* Your modal content goes here */}
            <button onClick={closeModal}>Close</button>
          </div>
        </div>
      )}
    </>
  );
};

export default Modal;
