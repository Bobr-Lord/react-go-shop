import React, {useState} from 'react';
import cl from './Admin.module.css';
import AddItemForm from "../../componends/UI/AddItemForm/AddItemForm";
import MenuItemLoader from "../../componends/MenuItemLoader/MenuItemLoader";

const Admin = () => {
    const [products, setProducts] = useState([]);
    const [form, setForm] = useState({
        id: '',
        title: '',
        price: '',
        description: '',
        image: ''
    });

    return (
        <div className={cl.admin}>
            <h2>Панель администратора</h2>
            <AddItemForm form={form} setForm={setForm} products={products} setProducts={setProducts} />
            <MenuItemLoader products={products} setProducts={setProducts}/>
        </div>
    );
};

export default Admin;
