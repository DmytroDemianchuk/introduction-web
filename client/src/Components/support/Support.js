import React, { useState } from 'react';
import './Support.css';

const Support = () => {
  const [copySuccess, setCopySuccess] = useState('');

  const copyToClipboard = () => {
    const address = '0x1234567890abcdef1234567890abcdef12345678';
    navigator.clipboard.writeText(address)
      .then(() => {
        setCopySuccess('Address copied to clipboard!');
      })
      .catch(() => {
        setCopySuccess('Failed to copy address.');
      });
  };

  return (
    <div className="support-parent-container">
      <div className="support-container">
        <h1>Support Us</h1>
        <p>If you find our work helpful, you can support us by donating to the following crypto address:</p>
        <div className="crypto-address">
          <h2>Crypto Address</h2>
          <div className="address-container">
            <div className="address">
              0x1234567890abcdef1234567890abcdef12345678
            </div>
            <button className="copy-button" onClick={copyToClipboard}>Copy Address</button>
          </div>
          {copySuccess && <p className="message success">{copySuccess}</p>}
        </div>
        <p>Your support helps us continue our work and improve our services. Thank you!</p>
        <a href="https://example.com/donate" target="_blank" rel="noopener noreferrer">Learn more about how you can help</a>
      </div>
    </div>
  );
};

export default Support;
