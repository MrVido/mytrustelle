import React from 'react';

// Define props for the InFeedAd component
interface InFeedAdProps {
  imageUrl: string;
  headline: string;
  description: string;
  targetUrl: string;
  onClick?: () => void; // Optional click handler for tracking ad clicks
}

const InFeedAd: React.FC<InFeedAdProps> = ({ imageUrl, headline, description, targetUrl, onClick }) => {
  return (
    <a href={targetUrl} target="_blank" rel="noopener noreferrer" onClick={onClick} style={{ textDecoration: 'none', color: 'inherit' }}>
      <div style={{ display: 'flex', alignItems: 'center', padding: '10px', backgroundColor: '#f4f4f4', borderRadius: '5px', margin: '20px 0', boxShadow: '0 2px 4px rgba(0,0,0,0.1)' }}>
        <img src={imageUrl} alt={headline} style={{ width: '120px', height: '120px', marginRight: '15px', borderRadius: '5px' }} />
        <div>
          <h3 style={{ margin: '0 0 10px 0' }}>{headline}</h3>
          <p style={{ margin: 0 }}>{description}</p>
        </div>
      </div>
    </a>
  );
};

export default InFeedAd;
