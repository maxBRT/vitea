import { Badge } from "@/components/ui/badge";

export default function HowItWorks() {
    return (
        <section id="how-it-works" className="bg-gray-50 flex w-full flex-col items-center justify-center text-center py-52">
            {/* Headline */}
            <h1 className="text-5xl font-bold text-foreground mb-4">
                How it works
            </h1>

            {/* Subtext */}
            <p className="text-lg text-muted-foreground mb-8 max-w-xl">
                Three simple steps to your professional resume
            </p>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 text-center">
                <div className="flex flex-col text-center p-8 max-w-100 h-full justify-center items-center ">
                    <Badge variant="secondary" className="h-20 min-w-20">
                        <span className="text-2xl">1</span>
                    </Badge>
                    {/* Headline */}
                    <h3 className="text-3xl font-bold my-4 tracking-tight text-foreground">
                        Write or Upload
                    </h3>
                    <p className="text-lg text-muted-foreground">
                        Start with a template, upload your existing markdown file, or write from scratch in our intuitive editor.
                    </p>
                </div>
                <div className="flex flex-col text-center p-8 max-w-100 h-full justify-center items-center">
                    <Badge variant="secondary" className="h-20 min-w-20">
                        <span className="text-2xl">2</span>

                    </Badge>
                    <h3 className="text-3xl font-bold my-4 tracking-tight text-foreground">
                        Choose Your Theme
                    </h3>
                    <p className="text-lg text-muted-foreground">
                        Select from our collection of professional themes. See live preview as you customize your resume.
                    </p>
                </div>
                <div className="flex flex-col text-center p-8 h-full max-w-100 items-center justify-center">
                    <Badge variant="secondary" className="h-20 min-w-20">
                        <span className="text-2xl">3</span>
                    </Badge>
                    <h3 className="text-3xl font-bold my-4 tracking-tight text-foreground">
                        Share Your Link
                    </h3>
                    <p className="text-lg text-muted-foreground">
                        Get your unique hosted link instantly. Share it on LinkedIn, job applications, or anywhere you need.
                    </p>
                </div>
            </div>
        </section >
    )
}
