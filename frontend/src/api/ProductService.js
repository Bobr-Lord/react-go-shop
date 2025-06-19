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
    static async deleteProduct(id) {
        return await axios.delete(`http://localhost:8080/api/product/${id}`, {
            withCredentials: true
        });
    }
}