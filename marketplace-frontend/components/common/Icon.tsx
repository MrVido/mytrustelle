import React from 'react';

interface IconProps {
  name: string; // Name of the icon (refer to your icon library's documentation)
  size?: number; // Optional size of the icon in pixels
  color?: string; // Optional color for the icon
}

const Icon: React.FC<IconProps> = ({ name, size = 24, color = 'currentColor' }) => {
  // Import the icon component based on your chosen library
  // For example, using @heroicons/react:
  const IconComponent = require(`@heroicons/react/outline/${name}`).default;

  return <IconComponent size={size} color={color} />;
};

export default Icon;
