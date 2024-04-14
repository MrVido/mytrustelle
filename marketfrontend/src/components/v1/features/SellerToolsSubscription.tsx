import React, { useState, useEffect } from 'react';

interface SellerToolsSubscriptionProps {
  isSubscribed: boolean; // Current subscription status
  onSubscribe: () => Promise<void>; // Function to handle subscription
  onUnsubscribe: () => Promise<void>; // Function to handle unsubscription
}

const SellerToolsSubscription: React.FC<SellerToolsSubscriptionProps> = ({
  isSubscribed,
  onSubscribe,
  onUnsubscribe,
}) => {
  const [subscriptionStatus, setSubscriptionStatus] = useState(isSubscribed);

  useEffect(() => {
    setSubscriptionStatus(isSubscribed);
  }, [isSubscribed]);

  const handleSubscriptionToggle = async () => {
    if (subscriptionStatus) {
      await onUnsubscribe();
    } else {
      await onSubscribe();
    }
    setSubscriptionStatus(!subscriptionStatus);
  };

  return (
    <div className="seller-tools-subscription">
      <h2>Seller Tools Subscription</h2>
      <p>
        Access advanced tools to manage your listings, gain insights into your sales, and optimize your store's performance.
      </p>
      <button onClick={handleSubscriptionToggle}>
        {subscriptionStatus ? 'Unsubscribe from Seller Tools' : 'Subscribe to Seller Tools'}
      </button>
    </div>
  );
};

export default SellerToolsSubscription;
