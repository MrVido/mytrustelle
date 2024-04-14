import React from 'react';

// Define the structure of a blog post's properties
interface BlogPostProps {
  id: number;
  title: string;
  content: string;
  author: string;
  datePublished: string; // You might need to format this date depending on how it's returned from WordPress
  imageUrl?: string; // Optional cover image for the post
}

// The BlogPost component that receives a single post's props
const BlogPost: React.FC<BlogPostProps> = ({ title, content, author, datePublished, imageUrl }) => {
  return (
    <article className="blog-post">
      {imageUrl && <img src={imageUrl} alt={`Cover for ${title}`} className="blog-post-image" />}
      <h2 className="blog-post-title">{title}</h2>
      <div className="blog-post-meta">
        <span className="blog-post-author">{author}</span>
        <span className="blog-post-date">{new Date(datePublished).toLocaleDateString()}</span>
      </div>
      <div className="blog-post-content" dangerouslySetInnerHTML={{ __html: content }} />
    </article>
  );
};

export default BlogPost;
