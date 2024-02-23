import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

import Header from './Components/header/Header';
import Landing from './Components/landing/Landing';
import Login from './Components/login/Login';
import Chat from './Components/chat/Chat';
import Register from './Components/register/Register';

import Footer from './Components/footer/Footer';

import theme from './theme';
import './App.css';

function App() {
  return (

        <BrowserRouter>
          <Header></Header>
          <Routes>
            <Route path="/" element={<Landing />} />
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route path="/chat" element={<Chat />} />
          </Routes>
          <Footer></Footer>
        </BrowserRouter>
  );
}

export default App;
