import React from 'react';

export default function Sidebar() {
  return (
    <div className="bg-blue-700 text-white h-screen w-1/4 fixed top-0 left-0 p-4">
      <h2 className="text-lg font-semibold mb-4">Sidebar Title</h2>
      <nav>
        <ul>
          <li className="mb-2">
            <a href="#" className="text-white hover:text-blue-300">
              Projects
            </a>
          </li>
          <li className="mb-2">
            <a href="#" className="text-white hover:text-blue-300">
              Provider
            </a>
          </li>
          <li>
            <a href="#" className="text-white hover:text-blue-300">
              Settings
            </a>
          </li>
        </ul>
      </nav>
    </div>
  );
};
