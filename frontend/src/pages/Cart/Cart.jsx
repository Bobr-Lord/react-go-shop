import React from 'react';
import cl from './Cart.module.css';

const Cart = () => {
    const cartItems = []
    const onRemove = {}
    const total = cartItems.reduce((acc, item) => acc + item.price * item.quantity, 0);

    return (
        <div className={cl.cart}>
            <h2 className={cl.title}>Корзина</h2>
            {cartItems.length === 0 ? (
                <p className={cl.empty}>Ваша корзина пуста</p>
            ) : (
                <div className={cl.items}>
                    {cartItems.map(item => (
                        <div key={item.id} className={cl.item}>
                            <img src={item.image} alt={item.title} className={cl.image} />
                            <div className={cl.info}>
                                <h3>{item.title}</h3>
                                <p>{item.price} ₽ × {item.quantity}</p>
                                {/*<MyButton onClick={() => onRemove(item.id)}>Удалить</MyButton>*/}
                            </div>
                        </div>
                    ))}
                    <div className={cl.total}>Итого: {total.toLocaleString()} ₽</div>
                </div>
            )}
        </div>
    );
};

export default Cart;
