import React from 'react';

interface Post {
  id: string;
  author: string;
  content: string;
  timestamp: string; // Consider using a date string for simplicity
}

interface SocialFeedProps {
  posts: Post[];
}

const SocialFeed: React.FC<SocialFeedProps> = ({ posts }) => {
  return (
    <div className="social-feed">
      {posts.map((post) => (
        <div key={post.id} className="post">
          <div className="post-author">{post.author}</div>
          <div className="post-content">{post.content}</div>
          <div className="post-timestamp">{post.timestamp}</div>
        </div>
      ))}
    </div>
  );
};

export default SocialFeed;
