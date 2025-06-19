import Home from "../pages/Home/Home";
import Admin from "../pages/Admin/Admin";
import Register from "../pages/Register/Register";
import Loader from "../componends/UI/Loader/Loader";
import Login from "../pages/Login/Login";

export const privateRoute = [
    {path: "/", component: <Home/>, exact: true},
    {path: "/admin", component: <Admin/>, exact: true},
    {path: "/login", component: <Login/>, exact: true},
    {path: "/register", component: <Register/>, exact: true},
]

export const publicRoute = [
    {path: "/", component: <Home/>, exact: true},
    {path: "/login", component: <Register/>, exact: true},
    {path: "/register", component: <Login/>, exact: true},
]