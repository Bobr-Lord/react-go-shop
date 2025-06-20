import React, {useEffect} from 'react';
import cl from './Cart.module.css';
import ProductService from "../../api/ProductService";
import MyButton from "../../componends/UI/MyButton/MyButton";

const Cart = () => {
    const [cartItems, setCartItems] = React.useState([]);
    const total = (cartItems === null)
        ? 0
        : cartItems.reduce((acc, item) => acc + item.price * item.quantity, 0);

    useEffect(() => {
        ProductService.getItemsCart().then((response) => {
            console.log(response);
            setCartItems(response.data);
        }).catch((error) => {
            console.log(error);
        })
    }, [])

    async function RemoveItem(item) {
        try {
            console.log(item.id);
            const res = await ProductService.deleteItemCart(item.id);
            console.log(res);
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className={cl.cart}>
            <h2 className={cl.title}>Корзина</h2>
            {(cartItems === null || cartItems.length === 0) ? (
                <p className={cl.empty}>Ваша корзина пуста</p>
            ) : (
                <div className={cl.items}>
                    {cartItems.map(item => (
                        <div key={item.id} className={cl.item}>
                            <img src={item.image} alt={item.name} className={cl.image} />
                            <div className={cl.info}>
                                <h3>{item.name}</h3>
                                <p>{item.price} ₽ × {item.quantity}</p>
                                <MyButton onClick={() => RemoveItem(item)}>Удалить</MyButton>
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
