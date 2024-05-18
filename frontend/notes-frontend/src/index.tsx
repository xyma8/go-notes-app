import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import AddNotePage from './pages/AddNotePage';
import NotePage from './pages/NotePage';

const router = createBrowserRouter([
  {
    path: '/',
    element: <AddNotePage/>
  },
  {
    path: '/notes/add',
    element: <AddNotePage/>
  },
  {
    path: '/notes/:noteId',
    element: <NotePage/>
  }
]);

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);


reportWebVitals();
