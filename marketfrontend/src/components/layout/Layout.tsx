import React from 'react';
import Header from './Header'; // Assuming Header.tsx is in the same directory
import Footer from './Footer'; // Assuming Footer.tsx is in the same directory

interface LayoutProps {
  children: React.ReactNode; // Content to be displayed in the main area
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="layout">
      <Header title="Your App Title" /> {/* Pass title prop to Header */}
      <main className="main-content">{children}</main>
      <Footer />
    </div>
  );
};

export default Layout;
