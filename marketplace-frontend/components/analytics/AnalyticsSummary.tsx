import React from 'react';

// Define props for the AnalyticsSummary component if needed
interface AnalyticsSummaryProps {
  totalVisitors: number;
  totalSales: number;
  totalRevenue: number;
  conversionRate: number;
}

const AnalyticsSummary: React.FC<AnalyticsSummaryProps> = ({ totalVisitors, totalSales, totalRevenue, conversionRate }) => {
  return (
    <div style={{ padding: '20px', display: 'flex', justifyContent: 'space-around', backgroundColor: '#f5f5f5', borderRadius: '8px', boxShadow: '0px 4px 6px rgba(0,0,0,0.1)' }}>
      <div>
        <h4>Total Visitors</h4>
        <p>{totalVisitors}</p>
      </div>
      <div>
        <h4>Total Sales</h4>
        <p>{totalSales}</p>
      </div>
      <div>
        <h4>Total Revenue</h4>
        <p>${totalRevenue.toFixed(2)}</p>
      </div>
      <div>
        <h4>Conversion Rate</h4>
        <p>{conversionRate}%</p>
      </div>
    </div>
  );
};

export default AnalyticsSummary;
