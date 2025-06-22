import axios from 'axios';

export default class ProductService {
    static async addProduct(product) {
        return await axios.post('http://localhost:8080/api/product', product, {
            withCredentials: true
        });
    }
    static async getProducts() {
        return await axios.get('http://localhost:8080/api/products', {
            withCredentials: true
        });
    }
    static async getProductsPrivate() {
        return await axios.get('http://localhost:8080/api/products/cart', {
            withCredentials: true
        });
    }
    static async deleteProduct(id) {
        return await axios.delete(`http://localhost:8080/api/product/${id}`, {
            withCredentials: true
        });
    }
    static async addItemCart(id) {
        return await axios.post(`http://localhost:8080/api/cart/item`, {
            id: id,
        }, {
            withCredentials: true
        });
    }
    static async deleteItemCart(id) {
        return await axios.delete(`http://localhost:8080/api/cart/item/${id}`, {
            withCredentials: true
        });
    }
    static async getItemsCart() {
        return await axios.get(`http://localhost:8080/api/cart/item`, {
            withCredentials: true
        });
    }
}