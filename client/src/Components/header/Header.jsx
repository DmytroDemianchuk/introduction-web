import React, { useState } from 'react';

import "./header.css";

function Header() {
    return (
      <div className='header'>
        <div>
          <div to="/">
            <div className='headet_text'>
               <a href="/">Introduction To the Web</a>
            </div>
          </div>
        </div>
      </div>
    );
  }

export default Header;