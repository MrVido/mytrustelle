import React, { useState } from 'react';

interface Product {
  id: string;
  name: string;
  price: number;
}

interface CustomBundleCreatorProps {
  availableProducts: Product[]; // List of products that can be included in the bundle
  onSubmit: (selectedProducts: Product[]) => void; // Function to call when a bundle is submitted
}

const CustomBundleCreator: React.FC<CustomBundleCreatorProps> = ({ availableProducts, onSubmit }) => {
  const [selectedProducts, setSelectedProducts] = useState<Product[]>([]);

  const toggleProductSelection = (product: Product) => {
    setSelectedProducts((prevSelected) =>
      prevSelected.find((p) => p.id === product.id)
        ? prevSelected.filter((p) => p.id !== product.id) // Remove the product if it's already selected
        : [...prevSelected, product] // Add the product if it's not already selected
    );
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(selectedProducts);
  };

  return (
    <div className="custom-bundle-creator">
      <h2>Create Your Custom Bundle</h2>
      <form onSubmit={handleSubmit}>
        {availableProducts.map((product) => (
          <div key={product.id} className="product-option">
            <label>
              <input
                type="checkbox"
                checked={selectedProducts.some((p) => p.id === product.id)}
                onChange={() => toggleProductSelection(product)}
              />
              {product.name} - ${product.price}
            </label>
          </div>
        ))}
        <button type="submit">Create Bundle</button>
      </form>
    </div>
  );
};

export default CustomBundleCreator;
