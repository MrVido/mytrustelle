import React, { useState, useEffect } from 'react';

interface Product {
  id: string;
  name: string;
  basePrice: number; // Initial price of the product
  // Add other product properties as needed (e.g., stock, category)
}

const DynamicPricingTool: React.FC = ({ product }: { product: Product }) => {
  // State variables
  const [currentPrice, setCurrentPrice] = useState(product.basePrice);
  const [discount, setDiscount] = useState(0); // Initial discount percentage

  // Function to simulate dynamic pricing based on logic (replace with your actual logic)
  const updatePrice = () => {
    // Example: Increase price by 10% during peak hours
    const newPrice = product.basePrice * 1.1;
    setCurrentPrice(newPrice);
  };

  // Fetch data or perform calculations to update price (replace with your implementation)
  useEffect(() => {
    // Example: Simulate fetching data every 5 seconds to update price
    const intervalId = setInterval(updatePrice, 5000); // Update price every 5 seconds

    return () => clearInterval(intervalId); // Cleanup function to stop interval on unmount
  }, []);

  // Function to handle user input for discount
  const handleDiscountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(event.target.value, 10); // Parse input to number
    if (!isNaN(value) && value >= 0 && value <= 100) {
      setDiscount(value);
      const discountedPrice = product.basePrice * (1 - value / 100); // Apply discount
      setCurrentPrice(discountedPrice);
    }
  };

  return (
    <div className="dynamic-pricing-tool">
      <h2>{product.name}</h2>
      <p>Base Price: ${product.basePrice.toFixed(2)}</p>
      <p>Current Price: ${currentPrice.toFixed(2)}</p>
      <label htmlFor="discount">Discount (%)</label>
      <input
        type="number"
        id="discount"
        value={discount}
        onChange={handleDiscountChange}
        min={0}
        max={100}
      />
      <button onClick={updatePrice}>Update Price (Simulate)</button>
    </div>
  );
};

export default DynamicPricingTool;
