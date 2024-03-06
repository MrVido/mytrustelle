import React from 'react';
// Assuming a chart library like Recharts or Chart.js is being used
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

// Sample data
const data = [
  {
    name: 'Jan',
    sales: 4000,
    visitors: 2400,
  },
  {
    name: 'Feb',
    sales: 3000,
    visitors: 1398,
  },
  // Add more months data
];

const DashboardChart: React.FC = () => {
  return (
    <LineChart
      width={500}
      height={300}
      data={data}
      margin={{
        top: 5,
        right: 30,
        left: 20,
        bottom: 5,
      }}
    >
      <CartesianGrid strokeDasharray="3 3" />
      <XAxis dataKey="name" />
      <YAxis />
      <Tooltip />
      <Legend />
      <Line type="monotone" dataKey="sales" stroke="#8884d8" activeDot={{ r: 8 }} />
      <Line type="monotone" dataKey="visitors" stroke="#82ca9d" />
    </LineChart>
  );
};

export default DashboardChart;
