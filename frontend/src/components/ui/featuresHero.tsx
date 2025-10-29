import { Card, CardContent, CardTitle, CardHeader } from "@/components/ui/card"
import { FcFlashOn, FcDocument } from "react-icons/fc"
import { FaPalette } from "react-icons/fa";

export default function FeaturesHero() {
    return (
        <section id="features" className="flex max-w-6xl flex-col items-center justify-center text-center py-64 px-6">
            {/* Headline */}
            <h1 className="text-5xl font-bold tracking-tight text-foreground mb-4 max-w-2xl">
                Everything You Need
            </h1>

            {/* Subtext */}
            <p className="text-lg text-muted-foreground mb-8 max-w-xl">
                Powerful features to create and share your perfect resume.
            </p>
            {/* Features cards */}
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 text-center">
                <Card className="flex flex-col text-left p-8 h-full shadow-sm hover:shadow-md transition-all">

                    <FcDocument className="text-5xl" />
                    <CardHeader >
                        <CardTitle className="text-2xl">Markdown Editor</CardTitle>
                    </CardHeader>
                    <CardContent className="p-0 text-sm text-muted-foreground">
                        Write your resume in clean, simple markdown with live preview.
                        Upload existing files or start from scratch.
                    </CardContent>
                </Card>

                <Card className="flex flex-col text-left p-8 h-full shadow-sm hover:shadow-md transition-all">

                    <FaPalette className="text-5xl" />
                    <CardHeader>
                        <CardTitle className="text-2xl">Beautiful Themes</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <p className="text-sm text-muted-foreground">
                            Choose from professionally designed themes. Free themes included, premium themes for advanced styling.
                        </p>
                    </CardContent>
                </Card>
                <Card className="flex flex-col text-left p-8 h-full shadow-sm hover:shadow-md transition-all">

                    <FcFlashOn className="text-5xl" />
                    <CardHeader>
                        <CardTitle className="text-2xl">Instant Hosting</CardTitle>
                    </CardHeader>

                    <CardContent>
                        <p className="text-sm text-muted-foreground">
                            Get a shareable link instantly. Your resume is hosted and accessible from anywhere, anytime.
                        </p>
                    </CardContent>
                </Card>

            </div>
        </section >
    )
}
