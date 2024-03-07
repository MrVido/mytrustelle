import React from 'react';
import Link  from 'gatsby'; // Assuming you're using Gatsby for WordPress integration

interface BlogPostProps {
  title: string;
  excerpt?: string; // Optional excerpt for the post
  slug: string; // Slug for the post URL
  date?: string; // Optional date for the post
  // You can add more props as needed for specific post data (e.g., author, featured image)
}

const BlogPost: React.FC<BlogPostProps> = ({ title, excerpt, slug, date }) => {
  return (
    <article className="blog-post">
      <h2><Link to={slug}>{title}</Link></h2>
      {date && <p className="post-date">{date}</p>}
      {excerpt && <p className="post-excerpt">{excerpt}</p>}
      {/* You can add additional content or components for the full post body */}
      <Link to={slug} className="read-more">Read More</Link>
    </article>
  );
};

export default BlogPost;
