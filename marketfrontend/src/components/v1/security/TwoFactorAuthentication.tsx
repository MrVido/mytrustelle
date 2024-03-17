import React, { useState } from 'react';

interface TwoFactorAuthenticationProps {
  onVerifyCode: (code: string) => Promise<boolean>; // Function to verify the code, returns true if successful
  onRequestNewCode: () => void; // Function to request a new code
}

const TwoFactorAuthentication: React.FC<TwoFactorAuthenticationProps> = ({
  onVerifyCode,
  onRequestNewCode,
}) => {
  const [code, setCode] = useState('');
  const [verificationStatus, setVerificationStatus] = useState<'idle' | 'verifying' | 'failed' | 'success'>('idle');

  const handleVerifyCode = async () => {
    setVerificationStatus('verifying');
    const isSuccess = await onVerifyCode(code);
    setVerificationStatus(isSuccess ? 'success' : 'failed');
  };

  const handleRequestNewCode = () => {
    onRequestNewCode();
    setVerificationStatus('idle'); // Reset status when requesting a new code
  };

  return (
    <div className="two-factor-authentication">
      <h2>Two-Factor Authentication</h2>
      {verificationStatus === 'failed' && <p className="error">Verification failed. Please try again.</p>}
      <input
        type="text"
        placeholder="Enter your 2FA code"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        disabled={verificationStatus === 'verifying'}
      />
      <button onClick={handleVerifyCode} disabled={verificationStatus === 'verifying'}>
        Verify Code
      </button>
      <button onClick={handleRequestNewCode} disabled={verificationStatus === 'verifying'}>
        Request New Code
      </button>
      {verificationStatus === 'success' && <p className="success">Verification successful!</p>}
    </div>
  );
};

export default TwoFactorAuthentication;
