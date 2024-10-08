// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './Components/header/Header';
import Dashboard from './Components/dashboard/Dashboard';
import Login from './Components/login/LoginForm';
import SignUpForm from './Components/register/SignUpForm';
import LandingPage from './Components/landing/LandingPage';
import ProtectedRoute from './Components/protected/ProtectedRoute';
import Support from './Components/support/Support';
import FAQ from './Components/faq/FAQ';
import Footer from './Components/footer/Footer';

const App = () => {
  return (
    <Router>
      <Header />
      <div className="main-content">
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/support" element={<Support />} />
          <Route path="/faq" element={<FAQ />} />
          <Route path="/dashboard" element={
            <ProtectedRoute>
              <Dashboard />
            </ProtectedRoute>
          } />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUpForm />} />
        </Routes>
      </div>
      <Footer />
    </Router>
  );
};

export default App;
