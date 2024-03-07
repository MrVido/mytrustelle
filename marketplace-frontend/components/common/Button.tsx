import React from 'react';

interface ButtonProps {
  text: string;
  type?: 'button' | 'submit' | 'reset'; // Optional button type
  variant?: 'primary' | 'secondary' | 'outline' | 'link'; // Optional variant
  onClick?: () => void; // Optional click handler function
  disabled?: boolean; // Optional flag to disable the button
}

const Button: React.FC<ButtonProps> = ({ text, type = 'button', variant = 'primary', onClick, disabled }) => {
  const buttonClasses = `button button-${variant}`; // Base class with variant suffix

  return (
    <button type={type} className={buttonClasses} onClick={onClick} disabled={disabled}>
      {text}
    </button>
  );
};

export default Button;
