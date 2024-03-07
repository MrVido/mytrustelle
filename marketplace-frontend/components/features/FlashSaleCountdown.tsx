import React, { useState, useEffect } from 'react';

interface FlashSaleCountdownProps {
  saleStartTime: Date; // The start time of the flash sale
}

interface TimeLeft {
  days: number;
  hours: number;
  minutes: number;
  seconds: number;
}

const FlashSaleCountdown: React.FC<FlashSaleCountdownProps> = ({ saleStartTime }) => {
  const calculateTimeLeft = (): TimeLeft => {
    const difference = +saleStartTime - +new Date();
    let timeLeft: TimeLeft = { days: 0, hours: 0, minutes: 0, seconds: 0 };

    if (difference > 0) {
      timeLeft = {
        days: Math.floor(difference / (1000 * 60 * 60 * 24)),
        hours: Math.floor((difference / (1000 * 60 * 60)) % 24),
        minutes: Math.floor((difference / 1000 / 60) % 60),
        seconds: Math.floor((difference / 1000) % 60),
      };
    }

    return timeLeft;
  };

  const [timeLeft, setTimeLeft] = useState<TimeLeft>(calculateTimeLeft());

  useEffect(() => {
    const timer = setTimeout(() => {
      setTimeLeft(calculateTimeLeft());
    }, 1000);

    return () => clearTimeout(timer);
  });

  const timerComponents: JSX.Element[] = [];

  (Object.keys(timeLeft) as (keyof TimeLeft)[]).forEach((interval) => {
    if (!timeLeft[interval]) {
      return;
    }

    timerComponents.push(
      <span key={interval.toString()}>
        {timeLeft[interval]} {interval}{" "}
      </span>
    );
  });

  return (
    <div>
      <h1>Flash Sale Countdown</h1>
      {timerComponents.length ? timerComponents : <span>Time's up!</span>}
    </div>
  );
};

export default FlashSaleCountdown;
