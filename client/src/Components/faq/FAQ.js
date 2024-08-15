// FAQ.js
import React from 'react';
import './FAQ.css';

const FAQ = () => {
  const faqs = [
    {
      question: 'What is this website about?',
      answer: 'This website provides resources and tools for managing your personal library of books and more.'
    },
    {
      question: 'How do I create an account?',
      answer: 'You can create an account by clicking on the "Sign Up" button on the login page and filling out the required information.'
    },
    {
      question: 'How do I reset my password?',
      answer: 'If you have forgotten your password, click on the "Forgot Password" link on the login page to reset it via email.'
    },
    {
      question: 'How do I contact support?',
      answer: 'You can contact support through the "Support" page where you will find our contact information and other ways to get in touch with us.'
    },
    {
      question: 'Where can I find more information?',
      answer: 'Additional information can be found on our landing page or by exploring other sections of our website.'
    }
  ];

  return (
    <div className="faq-container">
      <h1>Frequently Asked Questions</h1>
      <div className="faq-list">
        {faqs.map((faq, index) => (
          <div key={index} className="faq-card">
            <h2 className="faq-question">{faq.question}</h2>
            <p className="faq-answer">{faq.answer}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default FAQ;
