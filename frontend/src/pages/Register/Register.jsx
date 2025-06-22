import React from 'react';
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";
import cl from "./Register.module.css";
import AuthService from "../../api/AuthService";
import {useNavigate} from "react-router-dom";

const Register = () => {
    const [firstName, setFirstName] = React.useState('');
    const [lastName, setLastName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const navigate = useNavigate();

    async function handleSubmit(e) {
        e.preventDefault();
        console.log(firstName, lastName, email, password);
        try {
            const res = await AuthService.register(firstName, lastName, email, password);
            console.log(res.data.id);
            alert("Регистрация успешна!");
            navigate("/login");
        } catch (e) {
            console.error(e);
        } finally {
            setFirstName("");
            setLastName("");
            setEmail("");
            setPassword("");
        }

    }

    return (
        <div className={cl.main}>
            <div className={cl.container}>
                <div className={cl.title}>Регистрация</div>
                <form className={cl.form}>
                    <MyInput
                        className={cl.input}
                        value={firstName}
                        placeholder="Введите имя"
                        onChange={(e) => setFirstName(e.target.value)}
                    />
                    <MyInput
                        className={cl.input}
                        value={lastName}
                        placeholder="Введите фамилию"
                        onChange={(e) => setLastName(e.target.value)}
                    />
                    <MyInput
                        className={cl.input}
                        type="email"
                        value={email}
                        placeholder="Введите почту"
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <MyInput
                        className={cl.input}
                        value={password}
                        type="password"
                        placeholder="Введите пароль"
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <MyButton className={cl.button} onClick={handleSubmit}>
                        Отправить
                    </MyButton>
                </form>
            </div>
        </div>
    );
};

export default Register;
