import React from 'react';
import cl from './ProductCard.module.css';

const ProductCard = ({ title, price, description, image, onDelete }) => {
    return (
        <div className={cl.card}>
            <div className={cl.imageWrapper}>
                <img src={image} alt={title} className={cl.image} />
                {onDelete && (
                    <button className={cl.deleteButton} onClick={onDelete}>
                        &times;
                    </button>
                )}
            </div>
            <div className={cl.content}>
                <h3 className={cl.title}>{title}</h3>
                <p className={cl.price}>
                    {price.toLocaleString()} <span>₽</span>
                </p>
                <p className={cl.subtitle}>{description}</p>
                <button className={cl.button}>В корзину</button>
            </div>
        </div>
    );
};

export default ProductCard;
