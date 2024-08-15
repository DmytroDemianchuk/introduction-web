import React from 'react';
import './LandingPage.css';

const LandingPage = () => {
  return (
    <div className="landing-parent-container">
      <div className="landing-container">
        <h1>Register Your Account</h1>
        <p>Quick and easy registration to get started. Follow the simple steps to create your account and gain access to all features.</p>
        <a href="/register" className="landing-cta-button">Sign Up</a>
      </div>
    </div>
  );
};

export default LandingPage;
