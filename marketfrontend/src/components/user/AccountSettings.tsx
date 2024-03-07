import React, { useState } from 'react';

const AccountSettings = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Submit updated settings logic here
    console.log('Updated settings:', { username, email });
  };

  return (
    <div className="account-settings">
      <h2>Account Settings</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Username:
          <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} />
        </label>
        <label>
          Email:
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        </label>
        <button type="submit">Update</button>
      </form>
    </div>
  );
};

export default AccountSettings;
