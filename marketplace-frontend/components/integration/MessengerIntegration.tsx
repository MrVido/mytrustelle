import React, { useState } from 'react';

interface Message {
  id: string;
  sender: 'user' | 'bot';
  content: string;
}

interface MessengerIntegrationProps {
  onSendMessage: (content: string) => void; // Function to handle sending messages to the server/API
}

const MessengerIntegration: React.FC<MessengerIntegrationProps> = ({ onSendMessage }) => {
  const [messages, setMessages] = useState<Message[]>([]); // Stores the conversation
  const [newMessage, setNewMessage] = useState(''); // Stores the new message input by the user

  const handleNewMessageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setNewMessage(event.target.value);
  };

  const handleSendMessage = () => {
    // Example function to send a message. You would replace this with your actual message sending logic, e.g., calling an API.
    const message: Message = {
      id: Date.now().toString(), // Simple unique ID for each message
      sender: 'user',
      content: newMessage,
    };
    setMessages([...messages, message]);
    onSendMessage(newMessage); // Propagate the message up to be sent to the server/API
    setNewMessage(''); // Clear the input field after sending
  };

  return (
    <div className="messenger-integration">
      <div className="messages">
        {messages.map((msg) => (
          <div key={msg.id} className={`message ${msg.sender}`}>
            {msg.content}
          </div>
        ))}
      </div>
      <div className="message-input">
        <input
          type="text"
          value={newMessage}
          onChange={handleNewMessageChange}
          placeholder="Type a message..."
        />
        <button onClick={handleSendMessage}>Send</button>
      </div>
    </div>
  );
};

export default MessengerIntegration;
