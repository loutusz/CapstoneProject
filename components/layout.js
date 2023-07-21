import React from "react";
import Link from "next/link";

export default function Header() {
    return (
        <div>
            <header className="bg-white sticky top-0 h-14 flex font-bold text-xl text-purple-400">
                JICO
                <div className="text-right">
                <Link href="/">
                    Home 
                </Link>
                <Link href="/loginPage">
                    Login
                </Link>
                <Link href="/registerPage">
                    Register
                </Link>
                </div>
            </header>
        </div>
    )
}