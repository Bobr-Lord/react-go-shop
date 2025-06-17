import React from 'react';
import cl from './Navbar.module.css';

const Navbar = () => {
    return (
        <div className={cl.navbar}>
            <span className={cl.navbarItem}>Каталог</span>
            <span className={cl.navbarItem}>Корзина</span>
            <span className={cl.navbarItem}>Аккаунт</span>
        </div>
    );
};

export default Navbar;
