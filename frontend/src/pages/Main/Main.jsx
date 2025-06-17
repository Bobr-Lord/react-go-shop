import React from 'react';
import cl from './Main.module.css';
import ProductCard from "../../componends/UI/ProductCard/ProductCard";

const Main = () => {
    return (
        <div className={cl.main}>
            <h1 className={cl.title}>React-Go Shop</h1>
            <p className={cl.description}>
                Добро пожаловать в React-Go Shop — современный онлайн-магазин техники, аксессуаров и товаров для дома.
                Удобный интерфейс, быстрая доставка и качественный сервис.
            </p>
            <ProductCard/>
        </div>
    );
};

export default Main;
