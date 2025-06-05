// frontend/AuthContext.js
import React, { createContext, useState, useEffect } from 'react';
import { GoogleSignin } from '@react-native-google-signin/google-signin';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    GoogleSignin.configure({
      webClientId: '383142498918-arffvr6d3gntjl9u5veclurt5bqrcs4m.apps.googleusercontent.com', // keep same as backend verification
    });

    (async () => {
      const isSignedIn = await GoogleSignin.isSignedIn();
      if (isSignedIn) {
        const currentUser = await GoogleSignin.signInSilently();
        setUser(currentUser);
      }
    })();
  }, []);

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
};
