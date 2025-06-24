import React from 'react';
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";
import cl from "./Register.module.css";
import AuthService from "../../api/AuthService";
import { useNavigate } from "react-router-dom";

const Register = () => {
    const [firstName, setFirstName] = React.useState('');
    const [lastName, setLastName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [errors, setErrors] = React.useState({});
    const navigate = useNavigate();

    const validate = () => {
        const newErrors = {};
        if (firstName.trim().length < 3) {
            newErrors.firstName = "Имя должно быть не короче 3 символов";
        }
        if (lastName.trim().length < 3) {
            newErrors.lastName = "Фамилия должна быть не короче 3 символов";
        }
        if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
            newErrors.email = "Введите корректный email";
        }
        if (password.length < 8) {
            newErrors.password = "Пароль должен быть не менее 8 символов";
        }
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    async function handleSubmit(e) {
        e.preventDefault();
        if (!validate()) return;

        try {
            const res = await AuthService.register(email, password, firstName, lastName);
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
                        className={`${cl.input} ${errors.firstName ? cl.inputError : ''}`}
                        value={firstName}
                        placeholder="Введите имя"
                        onChange={(e) => setFirstName(e.target.value)}
                    />
                    {errors.firstName && <div className={cl.error}>{errors.firstName}</div>}

                    <MyInput
                        className={`${cl.input} ${errors.lastName ? cl.inputError : ''}`}
                        value={lastName}
                        placeholder="Введите фамилию"
                        onChange={(e) => setLastName(e.target.value)}
                    />
                    {errors.lastName && <div className={cl.error}>{errors.lastName}</div>}

                    <MyInput
                        className={`${cl.input} ${errors.email ? cl.inputError : ''}`}
                        type="email"
                        value={email}
                        placeholder="Введите почту"
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    {errors.email && <div className={cl.error}>{errors.email}</div>}

                    <MyInput
                        className={`${cl.input} ${errors.password ? cl.inputError : ''}`}
                        value={password}
                        type="password"
                        placeholder="Введите пароль"
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    {errors.password && <div className={cl.error}>{errors.password}</div>}

                    <MyButton className={cl.button} onClick={handleSubmit}>
                        Отправить
                    </MyButton>
                </form>
            </div>
        </div>
    );
};

export default Register;
