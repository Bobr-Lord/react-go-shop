import React, {useContext} from 'react';
import cl from "./Login.module.css";
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";
import AuthService from "../../api/AuthService";
import {Link, useNavigate} from "react-router-dom";
import {AuthContext} from "../../context";

const Login = () => {
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [authError, setAuthError] = React.useState('');
    const navigate = useNavigate();
    const {setIsLoggedIn} = useContext(AuthContext);

    async function handleLogin(e) {
        e.preventDefault();
        setAuthError(''); // сбрасываем ошибку перед новым запросом

        try {
            const res = await AuthService.login(email, password);
            setIsLoggedIn(true);
            navigate("/");
        } catch (e) {
            console.error(e);
            if (e.response?.status === 401) {
                setAuthError("Неверный логин или пароль");
            } else {
                setAuthError("Произошла ошибка. Повторите позже");
            }
        }
    }

    return (
        <div className={cl.main}>
            <div className={cl.container}>
                <div className={cl.title}>Вход</div>
                <form className={cl.form}>
                    <MyInput
                        className={`${cl.input} ${authError ? cl.inputError : ''}`}
                        value={email}
                        placeholder="Введите почту"
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <MyInput
                        className={`${cl.input} ${authError ? cl.inputError : ''}`}
                        type="password"
                        value={password}
                        placeholder="Введите пароль"
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    {authError && <div className={cl.error}>{authError}</div>}
                    <MyButton className={cl.button} onClick={handleLogin}>
                        Войти
                    </MyButton>
                </form>
            </div>
            <div className={cl.registerPrompt}>
                У вас нет учетной записи? <Link to="/register" className={cl.registerLink}>Зарегистрироваться</Link>
            </div>
        </div>
    );
};

export default Login;
