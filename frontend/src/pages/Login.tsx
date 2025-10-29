import { LoginForm } from "@/components/login-form";
import NavBar from "@/components/ui/navbar";

export default function Login() {

    return (
        <>
            <div className="h-screen w-full max-w-3xl">
                <NavBar />
                <LoginForm className="pb-30 pt-24 px-24 w-full max-w-3xl" />
            </div>
        </>
    );
}
