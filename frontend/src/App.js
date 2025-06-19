import React from 'react';
import './styles/App.css'
import Navbar from "./componends/UI/Navbar/Navbar";
import {BrowserRouter} from "react-router-dom";
import AppRouter from "./componends/AppRouter";
import {AuthContext} from "./context";

export default function App() {
    const [isLoggedIn, setIsLoggedIn] = React.useState(false);
    const [isAdmin, setIsAdmin] = React.useState(true);

    return (
      <AuthContext.Provider value={{
          isLoggedIn,
          setIsLoggedIn,
          isAdmin,
          setIsAdmin,
      }}>
          <BrowserRouter>
              <Navbar />
              <AppRouter/>
          </BrowserRouter>
      </AuthContext.Provider>
    );
}

