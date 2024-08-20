import React, { useState } from 'react';
import './BookList.css';

const BookList = ({ books, onDeleteSelected, onEditBook }) => {
  const [selectedBooks, setSelectedBooks] = useState([]);

  const toggleSelectBook = (bookId) => {
    setSelectedBooks((prevSelected) =>
      prevSelected.includes(bookId)
        ? prevSelected.filter((id) => id !== bookId)
        : [...prevSelected, bookId]
    );
  };

  return (
    <div className="book-list">
      {books.length > 0 ? (
        books.map((book) => (
          <div key={book.id} className="book-list-item">
            <input
              type="checkbox"
              checked={selectedBooks.includes(book.id)}
              onChange={() => toggleSelectBook(book.id)}
            />
            <span>{book.title} by {book.author}</span>
            <button onClick={() => onEditBook(book)} className="edit-btn">
              Edit
            </button>
          </div>
        ))
      ) : (
        <p>No books available.</p>
      )}
      <button onClick={() => onDeleteSelected(selectedBooks)} className="delete-btn">
        Delete Selected
      </button>
    </div>
  );
};

export default BookList;
