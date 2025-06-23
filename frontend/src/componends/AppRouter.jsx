import React, {useContext} from 'react';
import {Navigate, Route, Routes} from "react-router-dom";
import {privateAdminRoute, privateUserRoute, publicRoute} from "../routes";
import {AuthContext} from "../context";

const AppRouter = () => {
    const {isAdmin, isLoggedIn} = useContext(AuthContext);
    return (
        <Routes>
            {
                isLoggedIn
                    ?
                    (isAdmin
                        ? privateAdminRoute.map((route) =>
                            <Route path={route.path} element={route.component} />
                        )
                        : privateUserRoute.map((route) =>
                            <Route path={route.path} element={route.component} />
                        )
                    )
                    : publicRoute.map((route) =>
                        <Route path={route.path} element={route.component} />
                    )
            }

            <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
    );
};

export default AppRouter;