import React from "react";
import {FaEnvelope, FaWhatsapp} from 'react-icons/fa'
export default function Footer (){
    return (
        // py-8 px-6 flex justify-between items-center
        <footer className="bg-blue-700 h-1/4 w-full flex md:flex-row flex-col justify-around items-start p-15">
            <div className="text-white py-14 px-14 pl-24 ">
                <h2 className="text-4xl font-bold">JICO </h2>
                <p>© 2023</p>
                <p>Privacy - Terms</p>
            </div>
        </footer>
    )}