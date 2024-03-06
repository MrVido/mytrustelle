import React, { useState } from 'react';
import OpenAI from 'openai';

const openai = new OpenAI({
  apiKey: process.env.OPENAI_API_KEY,
});

const SEOGenerator = () => {
  const [content, setContent] = useState('');
  const [seoTitle, setSeoTitle] = useState('');
  const [seoDescription, setSeoDescription] = useState('');

  const handleContentChange = (event) => {
    setContent(event.target.value);
  };

  const generateSEO = async () => {
    try {
      // Generate SEO title
      const titleResponse = await openai.createCompletion({
        model: "text-davinci-003",
        prompt: `Generate an SEO-friendly title for the following content: ${content}`,
        temperature: 0.7,
        max_tokens: 60,
      });
      setSeoTitle(titleResponse.data.choices[0].text.trim());

      // Generate SEO description
      const descriptionResponse = await openai.createCompletion({
        model: "text-davinci-003",
        prompt: `Generate an SEO-friendly description for the following content: ${content}`,
        temperature: 0.7,
        max_tokens: 160,
      });
      setSeoDescription(descriptionResponse.data.choices[0].text.trim());

    } catch (error) {
      console.error('Error generating SEO content:', error);
    }
  };

  return (
    <div>
      <h2>SEO Content Generator</h2>
      <textarea value={content} onChange={handleContentChange} placeholder="Enter content to analyze..." />
      <button onClick={generateSEO}>Generate SEO Content</button>
      {seoTitle && <p><strong>SEO Title:</strong> {seoTitle}</p>}
      {seoDescription && <p><strong>SEO Description:</strong> {seoDescription}</p>}
    </div>
  );
};

export default SEOGenerator;
