import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Header from './Components/header/Header';
import Dashboard from './Components/dashboard/Dashboard';
import Login from './Components/login/LoginForm';
import SignUpForm from './Components/register/SignUpForm'; // Ensure this path is correct
import LandingPage from './Components/landing/LandingPage';

const App = () => {
  return (
    <Router>
      <Header />
      <Routes>
      <Route path="/" element={<LandingPage />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<SignUpForm />} /> {/* Updated to /register */}
      </Routes>
    </Router>
  );
};

export default App;
