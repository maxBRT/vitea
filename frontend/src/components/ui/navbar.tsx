import { Button } from '@/components/ui/button';
import { IoDocumentText } from "react-icons/io5";
import { useNavigate, useLocation } from "react-router-dom";
import { FaArrowAltCircleUp } from "react-icons/fa";

export default function NavBar() {
    const navigate = useNavigate();
    const location = useLocation();

    return (
        <div className="flex justify-center bg-transparent">
            <nav className="fixed rounded-2xl top-0 z-50 flex items-center justify-between w-full max-w-7xl px-6 md:px-10 py-3 bg-white/80 backdrop-blur-md border-b border-gray-200 shadow-sm transition-all duration-300">

                {/* Logo */}
                <div className="flex items-center gap-2 hover:opacity-90 transition-opacity">

                    <IoDocumentText className="text-3xl text-blue-600" />
                    <Button
                        onClick={() => {
                            navigate("/")
                        }}
                        variant="ghost"
                        size="sm"
                        className="text-gray-700 hover:text-blue-600 transition-colors"
                    >
                        <span className="text-2xl font-semibold tracking-tight text-gray-800">Vitae</span>
                    </Button>
                </div>

                {/* Navigation Links */}
                {location.pathname === "/" && (
                    <div className="hidden md:flex items-center gap-6">
                        <a href="#" className="text-gray-600 hover:text-blue-600 transition-colors">
                            <FaArrowAltCircleUp className="text-2xl text-gray-600 hover:text-blue-600 transition-colors" />
                        </a>
                        <a href="#features" className="text-gray-600 hover:text-blue-600 transition-colors">Features</a>
                        <a href="#how-it-works" className="text-gray-600 hover:text-blue-600 transition-colors">How it works</a>
                        <a href="#pricing" className="text-gray-600 hover:text-blue-600 transition-colors">Pricing</a>
                    </div>
                )}
                {/* Navigation Buttons */}
                <div className="flex items-center gap-3">
                    <Button
                        onClick={() => {
                            navigate("/login")
                        }}
                        variant="ghost"
                        size="sm"
                        className="text-gray-700 hover:text-blue-600 transition-colors"
                    >
                        Log in
                    </Button>
                    {location.pathname === "/" && (
                        <Button
                            onClick={() => {
                                navigate("/register")
                            }}
                            size="sm"
                            className="rounded-full px-6 bg-blue-600 hover:bg-blue-700 text-white transition-all"
                        >
                            Get Started
                        </Button>
                    )}
                </div>
            </nav >
        </div >
    );
}

