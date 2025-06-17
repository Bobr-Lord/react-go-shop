import React from 'react';
import cl from './MyInput.module.css'
const MyInput = (props) => {
    return (
        <input className={cl.input} {...props} />
    );
};

export default MyInput;