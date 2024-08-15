import React, { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import './ProtectedRoute.css';

const ProtectedRoute = ({ children }) => {
  const token = Cookies.get('token');
  const [showModal, setShowModal] = useState(!token);
  const navigate = useNavigate();

  const handleLoginRedirect = () => {
    setShowModal(false);
    navigate('/login');
  };

  if (showModal) {
    return (
      <div className="modal-overlay">
        <div className="modal">
          <h2>Access Denied</h2>
          <p>You need to be logged in to view this page.</p>
          <button onClick={handleLoginRedirect}>Go to Login</button>
        </div>
      </div>
    );
  }

  return children;
};

export default ProtectedRoute;
