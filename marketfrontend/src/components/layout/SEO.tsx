import React from 'react';
import  Helmet from 'react-helmet'; // Assuming you're using react-helmet

interface SEOProps {
  title?: string; // Title displayed in search engine results (SERP)
  description?: string; // Brief description of the page for SERP
  keywords?: string[]; // Optional keywords for SEO
  openGraph?: {
    url?: string; // URL of the current page
    type?: string; // e.g., "website", "article"
    title?: string; // Open Graph title (optional, defaults to title)
    description?: string; // Open Graph description (optional, defaults to description)
    image?: string; // URL of an image for the page
  };
}

const SEO: React.FC<SEOProps> = ({ title = '', description = '', keywords = [], openGraph }) => {
  const ogTitle = openGraph?.title || title;
  const ogDescription = openGraph?.description || description;

  return (
    <Helmet>
      <title>{title}</title>
      <meta name="description" content={description} />
      {keywords.length > 0 && <meta name="keywords" content={keywords.join(', ')} />}
      {openGraph && (
        <>
          <meta property="og:url" content={openGraph.url || window.location.href} />
          <meta property="og:type" content={openGraph.type || 'website'} />
          <meta property="og:title" content={ogTitle} />
          <meta property="og:description" content={ogDescription} />
          {openGraph.image && <meta property="og:image" content={openGraph.image} />}
        </>
      )}
    </Helmet>
  );
};

export default SEO;
