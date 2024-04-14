import React from 'react';
import EnvironmentalImpactScore from './EnvironmentalImpactScore';

interface Product {
  id: string;
  name: string;
  description: string;
  imageUrl: string;
  environmentalImpactScore: number;
}

interface ProductCardProps {
  product: Product;
}

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  return (
    <div className="product-card">
      <img src={product.imageUrl} alt={product.name} />
      <h3>{product.name}</h3>
      <p>{product.description}</p>
      <EnvironmentalImpactScore score={product.environmentalImpactScore} />
    </div>
  );
};

export default ProductCard;
