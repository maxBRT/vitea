import { Button } from "@/components/ui/button"

export default function Hero() {
    return (
        <section className="flex flex-col items-center justify-center text-center p-24 pb-6 px-6">
            {/* Headline */}
            <h1 className="text-8xl font-bold tracking-tight text-foreground mb-4 max-w-4xl">
                Your Resume, Powered by
                <span className="text-blue-600"> Vitae </span>

            </h1>

            {/* Subtext */}
            <p className="text-lg text-muted-foreground mb-8 max-w-xl">
                Write, style, and share your professional resume with beautiful themes. Get a permanent hosted link that you can share anywhere.
            </p>

            {/* Buttons */}
            <div className="flex flex-wrap justify-center gap-4">
                <Button size="lg" className="rounded-full px-8">
                    Join for free
                </Button>
                <Button onClick={() => {
                    document.getElementById("pricing")?.scrollIntoView({ behavior: "smooth" })
                }}
                    variant="outline" size="lg" className="rounded-full px-8">
                    See our plans →
                </Button>
            </div>
        </section>
    )
}
