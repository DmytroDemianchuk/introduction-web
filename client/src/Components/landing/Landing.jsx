import React from 'react';

import './Landing.css'; // Import the CSS file

function Landing ()  {
  return (
    <div className="container">
      <div className="text">
        This is a simple application of registerations 
      </div>
      <div>
        <div className="buttons" direction="row" spacing={7}>
            <div
              className="button register"
            >
              <a href="/register">Register</a>
            </div>

            <div className="button login" >
              <a href="/login">Login</a> 
            </div>
        </div>
      </div>
    </div>
  );
}

export default Landing;
