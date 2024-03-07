import React, { useState, useRef, useEffect } from 'react';
import  Link  from 'next/link'; // For routing if applicable

interface LoginFormProps {
  onSubmit: (username: string, password: string) => void; // Function to handle successful login
  errorMessage?: string; // Error message to display, if any
  isLoading?: boolean; // Flag indicating if login is in progress
}

const LoginForm: React.FC<LoginFormProps> = ({ onSubmit, errorMessage, isLoading }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const usernameRef = useRef<HTMLInputElement>(null); // Ref for username input
  const passwordRef = useRef<HTMLInputElement>(null); // Ref for password input

  // Focus username field on component mount
  useEffect(() => {
    usernameRef.current?.focus();
  }, []);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit(username.trim(), password.trim()); // Trim whitespace from inputs
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>
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
          autoComplete="off" // Prevent browser autocompletion
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
          ref={passwordRef}
          required
          autoComplete="off" // Prevent browser autocompletion
        />
      </label>
      <div className="form-actions">
        <button type="submit" disabled={isLoading}>
          {isLoading ? 'Loading...' : 'Login'}
        </button>
        <Link href="/forgot-password">Forgot Password?</Link>
      </div>
    </form>
  );
};

export default LoginForm;
