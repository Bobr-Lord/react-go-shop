import React, { useState } from 'react';
import cl from './Admin.module.css';
import ProductCard from "../../componends/UI/ProductCard/ProductCard";
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";
import AddItemForm from "../../componends/UI/AddItemForm/AddItemForm";

const Admin = () => {
    const [products, setProducts] = useState([]);
    const [form, setForm] = useState({
        title: '',
        price: '',
        description: '',
        image: ''
    });


    const handleDelete = (id) => {
        setProducts(products.filter(p => p.id !== id));
    };
    return (
        <div className={cl.admin}>
            <h2>Панель администратора</h2>
            <AddItemForm form={form} setForm={setForm} products={products} setProducts={setProducts} />

            <div className={cl.grid}>
                {products.map(product => (
                    <div key={product.id} className={cl.cardWrapper}>
                        <ProductCard
                            title={product.title}
                            price={product.price}
                            description={product.description}
                            image={product.image}
                        />
                        <button className={cl.deleteButton} onClick={() => handleDelete(product.id)}>
                            &times;
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Admin;
