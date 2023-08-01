import React from "react";
import {FaEnvelope, FaWhatsapp, FaLinkedin} from 'react-icons/fa'
export default function Footer (){
    return (
        <footer className="bg-blue-700 flex justify-between items-center px-5 py-3">
           <ul>
                <div className="text-white">
                    <h2 className="text-3xl font-bold">JICO </h2>
                    <p>© 2023 | Privacy - Terms</p>
                </div>
         </ul>
                  {/* Contact */}
                <div className="flex gap-4 items-end px-5">
                    <p>Contact US</p>
                    <FaEnvelope className="text-2xl cursor-pointer hover:text-red-600"/>
                    <FaWhatsapp className="text-2xl cursor-pointer hover:text-green-600"/>
                    <FaLinkedin className="text-2xl cursor-pointer hover:text-sky-400"/>               
                </div>   
        </footer>
    )}