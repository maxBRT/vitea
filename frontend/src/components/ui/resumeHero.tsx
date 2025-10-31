import { Card, CardContent } from "@/components/ui/card"

export default function ResumeHero() {
    return (
        <div className="max-w-6xl mx-auto mt-4 rounded-2xl border bg-linear-to-b from-muted/60 via-background to-background shadow-xl overflow-hidden">
            {/* Fake app window header */}
            <div className="flex items-center gap-2 px-4 py-2 border-b bg-muted/60 text-sm font-medium text-muted-foreground backdrop-blur">
                <div className="flex gap-1">
                    <span className="w-3 h-3 rounded-full bg-red-400" />
                    <span className="w-3 h-3 rounded-full bg-yellow-400" />
                    <span className="w-3 h-3 rounded-full bg-green-400" />
                </div>
                <span className="mx-auto text-xs tracking-wide text-muted-foreground">resume-editor</span>
            </div>

            {/* Two-column layout */}
            <div className="grid grid-cols-1 md:grid-cols-2 divide-y md:divide-x md:divide-y-0">
                {/* Markdown editor side */}
                <Card className="rounded-none border-0 bg-muted/20">
                    <CardContent className="p-6 font-mono text-sm leading-relaxed text-muted-foreground">
                        <p>
                            <span className="text-blue-600 font-bold"># John Doe</span><br />
                            <span className="text-green-600 font-bold">## Software Engineer</span><br /><br />
                            john@email.com<br />
                            (555) 123-4567<br /><br />
                            <span className="text-blue-600 font-bold">### Experience</span><br />
                            **Senior Developer** at *TechCorp*
                        </p>
                        <ul>
                            <li>Led a team of 4 engineers building scalable APIs.</li>
                            <li>Optimized CI/CD pipelines, cutting deploy time by 30%.</li>
                            <li>Refined front-end architecture with React & TypeScript.</li>
                        </ul>
                    </CardContent>
                </Card>

                {/* Live preview side */}
                <Card className="rounded-none border-0 bg-white/90 backdrop-blur-sm">
                    <CardContent className="p-10 text-left">
                        <div className="max-w-md mx-auto">
                            {/* Header */}
                            <header className="border-b border-muted pb-4 mb-6">
                                <h1 className="text-4xl font-extrabold text-foreground tracking-tight">
                                    John <span className="text-primary">Doe</span>
                                </h1>
                                <h2 className="text-lg text-muted-foreground font-medium">
                                    Software Engineer
                                </h2>
                                <div className="mt-2 text-sm text-muted-foreground">
                                    <p>john@email.com</p>
                                    <p>(555) 123-4567</p>
                                </div>
                            </header>

                            {/* Experience */}
                            <section>
                                <h3 className="uppercase tracking-[0.15em] text-xs font-semibold text-muted-foreground mb-3">
                                    EXPERIENCE
                                </h3>
                                <div className="space-y-3">
                                    <div className="rounded-lg border border-muted p-4 shadow-sm hover:shadow-md transition-shadow">
                                        <div className="flex items-baseline justify-between">
                                            <h4 className="text-base font-semibold text-foreground">
                                                Senior Developer <span className="text-primary font-normal">@ TechCorp</span>
                                            </h4>
                                        </div>
                                        <ul className="mt-2 list-disc list-inside text-sm text-muted-foreground leading-relaxed">
                                            <li>Led a team of 4 engineers building scalable APIs.</li>
                                            <li>Optimized CI/CD pipelines, cutting deploy time by 30%.</li>
                                            <li>Refined front-end architecture with React & TypeScript.</li>
                                        </ul>
                                    </div>
                                </div>
                            </section>

                        </div>
                    </CardContent>
                </Card>
            </div>
        </div>
    )
}
