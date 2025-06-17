import './styles/App.css'
import Main from "./pages/Main/Main";
import Navbar from "./componends/UI/Navbar/Navbar";

export default function App() {
  return (
    <div className="App">
        <Navbar />
        <Main/>
    </div>
  );
}

