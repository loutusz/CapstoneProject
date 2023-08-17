import React from "react";
import HeaderHome from "@/components/headerhome";
import ContentHome from "@/components/contenthome";

export default function Home() {
    return (
        <>
            <div className="flex flex-col w-full">
                <HeaderHome/>
                <ContentHome/>
            </div>
        </>
    );
}