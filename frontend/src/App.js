import React, { useEffect, useState } from 'react';
import './styles/App.css';
import Navbar from "./componends/UI/Navbar/Navbar";
import { useLocation } from "react-router-dom";
import AppRouter from "./componends/AppRouter";
import { AuthContext } from "./context";
import AuthService from "./api/AuthService";

export default function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [user, setUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true); // 👈

    const location = useLocation();
    const hideNavbarRoutes = ["/login", "/register"];
    const shouldHideNavbar = hideNavbarRoutes.includes(location.pathname);

    useEffect(() => {
        AuthService.getMe()
            .then((res) => {
                if (res.data.role === "admin") {
                    setIsAdmin(true);
                    setIsLoggedIn(true);
                    setUser(res.data);
                } else if (res.data.role === "user") {
                    setIsAdmin(false);
                    setIsLoggedIn(true);
                    setUser(res.data);
                }
            })
            .catch(err => {
                console.log("Not authenticated", err);
                setIsLoggedIn(false);
                setIsAdmin(false);
                setUser(null);
            })
            .finally(() => {
                setIsLoading(false); // 👈 теперь можно отрисовывать
            });
    }, []);

    if (isLoading) {
        return <div style={{ color: "#fff", textAlign: "center", marginTop: "100px" }}>Загрузка...</div>;
    }

    return (
        <AuthContext.Provider value={{
            isLoggedIn,
            setIsLoggedIn,
            isAdmin,
            setIsAdmin,
            user,
            setUser,
        }}>
            {!shouldHideNavbar && <Navbar />}
            <AppRouter />
        </AuthContext.Provider>
    );
}
