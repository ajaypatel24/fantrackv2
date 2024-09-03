import React from 'react';

const AuthButton = () => {
  const handleLogin = () => {
    window.location.href = 'http://localhost:8080/login'; // Redirect to your backend
  };

  return <button onClick={handleLogin}>Log in with Yahoo</button>;
};

export default AuthButton;