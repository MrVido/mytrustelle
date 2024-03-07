import React, { lazy } from 'react';
import Loader from './Loader';

const MyAsyncComponent = lazy(() => import('./MyComponent'));

const MyPage = () => {
  return (
    <>
      <Loader />
      <MyAsyncComponent />
    </>
  );
};

export default MyPage;
