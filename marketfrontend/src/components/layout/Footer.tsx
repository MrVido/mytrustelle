import React from 'react';

const Footer: React.FC = () => {
  return (
    <footer className="footer">
      <p>&copy; {new Date().getFullYear()} Mytrustelle</p>
    </footer>
  );
};

export default Footer;
