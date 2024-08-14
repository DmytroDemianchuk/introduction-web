import React from 'react';
import '../landing/LandingPage.css'; // Шлях до CSS файлу
import LandingPage from '../landing/LandingPage'; // Шлях до компонента LandingPage

const Header = () => {
  return (
    <div>
      {/* Navbar */}
      <nav>
        <div className="logo">
          <a href="/">YourLogo</a>
        </div>
        <ul className="nav-links">
          <li><a href="#features">Features</a></li>
          <li><a href="#pricing">Pricing</a></li>
          <li><a href="#faq">FAQ</a></li>
          <li><a href="#support">Support</a></li>
        </ul>
        <div className="download-btn">
          <a href="/register">Download</a>
        </div>
      </nav>
      
      {/* Landing Page Section */}
    </div>
  );
};

export default Header;
