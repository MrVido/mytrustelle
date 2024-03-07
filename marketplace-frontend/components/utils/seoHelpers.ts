// Example SEO helper function to update page metadata
export const updatePageMetadata = (title: string, description: string) => {
    document.title = title;
    const metaDescription = document.querySelector('meta[name="description"]');
    if (metaDescription) {
      metaDescription.setAttribute("content", description);
    }
  
    // You could also update other SEO-related tags (e.g., Open Graph tags) here
  };
  
  // Add more SEO-related functions here
  