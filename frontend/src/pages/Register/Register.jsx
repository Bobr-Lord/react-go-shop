import React from 'react';
import MyInput from "../../componends/UI/MyInput/MyInput";
import MyButton from "../../componends/UI/MyButton/MyButton";
import cl from "./Register.module.css";

const Register = () => {
    const [firstName, setFirstName] = React.useState('');
    const [lastName, setLastName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');

    function handleSubmit(e) {
        e.preventDefault();
        console.log(firstName, lastName, email, password);
        setFirstName("");
        setLastName("");
        setEmail("");
        setPassword("");
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
