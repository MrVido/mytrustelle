import React from 'react';

interface MembersOnlyContentProps {
  userHasAccess: boolean;
  content: JSX.Element; // Assuming the exclusive content is passed as a prop, can be any JSX element
  onUpgradeClick: () => void; // Function to handle upgrading the user's membership
}

const MembersOnlyContent: React.FC<MembersOnlyContentProps> = ({ userHasAccess, content, onUpgradeClick }) => {
  if (userHasAccess) {
    return <div className="members-only-content">{content}</div>;
  } else {
    return (
      <div className="access-denied">
        <p>This content is exclusive to our premium members.</p>
        <button onClick={onUpgradeClick}>Upgrade Now</button>
      </div>
    );
  }
};

export default MembersOnlyContent;
