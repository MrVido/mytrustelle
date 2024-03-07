import React from 'react';
import '../../styles/skeleton.css'; // Assuming you have a CSS file for styling

interface SkeletonLoaderProps {
  count?: number; // Number of skeleton items to display
  height?: string; // Height of each skeleton item
  width?: string; // Width of each skeleton item
  borderRadius?: string; // Border radius of each skeleton item
}

const SkeletonLoader: React.FC<SkeletonLoaderProps> = ({
  count = 1,
  height = '20px',
  width = '100%',
  borderRadius = '4px',
}) => {
  // Create an array with 'count' number of elements to render multiple loaders
  const loaders = Array.from({ length: count }, (_, index) => index);

  return (
    <>
      {loaders.map((_, index) => (
        <div
          key={index}
          className="skeleton-loader"
          style={{
            height,
            width,
            borderRadius,
          }}
        ></div>
      ))}
    </>
  );
};

export default SkeletonLoader;
