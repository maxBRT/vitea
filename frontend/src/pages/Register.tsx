import NavBar from "@/components/ui/navbar";
import { SignupForm } from "@/components/signup-form";

export default function Register() {
    return (
        <>

            <div className="h-screen w-full max-w-3xl">
                <NavBar />
                <SignupForm className="mt-20 w-full max-w-3xl" />
            </div>
        </>

    );
}

