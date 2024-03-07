import React, { useState, useEffect, MouseEvent, RefObject, useRef, LegacyRef } from 'react';

interface TooltipProps {
  content: string;
}

const Tooltip: React.FC<TooltipProps> = ({ content }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [tooltipCoords, setTooltipCoords] = useState({ top: 0, left: 0 });
  const triggerRef = useRef<HTMLElement>(null); // Use a ref for the trigger element

  const handleMouseEnter = (event: MouseEvent<HTMLElement>) => {
    if (triggerRef.current) { // Ensure ref is set
      setIsOpen(true);
      setTooltipCoords({ top: event.clientY + 10, left: event.clientX + 10 });
    }
  };

  const handleMouseLeave = () => {
    setIsOpen(false);
  }; // Define the missing handleMouseLeave function

  useEffect(() => {
    const handleResize = () => {
      // ... (rest of the resize logic)
    };

    window.addEventListener('resize', handleResize);

    return () => window.removeEventListener('resize', handleResize);
  }, [isOpen]); // Update dependencies to avoid infinite loop (isOpen only)

  return (
    <>
      <div ref={triggerRef as LegacyRef<HTMLDivElement>} onMouseEnter={handleMouseEnter} onMouseLeave={handleMouseLeave}>
  {/* Render the trigger element's content within this div */}
        </div>
      {isOpen && (
        <div className="tooltip" style={{ top: tooltipCoords.top, left: tooltipCoords.left }}>
          {content}
        </div>
      )}
    </>
  );
};

export default Tooltip;
