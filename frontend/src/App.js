import React from 'react';
import './styles/App.css'
import Home from "./pages/Home/Home";
import Navbar from "./componends/UI/Navbar/Navbar";
import Admin from "./pages/Admin/Admin";
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import AppRouter from "./componends/AppRouter";
import Loader from "./componends/UI/Loader/Loader";

export default function App() {
  return (
      <BrowserRouter>
          <Navbar />
          <AppRouter/>
      </BrowserRouter>

  );
}

