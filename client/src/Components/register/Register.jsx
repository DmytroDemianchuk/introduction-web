import React, { useState } from 'react';
import axios from 'axios';
import "./Register.css";

function Register() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
    phone: '',
    isInvalid: false,
    message: '',
  });

  const { name, email, password, phone, isInvalid, message } = formData;

  const onChange = e => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const onSubmit = async e => {
    e.preventDefault();

    try {
      const res = await axios.post('http://localhost:8000/create-name', {
        name,
        email,
        password,
        phone,
      });

      console.log('register', res);
      if (res.data.status) {
        // Handle success response here
        console.log('Success:', res.data);
      } else {
        // Handle failed response here
        setFormData({ ...formData, message: res.data.message, isInvalid: true });
      }
    } catch (error) {
      console.error('Error:', error);
      setFormData({ ...formData, message: 'Ця електронна пошта вже зайнята', isInvalid: true });
    }
  };

  return (
    <div className="signup_form">
      <div className="sign_form-body">
        <h2 className="head-name">Get Start Now</h2>
        <form onSubmit={onSubmit}>

          <div className="signup_form-head">
            <label htmlFor="request-title">Name</label>
            <input
              className="input"
              name="name"
              id="request-title"
              type="text"
              placeholder="Enter your name"
              value={name}
              onChange={onChange}
            />
            {/* <FormErrorMessage v-if="isInvalid">{message}</FormErrorMessage> */}
          </div>

          <div>
            <label htmlFor="email">Email address</label>
            <input
              name="email"
              id="request-email"
              className="input"
              type="text"
              placeholder="Enter your email"
              value={email}
              onChange={onChange}
            />
            {isInvalid && <span className="error-message">{message}</span>}
          </div>

          <div>
            <label htmlFor="request-email">Password</label>
            <input
              name="password"
              id="request-email"
              className="input"
              type="password"
              placeholder="Enter your password"
              value={password}
              onChange={onChange}
            />
          </div>

          <div className="signup_agreement">
            <label className="signup_for-label">
              <input className="label_controll" type="checkbox" checked />
              <i className="checkbox">I agree to the terms & policy</i>
            </label>
          </div>

          <button className="button_submit" type="submit" id="submitButton">
            <span className="button_content">Зареєструватись</span>
          </button>

          <button className="button_submit" type="submit" id="submitButton">
            <a className="button_content" href="signin">
              Увійти
            </a>
          </button>
        </form>
      </div>
    </div>
  );
}

export default Register;
