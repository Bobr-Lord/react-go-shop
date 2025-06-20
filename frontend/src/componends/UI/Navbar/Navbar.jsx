import React from 'react';
import cl from './Navbar.module.css';
import { Link } from "react-router-dom";
import { AuthContext } from "../../../context";

const Navbar = () => {
    const { isAdmin, isLoggedIn, user } = React.useContext(AuthContext);

    return (
        <div className={cl.navbar}>
            <div className={cl.leftSection}>
                <Link to={"/"} className={cl.navbarItem}>Каталог</Link>
                <Link to={"/cart"} className={cl.navbarItem}>Корзина</Link>
                <span className={cl.navbarItem}>Аккаунт</span>
                {isAdmin && (
                    <Link to={"/admin"} className={cl.navbarItem}>Админ-панель</Link>
                )}
            </div>
            {
                (isLoggedIn && user !== null)
                ? <span className={cl.navbarItem}>{user.first_name}</span>
                : <Link to={"/login"} className={cl.navbarItem}>Войти</Link>
            }
        </div>
    );
};

export default Navbar;
