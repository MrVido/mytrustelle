import React from 'react';

interface UserProfile {
  username: string;
  email: string;
  joinDate: string;
}

interface ProfileOverviewProps {
  profile: UserProfile;
}

const ProfileOverview: React.FC<ProfileOverviewProps> = ({ profile }) => {
  return (
    <div className="profile-overview">
      <h2>Profile Overview</h2>
      <p>Username: {profile.username}</p>
      <p>Email: {profile.email}</p>
      <p>Join Date: {profile.joinDate}</p>
    </div>
  );
};

export default ProfileOverview;
