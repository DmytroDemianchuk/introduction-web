// App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Header from './Components/header/Header';
import Dashboard from './Components/dashboard/Dashboard';
import Login from './Components/login/LoginForm';
import SignUpForm from './Components/register/SignUpForm';
import LandingPage from './Components/landing/LandingPage';
import ProtectedRoute from './Components/protected/ProtectedRoute';
import Support from './Components/support/Support'; // Updated to correct path
import FAQ from './Components/faq/FAQ'; // Import the FAQ component

const App = () => {
  return (
    <Router>
      <Header />
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/support" element={<Support />} />
        <Route path="/faq" element={<FAQ />} /> {/* Add the FAQ route */}
        <Route path="/dashboard" element={
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        } />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<SignUpForm />} />
      </Routes>
    </Router>
  );
};

export default App;
