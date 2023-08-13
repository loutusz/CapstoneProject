import React from "react";
import Sidebar from "@/components/sidebar";
import HeaderHome from "@/components/headerhome";

export default function Home() {
    return (
        <div className="flex">
            {/* <Sidebar/> */}
            <div className=""flex flex-col w-full>
                <HeaderHome/>
                <main className="p-8">Content</main>
            </div>
        </div>
    )
}