import axios from "axios";

export default class AuthService {
    static async login(email, password) {
        return await axios.post("/login", {
            email: email,
            password: password,
        })
    }
    static async register(email, password, firstname, lastname) {
        return await axios.post("/register", {
            email: email,
            password: password,
            firstname: firstname,
            lastname: lastname,
        })
    }
}