import React, { useState } from 'react';

interface SupportTicketFormProps {
  onSubmit: (ticket: { name: string; email: string; description: string }) => void;
}

const SupportTicketForm: React.FC<SupportTicketFormProps> = ({ onSubmit }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [description, setDescription] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({ name, email, description });
    setName('');
    setEmail('');
    setDescription('');
  };

  return (
    <form className="support-ticket-form" onSubmit={handleSubmit}>
      <div className="form-group">
        <label htmlFor="name">Name</label>
        <input
          type="text"
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
      </div>
      <div className="form-group">
        <label htmlFor="email">Email</label>
        <input
          type="email"
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
      </div>
      <div className="form-group">
        <label htmlFor="description">Issue Description</label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
        />
      </div>
      <button type="submit" className="submit-btn">
        Submit Ticket
      </button>
    </form>
  );
};

export default SupportTicketForm;
