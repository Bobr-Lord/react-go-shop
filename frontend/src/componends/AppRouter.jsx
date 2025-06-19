import React, {useContext} from 'react';
import {Navigate, Route, Routes} from "react-router-dom";
import {privateRoute, publicRoute} from "../routes";
import {AuthContext} from "../context";

const AppRouter = () => {
    const {isAdmin} = useContext(AuthContext);
    return (
        <Routes>
            {
                isAdmin
                    ?
                        privateRoute.map((route) =>
                            <Route path={route.path} element={route.component} />
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