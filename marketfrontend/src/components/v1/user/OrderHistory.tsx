import React from 'react';

interface Order {
  id: string;
  date: string;
  total: number;
  status: string;
}

interface OrderHistoryProps {
  orders: Order[];
}

const OrderHistory: React.FC<OrderHistoryProps> = ({ orders }) => {
  return (
    <div className="order-history">
      <h2>Order History</h2>
      {orders.length > 0 ? (
        <ul>
          {orders.map((order) => (
            <li key={order.id}>
              <span>Order ID: {order.id} - </span>
              <span>Date: {order.date} - </span>
              <span>Total: ${order.total} - </span>
              <span>Status: {order.status}</span>
            </li>
          ))}
        </ul>
      ) : (
        <p>No orders found.</p>
      )}
    </div>
  );
};

export default OrderHistory;
