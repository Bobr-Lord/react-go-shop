import React, {useEffect} from 'react';
import './styles/App.css'
import Navbar from "./componends/UI/Navbar/Navbar";
import {BrowserRouter, useLocation} from "react-router-dom";
import AppRouter from "./componends/AppRouter";
import {AuthContext} from "./context";


export default function App() {
    const [isLoggedIn, setIsLoggedIn] = React.useState(false);
    const [isAdmin, setIsAdmin] = React.useState(true);

    const location = useLocation();
    const hideNavbarRoutes = ["/login", "/register"];
    const shouldHideNavbar = hideNavbarRoutes.includes(location.pathname);

    useEffect(() => {
        const login = localStorage.getItem('isLoggedIn');
        const role = localStorage.getItem('role');

        if (login === 'true') {
            setIsLoggedIn(true);
        }
        if (role === 'admin') {
            setIsAdmin(true);
        }
    }, []);

    return (
      <AuthContext.Provider value={{
          isLoggedIn,
          setIsLoggedIn,
          isAdmin,
          setIsAdmin,
      }}>
          {!shouldHideNavbar && <Navbar />}
          <AppRouter/>
      </AuthContext.Provider>
    );
}

