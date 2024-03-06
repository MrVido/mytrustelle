# MyTrustelle Frontend

## Overview

MyTrustelle is a modern online marketplace platform designed to provide a seamless shopping experience. This repository contains the frontend implementation of the platform.

## Features

- **User Authentication & Management:** Secure handling of user data with JWT for session management.
- **Listing Operations:** Full CRUD functionality for listings with image support.
- **Real-Time Chat:** WebSocket-based communication for users.
- **Secure Transactions:** Integration with Stripe and PayPal for financial transactions.
- **Recommendations:** Dynamic listing recommendations based on user activity.
- **Admin Dashboard:** Tools for administrative oversight and analytics.
- **Social Sharing:** Functions to share listings on major social media platforms.

## Technologies Used

- **Framework:** Next.js
- **Styling:** Tailwind CSS, Emotion
- **State Management:** Redux Toolkit, React Query
- **Authentication:** NextAuth.js
- **HTTP Requests:** Axios
- **Charts:** Chart.js, Recharts
- **Forms:** Formik, React Hook Form
- **Date Management:** Day.js
- **UI Components:** Material-UI, Radix UI
- **Image Processing:** TensorFlow.js
- **API Integration:** OpenAI
- **Others:** Lodash, UUID, Yup, and more.

## Getting Started

### Prerequisites

- Node.js (>= v14)
- npm or yarn
- Backend API (refer to backend repository)

### Installation

1. Clone the repository:
git clone https://github.com/your-username/mytrustelle-frontend.git
cd mytrustelle-frontend

2. Install dependencies: npm install

3. Set up environment variables:

Copy the `.env.example` file to a new file named `.env.local` and adjust the variables to match your environment.

### Running the Application

npm run dev

The application will start on `http://localhost:3000` by default.

## Contributing

We welcome contributions from the community. Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) for detailed instructions on how to contribute, code of conduct, and the process for submitting pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.