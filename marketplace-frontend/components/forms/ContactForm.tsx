import React, { useState, useRef } from 'react';

interface ContactFormProps {
  onSubmit: (formData: { name: string; email: string; message: string }) => void; // Function to handle form submission
}

const ContactForm: React.FC<ContactFormProps> = ({ onSubmit }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const formRef = useRef<HTMLFormElement>(null); // Use a ref for the form element

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    // Input validation (optional)
    if (!name || !email || !message) {
      alert('Please fill in all required fields.');
      return;
    }

    const formData = { name, email, message };
    onSubmit(formData);

    // Clear form after successful submission
    setName('');
    setEmail('');
    setMessage('');
  };

  return (
    <form ref={formRef} onSubmit={handleSubmit}>
      <label htmlFor="name">
        Name:
        <input type="text" id="name" name="name" value={name} onChange={(e) => setName(e.target.value)} required />
      </label>
      <label htmlFor="email">
        Email:
        <input type="email" id="email" name="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
      </label>
      <label htmlFor="message">
        Message:
        <textarea id="message" name="message" value={message} onChange={(e) => setMessage(e.target.value)} required />
      </label>
      <button type="submit">Send Message</button>
    </form>
  );
};

export default ContactForm;
