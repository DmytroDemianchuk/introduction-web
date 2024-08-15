import React from 'react';
import './Header.css'; // Importing CSS for Header

const Header = () => {
  return (
    <header className="header-container"> {/* Unique class for Header */}
      <nav className="header-nav"> {/* Unique class for navigation */}
        <div className="header-logo">
          <a href="/">Introduction to web</a>
        </div>
        <ul className="header-nav-links"> {/* Unique class for navigation links */}
          <li><a href="#features">Features</a></li>
          <li><a href="#pricing">Pricing</a></li>
          <li><a href="#faq">FAQ</a></li>
          <li><a href="#support">Support</a></li>
        </ul>

        <div className='btn'> {/* Container for buttons */}
          <div className="header-signin-btn">
            <a href="/login">Log In</a>
          </div>
          <div className="header-download-btn">
            <a href="/register">Sign Up</a>
          </div>
        </div>
      </nav>
    </header>
  );
};

export default Header;
