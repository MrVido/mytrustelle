import React, { useState, useEffect } from 'react';
import  Link  from 'next/link'; // For routing to topic pages
import { Topic } from './Topic'; // Import Topic component (replace with actual path)

interface ForumSectionProps {
  // Props (customize based on your data structure)
  title: string;
  topics: Topic[]; // Array of topic objects
}

const ForumSection: React.FC<ForumSectionProps> = ({ title, topics }) => {
  // State for loading state (optional)
  const [isLoading, setIsLoading] = useState(true);

  // Fetch topics from an API or local data (replace with your logic)
  useEffect(() => {
    const fetchTopics = async () => {
      const response = await fetch('https://api.example.com/topics'); // Replace with your API endpoint
      const data = await response.json();
      setTopics(data); // Update topics state only if needed
      setIsLoading(false);
    };

    fetchTopics();
  }, []);

  return (
    <section className="forum-section">
      <h2>{title}</h2>
      {isLoading ? (
        <p>Loading topics...</p>
      ) : (
        <ul className="topics-list">
          {topics.map((topic) => (
            <li key={topic.id}>
              <Link href={`/forum/topic/${topic.id}`}>
                <a>{topic.title}</a>
              </Link>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
};

export default ForumSection;
function setTopics(data: any) {
    throw new Error('Function not implemented.');
}

