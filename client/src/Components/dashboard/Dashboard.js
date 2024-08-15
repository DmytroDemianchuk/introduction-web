import React, { useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import './Dashboard.css';

const Dashboard = () => {
  const [books, setBooks] = useState([]);
  const [bookId, setBookId] = useState('');
  const [bookTitle, setBookTitle] = useState('');
  const [bookAuthor, setBookAuthor] = useState('');
  const [bookPublishedYear, setBookPublishedYear] = useState('');
  const [message, setMessage] = useState('');
  const [showBooks, setShowBooks] = useState(false);

  const token = Cookies.get('token'); // Example of getting token from cookies

  useEffect(() => {
    if (showBooks) {
      fetchBooks();
    }
  }, [showBooks]);

  const fetchBooks = async () => {
    try {
      const response = await fetch('http://localhost:8080/books');
      const data = await response.json();
      setBooks(data);
    } catch (error) {
      setMessage('Failed to fetch books');
    }
  };

  const handleCreateBook = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/books/create', {
        method: 'POST',
        headers: { 
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}` // Sending token in the header
        },
        body: JSON.stringify({
          title: bookTitle,
          author: bookAuthor,
          publishedYear: bookPublishedYear,
        }),
      });
      if (response.ok) {
        setMessage('Book created successfully');
        fetchBooks(); // Refresh the list of books
      } else {
        setMessage('Failed to create book');
      }
    } catch (error) {
      setMessage('Error creating book');
    }
  };

  const handleGetBookByID = async () => {
    try {
      const response = await fetch(`http://localhost:8080/books?id=${bookId}`);
      const data = await response.json();
      // Handle the book data (e.g., display it)
      setBookTitle(data.title);
      setBookAuthor(data.author);
      setBookPublishedYear(data.publishedYear);
    } catch (error) {
      setMessage('Failed to fetch book');
    }
  };

  const handleUpdateBook = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/books/update', {
        method: 'PUT',
        headers: { 
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}` // Sending token in the header
        },
        body: JSON.stringify({
          id: bookId,
          title: bookTitle,
          author: bookAuthor,
          publishedYear: bookPublishedYear,
        }),
      });
      if (response.ok) {
        setMessage('Book updated successfully');
        fetchBooks(); // Refresh the list of books
      } else {
        setMessage('Failed to update book');
      }
    } catch (error) {
      setMessage('Error updating book');
    }
  };

  const handleDeleteBook = async () => {
    try {
      const response = await fetch(`http://localhost:8080/books/delete?id=${bookId}`, {
        method: 'DELETE',
        headers: { 
          'Authorization': `Bearer ${token}` // Sending token in the header
        },
      });
      if (response.ok) {
        setMessage('Book deleted successfully');
        fetchBooks(); // Refresh the list of books
      } else {
        setMessage('Failed to delete book');
      }
    } catch (error) {
      setMessage('Error deleting book');
    }
  };

  return (
    <div className="dashboard-container">
      <h1>Welcome to Your Dashboard</h1>
      <p>This is your personal dashboard. From here you can manage your books.</p>

      <div className="dashboard-section">
        <h2>Create a Book</h2>
        <form onSubmit={handleCreateBook} className="dashboard-form">
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
          <button type="submit">Create Book</button>
        </form>
      </div>

      <div className="dashboard-section">
        <h2>Get Book by ID</h2>
        <input
          type="text"
          placeholder="Book ID"
          value={bookId}
          onChange={(e) => setBookId(e.target.value)}
        />
        <button onClick={handleGetBookByID}>Get Book</button>
      </div>
      
      <div className="dashboard-section">
        <h2>Update a Book</h2>
        <form onSubmit={handleUpdateBook} className="dashboard-form">
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
          <button type="submit">Update Book</button>
        </form>
      </div>

      <div className="dashboard-section">
        <h2>Delete a Book</h2>
        <input
          type="text"
          placeholder="Book ID"
          value={bookId}
          onChange={(e) => setBookId(e.target.value)}
        />
        <button onClick={handleDeleteBook}>Delete Book</button>
      </div>

      <div className="dashboard-section">
        <h2>All Books</h2>
        <button onClick={() => setShowBooks(!showBooks)}>
          {showBooks ? 'Hide' : 'Show'} Books
        </button>
        {showBooks && (
          <ul>
            {books.map(book => (
              <li key={book.id}>{book.title} by {book.author} ({book.publishedYear})</li>
            ))}
          </ul>
        )}
      </div>

      {message && <p className="dashboard-message">{message}</p>}
    </div>
  );
};

export default Dashboard;
