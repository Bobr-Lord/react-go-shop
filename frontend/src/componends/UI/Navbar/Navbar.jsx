import React from 'react';
import cl from './Navbar.module.css';
import {Link} from "react-router-dom";
import {AuthContext} from "../../../context";

const Navbar = () => {
    const {isAdmin} = React.useContext(AuthContext);
    return (
        <div className={cl.navbar}>
            <Link to={"/"} className={cl.navbarItem} >Каталог</Link>
            <span className={cl.navbarItem}>Корзина</span>
            <span className={cl.navbarItem}>Аккаунт</span>
            {
                isAdmin && <Link to={"/admin"} className={cl.navbarItem}>Админ-панель</Link>
            }
        </div>
    );
};

export default Navbar;
