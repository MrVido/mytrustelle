import React from 'react';

interface SocialShareButtonProps {
  content: string; // The content to be shared
  onShare: () => void; // Callback function to handle the share action
}

const SocialShareButton: React.FC<SocialShareButtonProps> = ({ content, onShare }) => {
  const shareContent = () => {
    console.log(`Sharing content: ${content}`);
    onShare(); // Trigger the provided callback function
  };

  return (
    <button onClick={shareContent} className="social-share-button">
      Share
    </button>
  );
};

export default SocialShareButton;
