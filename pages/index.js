import Link from "next/link";
import Layout from "../components/layout"
import {FaDiscord} from 'react-icons/fa'

export default function Home() {
  <div>
    <head>
      <title>Home</title>
    </head>
  </div>
  return (
    <div>
      <Layout/>
      <div className="container mx-auto px-5 py-40 flex flex-col w-1/4 ml-8 min-h-screen justify-center items-center">
        <h1 className="text-xl text-start font-bold text-black">
          WELCOME TO JICO
        </h1>
        <div className="ml-50">
        <img src="/assets/teamsdc-removebg-preview.png"/>
        </div>
        <div>
        <p className="text-lg text-start font-medium text-black mb-4">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
         labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco 
         laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in 
         voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat 
         non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
        </p>
        </div>
        
        <button className="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-4 py-2 rounded">
          Get Started
        </button>
        <FaDiscord size="25px"/>
      </div>
    </div>
  )
}