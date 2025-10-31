import { useEffect, useState, useCallback, useRef } from "react"
import { useNavigate } from "react-router-dom"
import MDEditor from "@uiw/react-md-editor";
import Navbar from "@/components/ui/navbar";
import { Button } from "@/components/ui/button";
import { MdOutlineFileUpload } from "react-icons/md";

interface User {
    id: string;
    first_name: string;
    last_name: string;
    email: string;
    created_at: string;
    updated_at: string;
}

export default function Dashboard() {
    const navigate = useNavigate();
    const [user, setUser] = useState<User | null>(null);

    const fileInputRef = useRef<HTMLInputElement | null>(null);

    const handleButtonClick = () => {
        fileInputRef.current?.click();
    };

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (!file) return;

        const reader = new FileReader();
        reader.onload = (event) => {
            const text = event.target?.result as string;
            setValue(text);
        };
        reader.readAsText(file);
    };

    const fetchUser = useCallback(async () => {
        const token = localStorage.getItem("access_token");

        const res = await fetch("/api/me", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${token}`,
            },
        });

        if (res.status === 401) {
            console.warn("Access token expired, attempting refresh...");
            const refreshToken = localStorage.getItem("refresh_token");
            const refreshRes = await fetch("api/refresh", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${refreshToken}`,
                },
            });

            if (!refreshRes.ok) {
                throw new Error(`HTTP error! Status: ${refreshRes.status}`);
            }
            const refreshData = await refreshRes.json();
            localStorage.setItem("access_token", refreshData.access_token);
            return fetchUser();
        }

        if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);
        const data = await res.json();
        setUser(data);
    }, []);



    useEffect(() => {
        try {
            fetchUser();
        } catch (e) {
            console.error(e);
            navigate("/login");
        }
    }, [fetchUser, navigate]);


    const [value, setValue] = useState<string>('');

    return (
        <div className="h-screen w-full overflow-hidden pt-14 flex flex-col">
            <Navbar />

            {/* Content area that grows to fill space between Navbar and Button */}
            <div className="flex-1 overflow-hidden">
                <MDEditor
                    height="100%"
                    value={value}
                    data-color-mode="light"
                    preview="live"
                    onChange={(val) => setValue(val ?? "")}
                />
            </div>

            {/* Button area pinned to bottom */}
            <div className="flex justify-start items-center bg-white/80 backdrop-blur-md border-t border-gray-200 shadow-sm">
                <input
                    type="file"
                    accept=".md"
                    ref={fileInputRef}
                    onChange={handleFileChange}
                    className="hidden"
                />
                <Button
                    size="lg"
                    variant="outline"
                    onClick={handleButtonClick}
                >
                    <MdOutlineFileUpload className="text-xl" />

                    Upload .md file
                </Button>
            </div>
        </div>
    );

}
