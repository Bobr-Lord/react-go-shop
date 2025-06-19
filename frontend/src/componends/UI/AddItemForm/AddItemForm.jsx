import React from 'react';
import cl from "./AddItemForm.module.css";
import MyInput from "../MyInput/MyInput";
import MyButton from "../MyButton/MyButton";
import ProductService from "../../../api/ProductService";

const AddItemForm = ({form, setForm, products, setProducts}) => {
    const handleChange = (e) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };
    const handleAddProduct = async () => {
        const newProduct = {
            name: form.title,
            price: Number(form.price),
            description: form.description,
            image: form.image,
            category: "product",
        };
        try {
            const res = await ProductService.addProduct(newProduct);
            console.log('запрос на добавление товара:', res);
            newProduct.id = res.data.id
            console.log(newProduct);
            if (products === null) {
                setProducts([newProduct]);
            } else {
                setProducts([newProduct, ...products]);
            }
            setForm({ title: '', price: '', description: '', image: '' });
        } catch (err) {
            console.error('Ошибка при добавлении товара:', err);
        }
    };





    return (
        <div className={cl.form}>
            <MyInput
                type="text"
                name="title"
                placeholder="Название"
                value={form.title}
                onChange={handleChange}
            />
            <MyInput
                type="text"
                name="price"
                placeholder="Цена"
                value={form.price}
                onChange={handleChange}
            />
            <MyInput
                type="text"
                name="description"
                placeholder="Описание"
                value={form.description}
                onChange={handleChange}
            />
            <MyInput
                type="text"
                name="image"
                placeholder="URL изображения"
                value={form.image}
                onChange={handleChange}
            />
            <MyButton onClick={handleAddProduct}>Добавить товар</MyButton>
        </div>
    );
};

export default AddItemForm;