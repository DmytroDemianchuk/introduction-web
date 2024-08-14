import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

import Header from './Components/header/Header'; // Імпортуй новий Header
import Landing from './Components/landing/LandingPage'; // Імпортуй компонент LandingPage
import Login from './Components/login/Login';
import Register from './Components/register/SignUpForm';

import Footer from './Components/footer/Footer';

import './App.css';

function App() {
  return (
    <BrowserRouter>
      {/* Використовуй Header для заголовку та меню навігації */}
      <Header />
      <Routes>
        {/* Використовуй Landing як головну сторінку */}
        <Route path="/" element={<Landing />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
      </Routes>
      {/* Footer завжди знаходиться внизу сторінки */}
      <Footer />
    </BrowserRouter>
  );
}

export default App;
