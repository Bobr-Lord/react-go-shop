import './styles/App.css'
import Home from "./pages/Home/Home";
import Navbar from "./componends/UI/Navbar/Navbar";
import Admin from "./pages/Admin/Admin";

export default function App() {
  return (
    <div className="App">
        <Navbar />
        <Admin/>
    </div>
  );
}

