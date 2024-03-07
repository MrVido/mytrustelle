import React, { useState } from 'react';
import { useCheckCookie } from './useCheckCookie'; // Replace with your cookie checking hook

interface CookieConsentProps {
  onAccept: () => void; // Function to be called when user accepts cookies
  onDecline: () => void; // Function to be called when user declines cookies
}

const CookieConsentManager: React.FC<CookieConsentProps> = ({ onAccept, onDecline }) => {
  const [showConsentBanner, setShowConsentBanner] = useState(true);
  const hasConsentCookie = useCheckCookie('my-cookie-name'); // Replace with your cookie name

  const handleAccept = () => {
    setShowConsentBanner(false);
    onAccept(); // Call provided onAccept function
  };

  const handleDecline = () => {
    setShowConsentBanner(false);
    onDecline(); // Call provided onDecline function
  };

  if (!hasConsentCookie && showConsentBanner) {
    return (
      <div className="cookie-consent-banner">
        <p>
          This website uses cookies to improve your experience. By continuing to use this website,
          you consent to the use of cookies.
        </p>
        <button onClick={handleAccept}>Accept</button>
        <button onClick={handleDecline}>Decline</button>
      </div>
    );
  }

  return null; // Don't show banner if consent cookie exists or banner has been dismissed
};

export default CookieConsentManager;
