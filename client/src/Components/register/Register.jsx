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
        <h2 className="head-name">Реєстрація</h2>
        <form onSubmit={onSubmit}>
          <div className="signup_form-head">
            <label htmlFor="request-title">Ім'я та прізвище</label>
            <input
              className="input"
              name="name"
              id="request-title"
              type="text"
              placeholder="type your username"
              value={name}
              onChange={onChange}
            />
            {/* <FormErrorMessage v-if="isInvalid">{message}</FormErrorMessage> */}
          </div>

          <div>
            <label htmlFor="email">Електронна пошта</label>
            <input
              name="email"
              id="request-email"
              className="input"
              type="text"
              value={email}
              onChange={onChange}
            />
            {isInvalid && <span className="error-message">{message}</span>}
          </div>

          <div>
            <label htmlFor="request-email">Пароль</label>
            <input
              name="password"
              id="request-email"
              className="input"
              type="password"
              value={password}
              onChange={onChange}
            />
          </div>

          <div>
            <label htmlFor="request-phone">Телефон</label>
            <input
              id="request-phone"
              className="input"
              type="text"
              value={phone}
              onChange={onChange}
              required
              autoComplete="off"
              placeholder="980 123 540"
            />
          </div>

          <div className="signup_agreement">
            <label className="signup_for-label">
              <input className="label_controll" type="checkbox" checked />
              <i className="checkbox"></i>
              <span className="label_text">
                Підтверджую свою згоду з умовами
                <br />
                <a href="/license" target="_blank">
                  ліцензійного договору
                </a>
              </span>
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
