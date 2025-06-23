import React, {useState} from 'react';
import cl from './Home.module.css';
import MenuItemLoader from "../../componends/MenuItemLoader/MenuItemLoader";

const Home = () => {
    const [products, setProducts] = useState([]);
    return (
        <div className={cl.main}>
            <h1 className={cl.title}>React-Go Shop</h1>
            <p className={cl.description}>
                Добро пожаловать в React-Go Shop — современный онлайн-магазин техники, аксессуаров и товаров для дома.
                Удобный интерфейс, быстрая доставка и качественный сервис.
            </p>
            <MenuItemLoader products={products} setProducts={setProducts}/>

        </div>
    );
};

export default Home;
