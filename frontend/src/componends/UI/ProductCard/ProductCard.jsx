import React, { useEffect, useState} from 'react';
import cl from './ProductCard.module.css';
import { useLocation } from "react-router-dom";
import ProductService from "../../../api/ProductService";

const ProductCard = ({ idProduct, products, setProducts, onDelete }) => {
    const location = useLocation();

    const [itemInCart, setItemInCart] = useState(products.find(item => item.id === idProduct));


    const hideDeleterRoutes = ["/"];
    const shouldHideDelete = hideDeleterRoutes.includes(location.pathname);

    useEffect(() => {
        setItemInCart(products.find(item => item.id === idProduct));
    }, [products])

    const addCart = async (item) => {
        try {
            await ProductService.addItemCart(item.id);

            const exists = products.some(p => p.id === item.id);

            if (exists) {
                const updatedProducts = products.map(p =>
                    p.id === item.id ? { ...p, quantity: 1 } : p
                );
                setProducts(updatedProducts);
            } else {
                setProducts([...products, { ...item, quantity: 1 }]);
            }
        } catch (e) {
            if (e.response?.status === 401) {
                alert("Не авторизованы");
            } else {
                console.error(e);
            }
        }
    };



    const increment = async (id) => {
        try {
            await ProductService.addItemCart(id);
            const updated = products
                .map(item =>
                    item.id === id
                        ? { ...item, quantity: item.quantity + 1 }
                        : item
                )
            setProducts(updated);
        } catch (e) {
            if (e.status === 401) {
                alert("Не авторизованы");
            } else {
                console.error(e);
                alert("Ошибка")
            }
        }
    };

    const decrement = async (id) => {
        try {
            await ProductService.decrementProduct(id);
            const updated = products
                .map(item =>
                    item.id === id
                        ? { ...item, quantity: item.quantity - 1 }
                        : item
                )
            setProducts(updated);
        } catch (e) {
            if (e.status === 401) {
                alert("Не авторизованы");
            } else {
                console.error(e);
                alert("Ощибка")
            }
        }
    };

    console.log(itemInCart);
    return (
        <div className={cl.card}>
            <div className={cl.imageWrapper}>
                <img src={itemInCart.image} alt={itemInCart.title} className={cl.image} />
                {(!shouldHideDelete) && (
                    <button className={cl.deleteButton} onClick={onDelete}>
                        &times;
                    </button>
                )}
            </div>
            <div className={cl.content}>
                <h3 className={cl.title}>{itemInCart.name}</h3>
                <p className={cl.price}>
                    {itemInCart.price.toLocaleString()} <span>₽</span>
                </p>
                {itemInCart.quantity !== 0 ? (
                    <div className={cl.counter}>
                        <button onClick={() => decrement(itemInCart.id)}>-</button>
                        <span>{itemInCart.quantity}</span>
                        <button onClick={() => increment(itemInCart.id)}>+</button>
                    </div>
                ) : (
                    <button className={cl.button} onClick={() => addCart(itemInCart)}>В корзину</button>
                )}
            </div>
        </div>
    );
};

export default ProductCard;
