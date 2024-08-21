import React, { useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import './Profile.css'; // Імпортуйте ваш CSS файл

const ProfilePage = () => {
  const [profile, setProfile] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const token = Cookies.get('token');

        if (!token) {
          setError('No token found');
          return;
        }

        const response = await fetch('http://localhost:8080/profile', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${token}`
          }
        });

        if (response.ok) {
          const data = await response.json();
          console.log('Profile data:', data); // Для перевірки
          setProfile(data);
        } else {
          const errorText = await response.text();
          setError(`Failed to fetch profile: ${errorText}`);
        }
      } catch (error) {
        setError(`Failed to fetch profile: ${error.message}`);
      }
    };

    fetchProfile();
  }, []);

  if (error) return <div className="error">{error}</div>;
  if (!profile) return <div className="loading">Loading...</div>;

  return (
    <div className="container">
      <div className="profile-container">
        <div className="profile-header">
          <img src="/icons/profile.png" alt="Profile" /> {/* Виправлений шлях до зображення */}
          <h2>{profile.Name}</h2>
          <p>{profile.Email}</p>
        </div>
        {/* <div className="profile-details">
          <p><label>Name:</label> {profile.Name}</p>
          <p><label>Email:</label> {profile.Email}</p>
        </div> */}
      </div>
    </div>
  );
};

export default ProfilePage;
