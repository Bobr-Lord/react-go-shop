import axios from 'axios';

export default class ProductService {
    static async addProduct(product) {
        return await axios.post('http://localhost:8080/api/product', product);
    }
    static async getProducts() {
        return await axios.get('http://localhost:8080/api/products');
    }
}