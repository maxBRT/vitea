import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Landing from './pages/Landing.tsx';
import Dashboard from './pages/Dashboard.tsx';
import Register from './pages/Register.tsx';
import Login from './pages/Login.tsx';

function App() {
    return (
        <div className="flex flex-col items-center justify-center bg-linear-to-b from-blue-100 to-stone-50 repeat-infinite">
            <BrowserRouter >
                <Routes >
                    <Route path="/" element={<Landing />} />
                    <Route path="/dashboard" element={<Dashboard />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/login" element={<Login />} />
                </Routes >
            </BrowserRouter >
        </div>
    );
}

export default App
