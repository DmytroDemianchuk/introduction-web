import React, { useState } from 'react';
import './SignUpForm.css';

const SignUpForm = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission here
    console.log('Name:', name);
    console.log('Email:', email);
    console.log('Password:', password);
  };

  return (
    <div className="signup-parent-container">
      <div className="signup-container">
        <h2>Sign Up</h2>
        <p>Let's get started with your 30 days free trial</p>
        <form onSubmit={handleSubmit}>
          <input 
            type="text" 
            placeholder="Name" 
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <input 
            type="email" 
            placeholder="Email" 
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <input 
            type="password" 
            placeholder="Password" 
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit">Sign Up</button>
        </form>
        <p className="login-link">
          Already have an account? <a href="/login">Log In</a>
        </p>
        <div className="or-divider">or</div>
        <button className="google-signup-btn">Sign up with Google</button>
        <p className="terms">
          By signing up you accept the Company's <a href="/terms">Terms of Use</a> and <a href="/privacy">Privacy Policy</a>.
        </p>
      </div>
    </div>
  );
};

export default SignUpForm;
