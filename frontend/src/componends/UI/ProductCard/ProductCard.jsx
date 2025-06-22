import React, { useContext } from 'react';
import cl from './ProductCard.module.css';
import { useLocation } from "react-router-dom";
import ProductService from "../../../api/ProductService";
import {CartContext} from "../../../context";

const ProductCard = ({ product, onDelete }) => {
    const location = useLocation();
    const { cart, setCart } = useContext(CartContext);

    const itemInCart = cart.find(item => item.productId === product.id);

    const hideDeleterRoutes = ["/"];
    const shouldHideDelete = hideDeleterRoutes.includes(location.pathname);

    const addCart = async (id) => {
        try {
            await ProductService.addItemCart(id);
            const updatedCart = [...cart, { productId: id, quantity: 1 }];
            setCart(updatedCart);
        } catch (e) {
            console.error(e);
        }
    };

    const increment = async (id) => {
        try {
            await ProductService.addItemCart(id);
            const updatedCart = [...cart, { productId: id, quantity: 1 }];
            setCart(updatedCart);
            setCart(cart.map(item =>
                item.productId === id
                    ? { ...item, quantity: item.quantity + 1 }
                    : item
            ));
        } catch (e) {
            console.error(e);
            alert("Ощибка")
        }
    };

    const decrement = (id) => {
        const updated = cart
            .map(item =>
                item.productId === id
                    ? { ...item, quantity: item.quantity - 1 }
                    : item
            )
            .filter(item => item.quantity > 0);

        setCart(updated);
    };

    return (
        <div className={cl.card}>
            <div className={cl.imageWrapper}>
                <img src={product.image} alt={product.title} className={cl.image} />
                {(!shouldHideDelete) && (
                    <button className={cl.deleteButton} onClick={onDelete}>
                        &times;
                    </button>
                )}
            </div>
            <div className={cl.content}>
                <h3 className={cl.title}>{product.title}</h3>
                <p className={cl.price}>
                    {product.price.toLocaleString()} <span>₽</span>
                </p>
                {itemInCart ? (
                    <div className={cl.counter}>
                        <button onClick={() => decrement(product.id)}>-</button>
                        <span>{itemInCart.quantity}</span>
                        <button onClick={() => increment(product.id)}>+</button>
                    </div>
                ) : (
                    <button className={cl.button} onClick={() => addCart(product.id)}>В корзину</button>
                )}
            </div>
        </div>
    );
};

export default ProductCard;
