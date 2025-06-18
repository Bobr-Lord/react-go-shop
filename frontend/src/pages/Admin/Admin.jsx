import React, {useEffect, useState} from 'react';
import cl from './Admin.module.css';
import AddItemForm from "../../componends/UI/AddItemForm/AddItemForm";
import MenuItem from "../../componends/UI/MenuItem/MenuItem";
import ProductService from "../../api/ProductService";
import {useFetching} from "../../hooks/useFetching";
import Loader from "../../componends/UI/Loader/Loader";
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
