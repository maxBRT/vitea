import NavBar from '@/components/ui/navbar';
import Hero from '@/components/ui/hero';
import ResumeHero from '@/components/ui/resumeHero';
import FeaturesHero from '@/components/ui/featuresHero';
import HowItWorks from '@/components/ui/howitworks';
import Pricing from '@/components/ui/pricing';


export default function Landing() {
    return (
        <>
            <NavBar />
            <Hero />
            <ResumeHero />
            <FeaturesHero />
            <HowItWorks />
            <Pricing />
        </>
    )
}

