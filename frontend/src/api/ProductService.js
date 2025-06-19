import axios from 'axios';

export default class ProductService {
    static async addProduct(product) {
        return await axios.post('http://192.168.1.69:8080/api/product', product);
    }
    static async getProducts() {
        return await axios.get('http://192.168.1.69:8080/api/products');
    }
    static async deleteProduct(id) {
        return await axios.delete(`http://192.168.1.69:8080/api/product/${id}`);
    }
}