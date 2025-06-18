import React from 'react';
import cl from './Navbar.module.css';
import {Link} from "react-router-dom";

const Navbar = () => {
    return (
        <div className={cl.navbar}>
            <Link to={"/"} className={cl.navbarItem} >Каталог</Link>
            <span className={cl.navbarItem}>Корзина</span>
            <span className={cl.navbarItem}>Аккаунт</span>
            <Link to={"/admin"} className={cl.navbarItem}>Админ-панель</Link>
        </div>
    );
};

export default Navbar;
