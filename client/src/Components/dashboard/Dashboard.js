import React, { useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import './Dashboard.css';
import BookForm from '../bookForm/BookForm';
import BookList from '../bookList/BookList';

const Dashboard = () => {
  const [books, setBooks] = useState([]);
  const [bookId, setBookId] = useState('');
  const [bookTitle, setBookTitle] = useState('');
  const [bookAuthor, setBookAuthor] = useState('');
  const [bookPublishedYear, setBookPublishedYear] = useState('');
  const [message, setMessage] = useState('');
  const [selectedAction, setSelectedAction] = useState('create');
  const [showAllBooks, setShowAllBooks] = useState(false);

  const token = Cookies.get('token');
  const booksPerPage = 5;

  useEffect(() => {
    fetchBooks();
  }, []);

  const fetchBooks = async () => {
    try {
      const response = await fetch('http://localhost:8080/books', {
        headers: {
          'Cache-Control': 'no-cache',
          'Pragma': 'no-cache',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      setBooks(data);
    } catch (error) {
      setMessage('Failed to fetch books');
    }
  };

  const handleCreateOrUpdateBook = async (e) => {
    e.preventDefault();
    const url = selectedAction === 'create'
      ? 'http://localhost:8080/books/create'
      : 'http://localhost:8080/books/update';
    const method = selectedAction === 'create' ? 'POST' : 'PUT';

    try {
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          id: bookId,
          title: bookTitle,
          author: bookAuthor,
          publishedYear: bookPublishedYear,
        }),
      });

      if (response.ok) {
        setMessage(selectedAction === 'create' ? 'Book created successfully' : 'Book updated successfully');
        fetchBooks(); // Оновлюємо список книг після успішного редагування або створення
        resetForm(); // Скидаємо форму
      } else {
        setMessage('Failed to save book');
      }
    } catch (error) {
      setMessage('Error saving book');
    }
  };

  const handleDeleteSelectedBooks = async (selectedBookIds) => {
    try {
      await Promise.all(
        selectedBookIds.map((id) =>
          fetch(`http://localhost:8080/books/delete?id=${id}`, {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${token}`,
            },
          })
        )
      );
      setMessage('Selected books deleted successfully');
      setBooks(books.filter((book) => !selectedBookIds.includes(book.id))); // Видаляємо зі списку
    } catch (error) {
      setMessage('Error deleting selected books');
    }
  };

  const handleEditBook = (book) => {
    setBookId(book.id);
    setBookTitle(book.title);
    setBookAuthor(book.author);
    setBookPublishedYear(book.publishedYear);
    setSelectedAction('update');
  };

  const resetForm = () => {
    setBookId('');
    setBookTitle('');
    setBookAuthor('');
    setBookPublishedYear('');
    setSelectedAction('create');
  };

  const toggleShowAllBooks = () => {
    setShowAllBooks(!showAllBooks);
  };

  const displayedBooks = showAllBooks ? books : books.slice(0, booksPerPage);

  return (
    <div className="dashboard">
      <h1>Dashboard</h1>

      <div className="actions-container">
        <button className="action-btn" onClick={() => setSelectedAction('create')}>
          Create Book
        </button>
        <button className="action-btn" onClick={() => setSelectedAction('view')}>
          View Books
        </button>
      </div>

      {(selectedAction === 'create' || selectedAction === 'update') && (
        <BookForm
          onSubmit={handleCreateOrUpdateBook}
          bookId={bookId}
          setBookId={setBookId}
          bookTitle={bookTitle}
          setBookTitle={setBookTitle}
          bookAuthor={bookAuthor}
          setBookAuthor={setBookAuthor}
          bookPublishedYear={bookPublishedYear}
          setBookPublishedYear={setBookPublishedYear}
          action={selectedAction}
        />
      )}

      {selectedAction === 'view' && (
        <div>
          <BookList
            books={displayedBooks}
            onDeleteSelected={handleDeleteSelectedBooks}
            onEditBook={handleEditBook}
          />
          <button onClick={toggleShowAllBooks} className="view-more-btn">
            {showAllBooks ? 'Show Less' : 'Show All'}
          </button>
        </div>
      )}

      {message && <p className="message">{message}</p>}
    </div>
  );
};

export default Dashboard;
