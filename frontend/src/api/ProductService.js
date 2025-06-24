import axios from 'axios';

export default class ProductService {
    static async addProduct(product) {
        return await axios.post('/api/shop/product', product, {
            withCredentials: true
        });
    }
    static async getProducts() {
        return await axios.get('/api/shop/products', {
            withCredentials: true
        });
    }
    static async getProductsPrivate() {
        return await axios.get('/api/shop/products/cart', {
            withCredentials: true
        });
    }
    static async deleteProduct(id) {
        return await axios.delete(`/api/shop/product/${id}`, {
            withCredentials: true
        });
    }
    static async addItemCart(id) {
        return await axios.post(`/api/shop/cart/item`, {
            id: id,
        }, {
            withCredentials: true
        });
    }
    static async deleteItemCart(id) {
        return await axios.delete(`/api/shop/cart/item/${id}`, {
            withCredentials: true
        });
    }
    static async getItemsCart() {
        return await axios.get(`/api/shop/cart/item`, {
            withCredentials: true
        });
    }
    static async decrementProduct(id) {
        return await axios.put(`/api/shop/cart/item/${id}`, {}, {
            withCredentials: true
        });
    }
}