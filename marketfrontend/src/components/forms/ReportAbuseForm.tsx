import React, { useState } from 'react';

interface ReportAbuseFormProps {
  onSubmit: (formData: { reason: string; details: string }) => void; // Function to handle form submission
}

const ReportAbuseForm: React.FC<ReportAbuseFormProps> = ({ onSubmit }) => {
  const [reason, setReason] = useState('');
  const [details, setDetails] = useState('');

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit({ reason, details });
    setReason(''); // Clear form after submission (optional)
    setDetails('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Report Abuse</h2>
      <label htmlFor="reason">
        Reason for Reporting:
        <select id="reason" name="reason" value={reason} onChange={(e) => setReason(e.target.value)} required>
          <option value="">Select a reason</option>
          <option value="bullying">Bullying or harassment</option>
          <option value="spam">Spam or irrelevant content</option>
          <option value="threats">Threats or violence</option>
          <option value="other">Other (please specify in details)</option>
        </select>
      </label>
      <label htmlFor="details">
        Details (Optional):
        <textarea id="details" name="details" value={details} onChange={(e) => setDetails(e.target.value)} />
      </label>
      <button type="submit">Submit Report</button>
    </form>
  );
};

export default ReportAbuseForm;
