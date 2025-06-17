import React from 'react';
import cl from './ProductCard.module.css';
import productImage from '../../../img/XXL.jpeg'; // путь к изображению

const ProductCard = () => {
    return (
        <div className={cl.card}>
            <img src={productImage} alt="Костюм" className={cl.image} />
            <div className={cl.content}>
                <h3 className={cl.title}>чёрный классический костюм</h3>
                <p className={cl.price}>11 900 <span>₽</span></p>
                <p className={cl.subtitle}>чёрный костюм тройка</p>
                <button className={cl.button}>В корзину</button>
            </div>
        </div>
    );
};

export default ProductCard;