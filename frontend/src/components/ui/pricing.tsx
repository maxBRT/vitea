import { Card, CardContent, CardTitle, CardHeader, CardDescription, CardFooter } from "@/components/ui/card"
import { Badge } from "./badge"
import { Button } from "./button"


export default function Pricing() {
    return (
        <section id="pricing" className="flex max-w-6xl flex-col items-center justify-center text-center py-64 pb-24 px-6">
            {/* Headline */}
            <h1 className="text-6xl font-bold tracking-tight text-foreground mb-4 max-w-4xl">
                Simple, Transparent Pricing
            </h1>
            <p className="text-lg text-muted-foreground mb-8 max-w-xl">
                Choose the plan that works for you
            </p>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-6xl mx-auto py-12">
                {/* Free Plan */}
                <Card className="flex text-left flex-col justify-between border border-gray-200 shadow-sm">
                    <CardHeader>
                        <CardTitle>Free</CardTitle>
                        <CardDescription>Perfect for trying out the platform</CardDescription>
                    </CardHeader>
                    <CardContent className="space-y-3">
                        <h2 className="text-3xl font-bold">$0<span className="text-base font-normal text-gray-500">/forever</span></h2>
                        <ul className="space-y-2 text-sm text-gray-700">
                            <li>- Limited theme selection</li>
                            <li>- 30-day hosted links</li>
                            <li>- Markdown editor</li>
                            <li>- PDF export</li>
                            <li>- Re-upload every 30 days</li>
                        </ul>
                    </CardContent>
                    <CardFooter>
                        <Button className="w-full" variant="outline">Get Started</Button>
                    </CardFooter>
                </Card>

                {/* Basic Plan (Most Popular) */}
                <Card className="relative text-left flex flex-col justify-between border-2 border-blue-600 shadow-md">
                    <Badge className="absolute -top-3 left-1/2 -translate-x-1/2 bg-blue-600 text-white text-xs px-3 py-1 rounded-full">
                        Most Popular
                    </Badge>
                    <CardHeader>
                        <CardTitle>Basic</CardTitle>
                        <CardDescription>For professionals who need reliability</CardDescription>
                    </CardHeader>
                    <CardContent className="space-y-3">
                        <h2 className="text-3xl font-bold">$1<span className="text-base font-normal text-gray-500">/month</span></h2>
                        <ul className="space-y-2 text-sm text-gray-700">
                            <li>- Limited theme selection</li>
                            <li className="font-semibold">- Permanent hosted links</li>
                            <li>- Markdown editor</li>
                            <li>- PDF export</li>
                            <li>- Priority support</li>
                        </ul>
                    </CardContent>
                    <CardFooter>
                        <Button className="w-full bg-blue-600 hover:bg-blue-700 text-white">Start Basic Plan</Button>
                    </CardFooter>
                </Card>

                {/* Premium Plan */}
                <Card className="flex flex-col text-left justify-between border border-gray-200 shadow-sm">
                    <CardHeader>
                        <CardTitle>Premium</CardTitle>
                        <CardDescription>For those who want the best</CardDescription>
                    </CardHeader>
                    <CardContent className="space-y-3">
                        <h2 className="text-3xl font-bold">$3<span className="text-base font-normal text-gray-500">/month</span></h2>
                        <ul className="space-y-2 text-sm text-gray-700">
                            <li className="font-semibold">- All premium themes</li>
                            <li className="font-semibold">- Permanent hosted links</li>
                            <li>- Markdown editor</li>
                            <li>- PDF export</li>
                            <li>- Custom domain support</li>
                            <li>- Analytics dashboard</li>
                        </ul>
                    </CardContent>
                    <CardFooter>
                        <Button className="w-full" variant="outline">Start Premium</Button>
                    </CardFooter>
                </Card>
            </div>
        </section >
    )
}
