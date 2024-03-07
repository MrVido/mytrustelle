import React from 'react';

// Define props for the BannerAd component
interface BannerAdProps {
  imageUrl: string;
  altText?: string;
  targetUrl: string;
  onClick?: () => void; // Optional click handler for tracking ad clicks
}

const BannerAd: React.FC<BannerAdProps> = ({ imageUrl, altText = "Advertisement", targetUrl, onClick }) => {
  return (
    <a href={targetUrl} target="_blank" rel="noopener noreferrer" onClick={onClick} style={{ display: 'block', cursor: 'pointer' }}>
      <img src={imageUrl} alt={altText} style={{ width: '100%', height: 'auto' }} />
    </a>
  );
};

export default BannerAd;
