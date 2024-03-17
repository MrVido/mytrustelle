import React, { useState } from 'react';
import OpenAI from 'openai';

const openai = new OpenAI({
  apiKey: process.env.OPENAI_API_KEY,
});

const Chatbot = () => {
  const [userInput, setUserInput] = useState('');
  const [chatHistory, setChatHistory] = useState([]);

  const handleUserInput = (event) => {
    setUserInput(event.target.value);
  };

  const submitChat = async () => {
    try {
      const response = await openai.createCompletion({
        model: "text-davinci-003",
        prompt: userInput,
        temperature: 0.7,
        max_tokens: 150,
      });
      setChatHistory([...chatHistory, { role: 'user', content: userInput }, { role: 'bot', content: response.data.choices[0].text }]);
      setUserInput(''); // Reset input after submission
    } catch (error) {
      console.error('Error submitting chat to OpenAI:', error);
    }
  };

  return (
    <div>
      <h2>Chatbot</h2>
      <div className="chat-history">
        {chatHistory.map((chat, index) => (
          <div key={index} className={`chat-message ${chat.role}`}>
            <p>{chat.content}</p>
          </div>
        ))}
      </div>
      <input type="text" value={userInput} onChange={handleUserInput} />
      <button onClick={submitChat}>Send</button>
    </div>
  );
};

export default Chatbot;
