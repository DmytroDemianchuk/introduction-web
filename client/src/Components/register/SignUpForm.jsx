import React, { useState } from 'react';
import axios from 'axios';
import './SignUpForm.css';

const SignUpForm = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [isError, setIsError] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const requestBody = { name, email, password };

    try {
      const response = await axios.post('http://localhost:8080/signup', requestBody);

      setMessage('Signup successful!');
      setIsError(false);
    } catch (error) {
      console.error('Error:', error.response ? error.response.data : error.message);
      setMessage(error.response ? error.response.data.message : 'Something went wrong');
      setIsError(true);
    }
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
        {message && (
          <p className={isError ? 'error-message' : 'success-message'}>{message}</p>
        )}
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
