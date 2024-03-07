import React, { useState } from 'react';

interface SocialMediaIntegrationManagerProps {
  onShare: (platform: string, content: string) => void; // Function to handle sharing content
}

const SocialMediaIntegrationManager: React.FC<SocialMediaIntegrationManagerProps> = ({
  onShare,
}) => {
  const [content, setContent] = useState('');

  const platforms = ['Facebook', 'Twitter', 'Instagram', 'LinkedIn']; // List of social media platforms

  const handleShare = (platform: string) => {
    onShare(platform, content);
    alert(`Content shared on ${platform}!`);
  };

  return (
    <div className="social-media-integration-manager">
      <h2>Social Media Integration Manager</h2>
      <textarea
        placeholder="Write something to share..."
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />
      <div className="platform-buttons">
        {platforms.map((platform) => (
          <button key={platform} onClick={() => handleShare(platform)}>
            Share on {platform}
          </button>
        ))}
      </div>
    </div>
  );
};

export default SocialMediaIntegrationManager;
