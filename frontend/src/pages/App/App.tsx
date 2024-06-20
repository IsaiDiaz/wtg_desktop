import { useState } from 'react';
import './App.css';
import { Greet } from "../../../wailsjs/go/main/App";
import Navbar from '../../components/Navbar';
import { Outlet } from 'react-router-dom';

function App() {
    const [resultText, setResultText] = useState("Porfavor, ingresa tu nombre.");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <div style={{ display: 'flex' }}>
            <Navbar />
            <div style={{ flex: 1, padding: '10px' }}>
                <Outlet />
            </div>
        </div>
    )
}

export default App
