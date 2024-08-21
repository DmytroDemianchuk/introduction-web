import React, { useEffect } from 'react';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router-dom';
import './Logout.css'; // Якщо у вас є стилі для Logout

const Logout = () => {
  const navigate = useNavigate();

  useEffect(() => {
    // Видаляємо токен з cookies
    Cookies.remove('token');
    // Перенаправляємо користувача на сторінку входу
    navigate('/login');
  }, [navigate]);

  return (
    <div className="logout-container">
      <h1>Logging out...</h1>
      <p>You are being logged out. Please wait...</p>
    </div>
  );
};

export default Logout;
