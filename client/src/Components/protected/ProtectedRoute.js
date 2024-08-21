import React from 'react';
import { Navigate } from 'react-router-dom';
import Cookies from 'js-cookie';

const ProtectedRoute = ({ children }) => {
  const token = Cookies.get('token');

  if (!token) {
    // Якщо токену немає, перенаправляємо на сторінку входу
    return <Navigate to="/login" />;
  }

  return children;
};

export default ProtectedRoute;
