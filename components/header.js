import React from "react";
import Link from "next/link";

export default function Header() {
    return (
        <header className="bg-custombluegray py-4 px-6 flex justify-between items-center">
            <div className="text-white text-5xl font-bold">Jico</div>
            <nav className="flex justify-center items-center space-x-8">
                <Link className="text-white hover:text-gray-300 font-semibold" href="/">
                    Home
        </Link>
        <Link className="text-white hover:text-gray-300 font-semibold" href="/">
         Features
        </Link>
      </nav>
      <div className="space-x-4">
        <Link className="text-white hover:text-gray-300 font-semibold" href="/registerPage">
          Register
        </Link>
        <Link className="text-white hover:text-gray-300 font-semibold" href="/loginPage">
          Login
        </Link>
      </div>
    </header>
  );
};


