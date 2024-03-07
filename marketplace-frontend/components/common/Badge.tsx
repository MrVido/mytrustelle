import React from 'react';

interface BadgeProps {
  text: string;
  variant?: 'primary' | 'secondary' | 'success' | 'error' | 'info'; // Optional variant
}

const Badge: React.FC<BadgeProps> = ({ text, variant = 'primary' }) => {
  const badgeClasses = `badge badge-${variant}`; // Base class with variant suffix

  return <span className={badgeClasses}>{text}</span>;
};

export default Badge;
