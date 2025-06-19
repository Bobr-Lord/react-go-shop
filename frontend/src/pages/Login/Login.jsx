import React from 'react';
import cl from "./Login.module.css";
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";

const Login = () => {
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');

    function handleLogin(e) {
        e.preventDefault();
        console.log("Логин:", email, password);
        setEmail('');
        setPassword('');
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
