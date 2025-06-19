import React from 'react';
import cl from "./Login.module.css";
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";
import AuthService from "../../api/AuthService";
import {useNavigate} from "react-router-dom";

const Login = () => {
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const navigate = useNavigate();
    async function handleLogin(e) {
        e.preventDefault();
        console.log("Логин:", email, password);

        try {
            const res = await AuthService.login(email, password);
            console.log(res.data);
            navigate("/");
        } catch (e) {
            console.error(e);
        } finally {
            setEmail('');
            setPassword('');
        }
    }

    return (
        <div className={cl.main}>
            <div className={cl.container}>
                <div className={cl.title}>Вход</div>
                <form className={cl.form}>
                    <MyInput
                        className={cl.input}
                        value={email}
                        placeholder="Введите почту"
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <MyInput
                        className={cl.input}
                        type="password"
                        value={password}
                        placeholder="Введите пароль"
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <MyButton className={cl.button} onClick={handleLogin}>
                        Войти
                    </MyButton>
                </form>
            </div>
        </div>
    );
};

export default Login;
