import React, { useState, useRef, useEffect } from 'react';
import  Link  from 'next/link'; // For routing if applicable

interface RegisterFormProps {
  onSubmit: (formData: { username: string; email: string; password: string; confirmPassword: string }) => void; // Function to handle successful registration
  errorMessage?: string; // Error message to display, if any
  isLoading?: boolean; // Flag indicating if registration is in progress
}

const RegisterForm: React.FC<RegisterFormProps> = ({ onSubmit, errorMessage, isLoading }) => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const confirmPasswordRef = useRef<HTMLInputElement>(null); // Ref for confirm password input
  const [confirmPassword, setConfirmPassword] = useState('');
  const usernameRef = useRef<HTMLInputElement>(null); // Moved outside useEffect

  // Focus username field on component mount (optional)
  useEffect(() => {
    const usernameInput = usernameRef.current;
    if (usernameInput) {
      usernameInput.focus();
    }
  }, []);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (password !== confirmPassword) {
      alert('Passwords do not match!');
      return;
    }
    onSubmit({
        username, email, password,
        confirmPassword: ''
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Register</h2>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
      <label htmlFor="username">
        Username:
        <input
          type="text"
          id="username"
          name="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          ref={usernameRef}
          required
        />
      </label>
      <label htmlFor="email">
        Email:
        <input
          type="email"
          id="email"
          name="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
      </label>
      <label htmlFor="password">
        Password:
        <input
          type="password"
          id="password"
          name="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
      </label>
      <label htmlFor="confirm-password">
        Confirm Password:
        <input
          type="password"
          id="confirm-password"
          name="confirm-password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          ref={confirmPasswordRef}
          required
        />
      </label>
      <div className="form-actions">
        <button type="submit" disabled={isLoading}>
          {isLoading ? 'Loading...' : 'Register'}
        </button>
        <Link href="/login">Already have an account? Login</Link>
      </div>
    </form>
  );
};

export default RegisterForm;
