import axios from "axios";

export default class AuthService {
    static async login(email, password) {
        return await axios.post("http://localhost:8081/api/login", {
            email: email,
            password: password,
        }, {
            withCredentials: true
        })
    }
    static async register(email, password, firstname, lastname) {
        return await axios.post("http://localhost:8081/api/reg", {
            email: email,
            password: password,
            firstname: firstname,
            lastname: lastname,
        }, {
            withCredentials: true
        })
    }
    static async getMe() {
        return await axios.get(`http://localhost:8081/api/me`, {
            withCredentials: true
        });
    }
}