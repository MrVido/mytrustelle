import React from 'react';

interface HeaderProps {
  title?: string; // Optional title for the header
}

const Header: React.FC<HeaderProps> = ({ title = 'My App' }) => {
  return (
    <header className="header">
      <h1>{title}</h1>
    </header>
  );
};

export default Header;
