import React from 'react';
import EnvironmentalImpactScore from './EnvironmentalImpactScore';
import { Product } from '../../../../marketplace-frontend/types';

interface ProductPageProps {
  product: Product;
}

const ProductPage: React.FC<ProductPageProps> = ({ product }) => {
  return (
    <div className="product-page">
      <img src={product.imageUrl} alt={product.name} />
      <h1>{product.name}</h1>
      <p>{product.description}</p>
      <EnvironmentalImpactScore score={product.environmentalImpactScore} />
      {/* More product details here */}
    </div>
  );
};

export default ProductPage;
