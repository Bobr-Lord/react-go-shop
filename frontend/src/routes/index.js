import Home from "../pages/Home/Home";
import Admin from "../pages/Admin/Admin";

export const privateRoute = [
    {path: "/", component: <Home/>, exact: true},
    {path: "/admin", component: <Admin/>, exact: true},
]

export const publicRoute = [
    {path: "/", component: <Home/>, exact: true},
]