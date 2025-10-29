import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    Field,
    FieldDescription,
    FieldGroup,
    FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"

export function SignupForm({
    className,
    ...props
}: React.ComponentProps<typeof Card>) {
    const navigate = useNavigate()

    const [firstName, setFirstName] = useState("")
    const [lastName, setLastName] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (password !== confirmPassword) {
            alert("Passwords do not match")
            return
        }

        const userData = { firstName, lastName, email, password }

        try {
            const response = await fetch("/api/users", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(userData),
            })

            if (!response.ok) {
                const errorText = await response.text()
                console.error("Signup failed:", errorText)
                alert("Signup failed. Please try again.")
                return
            }

            navigate("/login")
        } catch (error) {
            console.error("Network or fetch error:", error)
            alert("An error occurred. Please check your connection.")
        }
    }

    return (
        <Card className={cn("max-w-md mx-auto mt-10", className)} {...props}>
            <CardHeader>
                <CardTitle>Create an account</CardTitle>
                <CardDescription>
                    Enter your information below to create your account
                </CardDescription>
            </CardHeader>

            <CardContent>
                <form onSubmit={handleSubmit} className="flex flex-col gap-6">
                    <FieldGroup>
                        <Field>
                            <FieldLabel htmlFor="first-name">First Name</FieldLabel>
                            <Input
                                id="first-name"
                                type="text"
                                placeholder="John"
                                required
                                value={firstName}
                                onChange={(e) => setFirstName(e.target.value)}
                            />
                        </Field>

                        <Field>
                            <FieldLabel htmlFor="last-name">Last Name</FieldLabel>
                            <Input
                                id="last-name"
                                type="text"
                                placeholder="Doe"
                                required
                                value={lastName}
                                onChange={(e) => setLastName(e.target.value)}
                            />
                        </Field>

                        <Field>
                            <FieldLabel htmlFor="email">Email</FieldLabel>
                            <Input
                                id="email"
                                type="email"
                                placeholder="m@example.com"
                                required
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                            />
                            <FieldDescription>
                                We&apos;ll use this to contact you. We will not share your email
                                with anyone else.
                            </FieldDescription>
                        </Field>

                        <Field>
                            <FieldLabel htmlFor="password">Password</FieldLabel>
                            <Input
                                id="password"
                                type="password"
                                required
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            <FieldDescription>
                                Must be at least 8 characters long.
                            </FieldDescription>
                        </Field>

                        <Field>
                            <FieldLabel htmlFor="confirm-password">
                                Confirm Password
                            </FieldLabel>
                            <Input
                                id="confirm-password"
                                type="password"
                                required
                                value={confirmPassword}
                                onChange={(e) => setConfirmPassword(e.target.value)}
                            />
                            <FieldDescription>Please confirm your password.</FieldDescription>
                        </Field>

                        <div className="flex flex-col gap-2">
                            <Button type="submit">Create Account</Button>
                            <Button variant="outline" type="button">
                                Sign up with Google
                            </Button>
                        </div>

                        <FieldDescription className="px-6 text-center">
                            Already have an account?{" "}
                            <a href="/login" className="underline">
                                Sign in
                            </a>
                        </FieldDescription>
                    </FieldGroup>
                </form>
            </CardContent>
        </Card>
    )
}

