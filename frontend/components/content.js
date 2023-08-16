import React from "react";

export default function content() {
  return (
    <div className="container px-5 py-24 mx-auto">
      <div className="lg:w-4/5 mx-auto flex flex-wrap">
        <div className="lg:w-1/2 w-full lg:pr-10 lg:py-6 mb-6 lg:mb-0">
          <h2 className="text-sm title-font text-gray-500 tracking-widest">JICO</h2>
          <h2 className="text-gray-900 text-3xl title-font font-medium mb-4">WELCOME TO JICO</h2>
          <div className="flex mb-4">
            <span className="flex-grow text-black border-b-2 border-black py-2 text-lg px-1">Description</span>
          </div>
          <p className="leading-relaxed mb-4">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
          </p>
          <div className="flex">
            <a className="flex item-center text-white bg-blue-500 border-0 py-2 px-6 focus:outline-none hover:bg-blue-700 rounded" href="/registerPage" role="button">Get Started</a>
          </div>
        </div>
        <img alt="ecommerce" className="lg:w-1/2 w-ful lg:h-auto h-64 object-cover object-center rounded" src="./assets/oranglaptop.jpeg"/>
      </div>
    </div>
  )
}