import React from 'react';

// Define props for the SideAd component
interface SideAdProps {
  imageUrl: string;
  altText?: string;
  targetUrl: string;
  onClick?: () => void; // Optional click handler for tracking ad clicks
}

const SideAd: React.FC<SideAdProps> = ({ imageUrl, altText = "Advertisement", targetUrl, onClick }) => {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', margin: '10px 0' }}>
      <a href={targetUrl} target="_blank" rel="noopener noreferrer" onClick={onClick}>
        <img src={imageUrl} alt={altText} style={{ maxWidth: '100%', maxHeight: '300px', display: 'block' }} />
      </a>
    </div>
  );
};

export default SideAd;
