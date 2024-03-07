import React, { useState } from 'react';

const GiftCardDigital: React.FC = () => {
  const [selectedAmount, setSelectedAmount] = useState<number>(50); // Default gift card amount
  const [email, setEmail] = useState<string>('');

  const handleAmountChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedAmount(Number(event.target.value));
  };

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // Implement your submission logic here, e.g., API call to process the gift card purchase
    console.log(`Gift card of $${selectedAmount} purchased for ${email}`);
  };

  return (
    <div className="gift-card-digital">
      <h2>Buy a Digital Gift Card</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="giftCardAmount">Amount</label>
          <select id="giftCardAmount" value={selectedAmount} onChange={handleAmountChange}>
            <option value="25">$25</option>
            <option value="50">$50</option>
            <option value="100">$100</option>
            <option value="200">$200</option>
          </select>
        </div>
        <div className="form-group">
          <label htmlFor="recipientEmail">Recipient's Email</label>
          <input
            type="email"
            id="recipientEmail"
            value={email}
            onChange={handleEmailChange}
            required
          />
        </div>
        <button type="submit">Purchase Gift Card</button>
      </form>
    </div>
  );
};

export default GiftCardDigital;
