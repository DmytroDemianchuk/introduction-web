import React from 'react';
import './BookForm.css';

const BookForm = ({
  onSubmit,
  bookId = '',
  setBookId,
  bookTitle = '',
  setBookTitle,
  bookAuthor = '',
  setBookAuthor,
  bookPublishedYear = '',
  setBookPublishedYear,
  action
}) => {
  return (
    <div className="dashboard-section">
      <h2>{action === 'create' ? 'Create a Book' : 'Update a Book'}</h2>
      <form onSubmit={onSubmit} className="dashboard-form">
        {action === 'update' && (
          <input
            type="text"
            placeholder="Book ID"
            value={bookId}
            onChange={(e) => setBookId(e.target.value)}
            required
          />
        )}
        <input
          type="text"
          placeholder="Title"
          value={bookTitle}
          onChange={(e) => setBookTitle(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Author"
          value={bookAuthor}
          onChange={(e) => setBookAuthor(e.target.value)}
          required
        />
        <input
          type="number"
          placeholder="Published Year"
          value={bookPublishedYear}
          onChange={(e) => setBookPublishedYear(e.target.value)}
          required
        />
        <button type="submit">
          {action === 'create' ? 'Create Book' : 'Update Book'}
        </button>
      </form>
    </div>
  );
};

export default BookForm;
