import React, {useEffect} from 'react';
import './styles/App.css'
import Navbar from "./componends/UI/Navbar/Navbar";
import {useLocation} from "react-router-dom";
import AppRouter from "./componends/AppRouter";
import {AuthContext} from "./context";
import ProductService from "./api/ProductService";


export default function App() {
    const [isLoggedIn, setIsLoggedIn] = React.useState(false);
    const [isAdmin, setIsAdmin] = React.useState(false);

    const location = useLocation();
    const hideNavbarRoutes = ["/login", "/register"];
    const shouldHideNavbar = hideNavbarRoutes.includes(location.pathname);

    useEffect( () => {
        ProductService.getMe().then((res) => {
            console.log(res);
            if (res.data.role === "admin") {
                setIsAdmin(true);
            }
        }).catch(err=>{
            console.log(err)
        })
    }, [isLoggedIn]);

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

