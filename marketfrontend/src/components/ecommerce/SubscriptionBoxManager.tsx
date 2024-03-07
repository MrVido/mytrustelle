import React, { useState, useEffect } from 'react';
import { SubscriptionBox } from './SubscriptionBox'; // Import SubscriptionBox component

interface Subscription {
  id: string;
  name: string;
  description: string;
  price: number;
  // Add other subscription properties as needed
}

const SubscriptionBoxManager: React.FC = () => {
  // State variables
  const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);

  // Fetch subscriptions from an API or local data (replace with your logic)
  useEffect(() => {
    const fetchSubscriptions = async () => {
      const response = await fetch('https://api.example.com/subscriptions'); // Replace with your API endpoint
      const data = await response.json();
      setSubscriptions(data);
    };

    fetchSubscriptions();
  }, []);

  return (
    <div className="subscription-box-manager">
      <h2>Subscription Boxes</h2>
      <ul className="subscriptions-list">
        {subscriptions.map((subscription) => (
          <li key={subscription.id}>
            <SubscriptionBox subscription={subscription} /> {/* Pass subscription data to SubscriptionBox component */}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SubscriptionBoxManager;
