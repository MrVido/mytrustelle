import React, { useState, useEffect } from 'react';
import  Link  from 'next/link'; // For routing to event details
import { EventCard } from './EventCard'; // Import EventCard component (replace with actual path)

interface LiveEvent {
  // Define the structure of an event object
  id: string;
  title: string;
  date: string; // Adjust format based on your data
  time: string;  // Adjust format based on your data
  description: string;
  // Add additional properties as needed (e.g., speaker, image URL)
}

const LiveEvents: React.FC = () => {
  // State for events array and loading state (optional)
  const [events, setEvents] = useState<LiveEvent[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  // Fetch live events from an API or local data (replace with your logic)
  useEffect(() => {
    const fetchEvents = async () => {
      const response = await fetch('https://api.example.com/live-events'); // Replace with your API endpoint
      const data = await response.json();
      setEvents(data);
      setIsLoading(false);
    };

    fetchEvents();
  }, []);

  return (
    <section className="live-events">
      <h2>Live Events</h2>
      {isLoading ? (
        <p>Loading events...</p>
      ) : (
        <ul className="events-list">
          {events.map((event) => (
            <li key={event.id}>
              <EventCard event={event} /> {/* Pass event data to EventCard component */}
            </li>
          ))}
        </ul>
      )}
    </section>
  );
};

export default LiveEvents;
