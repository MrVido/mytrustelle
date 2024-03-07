import React, { useState } from 'react';

interface DigitalCurrency {
  name: string;
  symbol: string;
  currentPrice: number;
}

interface DigitalCurrencyPurchaseProps {
  currencies: DigitalCurrency[]; // List of digital currencies available for purchase
  onPurchase: (currencySymbol: string, amount: number) => void; // Function to call when currency is purchased
}

const DigitalCurrencyPurchase: React.FC<DigitalCurrencyPurchaseProps> = ({ currencies, onPurchase }) => {
  const [selectedCurrency, setSelectedCurrency] = useState<string>('');
  const [amount, setAmount] = useState<number>(0);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onPurchase(selectedCurrency, amount);
  };

  return (
    <div className="digital-currency-purchase">
      <h2>Purchase Digital Currency</h2>
      <form onSubmit={handleSubmit}>
        <select
          value={selectedCurrency}
          onChange={(e) => setSelectedCurrency(e.target.value)}
          required
        >
          <option value="">Select Currency</option>
          {currencies.map((currency) => (
            <option key={currency.symbol} value={currency.symbol}>
              {currency.name} ({currency.symbol})
            </option>
          ))}
        </select>
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(Number(e.target.value))}
          min="0"
          step="any"
          placeholder="Amount to Purchase"
          required
        />
        <button type="submit">Purchase</button>
      </form>
    </div>
  );
};

export default DigitalCurrencyPurchase;
