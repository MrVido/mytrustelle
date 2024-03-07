import React from 'react';

interface HelpArticleProps {
  id: number;
  title: string;
  content: string;
  datePublished: string;
}

const HelpArticle: React.FC<HelpArticleProps> = ({ title, content, datePublished }) => {
  return (
    <article className="help-article">
      <h2>{title}</h2>
      <div className="help-article-date">{new Date(datePublished).toLocaleDateString()}</div>
      <div className="help-article-content" dangerouslySetInnerHTML={{ __html: content }} />
    </article>
  );
};

export default HelpArticle;
