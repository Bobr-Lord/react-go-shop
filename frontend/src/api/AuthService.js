import axios from "axios";

export default class AuthService {
    static async login(email, password) {
        return await axios.post("/api/auth/login", {
            email: email,
            password: password,
        }, {
            withCredentials: true
        })
    }
    static async register(email, password, firstname, lastname) {
        return await axios.post("/api/auth/reg", {
            email: email,
            password: password,
            first_name: firstname,
            last_name: lastname,
        }, {
            withCredentials: true
        })
    }
    static async getMe() {
        return await axios.get(`/api/auth/me`, {
            withCredentials: true
        });
    }
}