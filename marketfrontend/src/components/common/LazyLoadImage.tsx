import React from 'react';
import Image from 'next/image';

interface LazyLoadingImageProps {
  src: string;
  alt: string;
  width?: number;
  height?: number;
  layout?: 'fill' | 'intrinsic' | 'responsive'; // Optional layout prop
}

const LazyLoadingImage: React.FC<LazyLoadingImageProps> = ({ src, alt, width, height, layout }) => {
  return (
    <Image
      src={src}
      alt={alt}
      width={width}
      height={height}
      layout={layout} // Optional layout for different image display behavior
      priority={false} // Explicitly set to false for lazy loading
    />
  );
};

export default LazyLoadingImage;
