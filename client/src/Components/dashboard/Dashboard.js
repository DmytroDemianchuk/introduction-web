import React, { useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import './Dashboard.css';
import createIcon from '../../icons/create.png';
import deleteIcon from '../../icons/delete.png';
import viewIcon from '../../icons/view.png';
import editIcon from '../../icons/update.png';

const Dashboard = () => {
  const [books, setBooks] = useState([]);
  const [bookId, setBookId] = useState('');
  const [bookTitle, setBookTitle] = useState('');
  const [bookAuthor, setBookAuthor] = useState('');
  const [bookPublishedYear, setBookPublishedYear] = useState('');
  const [message, setMessage] = useState('');
  const [showBooks, setShowBooks] = useState(false);
  const [selectedAction, setSelectedAction] = useState(null);

  const token = Cookies.get('token');

  useEffect(() => {
    if (showBooks) {
      fetchBooks();
    }
  }, [showBooks]);

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
      console.log('Books fetched:', data); // Log books data
    } catch (error) {
      console.error('Error fetching books:', error); // Log error
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
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          title: bookTitle,
          author: bookAuthor,
          publishedYear: bookPublishedYear,
        }),
      });
      const data = await response.json();
      console.log('Create response:', data); // Логування відповіді
      if (response.ok) {
        setMessage('Book created successfully');
        fetchBooks();
      } else {
        setMessage('Failed to create book');
      }
    } catch (error) {
      setMessage('Error creating book');
    }
  };

  const handleUpdateBook = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/books/update', {
        method: 'PUT',
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
      const data = await response.json();
      console.log('Update response:', data); // Логування відповіді
      if (response.ok) {
        setMessage('Book updated successfully');
        fetchBooks();
      } else {
        setMessage('Failed to update book');
      }
    } catch (error) {
      setMessage('Error updating book');
    }
  };

  const handleDeleteBook = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/books/delete?id=${id}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Cache-Control': 'no-cache',
          'Pragma': 'no-cache',
        },
      });

      if (response.ok) {
        console.log('Delete response:', await response.text()); // Логування відповіді
        setMessage('Book deleted successfully');
        // Видалення книги з локального стану
        setBooks(prevBooks => prevBooks.filter(book => book.id !== id));
      } else {
        setMessage('Failed to delete book');
      }
    } catch (error) {
      console.error('Error deleting book:', error); // Логування помилки
      setMessage('Error deleting book');
    }
  };

  const handleEditClick = (book) => {
    setBookId(book.id);
    setBookTitle(book.title);
    setBookAuthor(book.author);
    setBookPublishedYear(book.publishedYear);
    setSelectedAction('update');
  };

  const renderAction = () => {
    switch (selectedAction) {
      case 'create':
        return (
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
        );
      case 'update':
        return (
          <div className="dashboard-section">
            <h2>Update a Book</h2>
            <form onSubmit={handleUpdateBook} className="dashboard-form">
              <input
                type="text"
                placeholder="Book ID"
                value={bookId}
                onChange={(e) => setBookId(e.target.value)}
                required
              />
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
        );
      case 'delete':
        return (
          <div className="dashboard-section">
            <h2>Delete a Book</h2>
            <input
              type="text"
              placeholder="Book ID"
              value={bookId}
              onChange={(e) => setBookId(e.target.value)}
            />
            <button onClick={() => handleDeleteBook(bookId)}>Delete Book</button>
          </div>
        );
      case 'view':
        return (
          <div className="dashboard-section">
            <h2>All Books</h2>
            <button onClick={() => setShowBooks(!showBooks)}>
              {showBooks ? 'Hide' : 'Show'} Books
            </button>
            {showBooks && (
              <ul>
                {books.map(book => (
                  <li key={book.id}>
                    {book.title} by {book.author} ({book.publishedYear})
                    <button 
                      onClick={() => handleEditClick(book)} 
                      style={{ marginLeft: '10px', cursor: 'pointer', color: 'blue' }}
                    >
                      <img src={editIcon} alt="Edit" style={{ width: '20px', height: '20px' }} />
                    </button>
                    <button 
                      onClick={() => handleDeleteBook(book.id)} 
                      style={{ marginLeft: '10px', cursor: 'pointer', color: 'red' }}
                    >
                      ❌
                    </button>
                  </li>
                ))}
              </ul>
            )}
          </div>
        );
      default:
        return <p>Please select an action above to continue.</p>;
    }
  };

  return (
    <div className="dashboard-container">
      <h1>Welcome to Your Dashboard</h1>
      <p>This is your personal dashboard. From here you can manage your books.</p>

      <div className="options-container">
        <div className="option-card" onClick={() => setSelectedAction('create')}>
          <img src={createIcon} alt="Create" />
          <h2>Create</h2>
        </div>
        <div className="option-card" onClick={() => setSelectedAction('update')}>
          <img src={editIcon} alt="Update" />
          <h2>Update</h2>
        </div>
        <div className="option-card" onClick={() => setSelectedAction('delete')}>
          <img src={deleteIcon} alt="Delete" />
          <h2>Delete</h2>
        </div>
        <div className="option-card" onClick={() => setSelectedAction('view')}>
          <img src={viewIcon} alt="View All" />
          <h2>View All</h2>
        </div>
      </div>

      <div className="action-container">
        {renderAction()}
      </div>

      {message && <p>{message}</p>}
    </div>
  );
};

export default Dashboard;
