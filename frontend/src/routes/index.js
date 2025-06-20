import Home from "../pages/Home/Home";
import Admin from "../pages/Admin/Admin";
import Register from "../pages/Register/Register";
import Login from "../pages/Login/Login";
import Cart from "../pages/Cart/Cart";

export const privateAdminRoute = [
    {path: "/", component: <Home/>, exact: true},
    {path: "/admin", component: <Admin/>, exact: true},
    {path: "/cart", component: <Cart/>, exact: true},
]

export const privateUserRoute = [
    {path: "/", component: <Home/>, exact: true},
    {path: "/cart", component: <Cart/>, exact: true},
]

export const publicRoute = [
    {path: "/", component: <Home/>, exact: true},
    {path: "/login", component: <Login/>, exact: true},
    {path: "/register", component: <Register/>, exact: true},
]