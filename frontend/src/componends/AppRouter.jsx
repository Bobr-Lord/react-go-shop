import React from 'react';
import {Navigate, Route, Routes} from "react-router-dom";
import Home from "../pages/Home/Home";
import Admin from "../pages/Admin/Admin";

const AppRouter = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/admin" element={<Admin />} />
            <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
    );
};

export default AppRouter;