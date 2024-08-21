import React from 'react';
import { Link } from 'react-router-dom';
import Cookies from 'js-cookie';
import './Header.css';
import createIcon from './icons/create.png'; // Імпортуємо іконку

const Header = () => {
  const token = Cookies.get('token'); // Перевіряємо наявність токена

  const handleLogout = () => {
    Cookies.remove('token'); // Видаляємо токен при виході
    window.location.href = '/login'; // Перенаправляємо на сторінку входу
  };

  return (
    <header className="header-container">
      <nav className="header-nav">
        <div className="header-logo">
          <Link to="/">Introduction to web</Link>
        </div>
        <ul className="header-nav-links">
          <li><Link to="/dashboard">Dashboard</Link></li>
          {/* <li><Link to="/faq">FAQ</Link></li> */}
          {/* <li><Link to="/support">Support</Link></li> */}
        </ul>
        <div className="header-btns">
          {!token ? (
            <>
              <div className="header-signin-btn">
                <Link to="/login">Log In</Link>
              </div>
              <div className="header-download-btn">
                <Link to="/signup">Sign Up</Link>
              </div>
            </>
          ) : (
            <>
              <div className="header-profile-btn">
                <Link to="/profile">
                  <img src={createIcon} alt="Profile" />
                </Link>
              </div>
              <div className="header-logout-btn">
                <button onClick={handleLogout}>Log Out</button>
              </div>
            </>
          )}
        </div>
      </nav>
    </header>
  );
};

export default Header;
