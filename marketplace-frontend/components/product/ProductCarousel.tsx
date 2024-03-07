import React from 'react';
import ProductCard from './ProductCard';
import { Product } from '../../types';

interface ProductCarouselProps {
  products: Product[]; // Use the Product interface from ProductCard.tsx
}

const ProductCarousel: React.FC<ProductCarouselProps> = ({ products }) => {
  return (
    <div className="product-carousel">
      {products.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  );
};

export default ProductCarousel;
