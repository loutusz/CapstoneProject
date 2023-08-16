import { useState } from 'react';
import Image from 'next/image';
import { MdClose } from 'react-icons/md';

const HeaderHome = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };

  return (
    <header className="py-4 px-6 flex justify-between items-center bg-blue-700">
      <button
        className="text-white text-xl focus:outline-none"
        onClick={toggleSidebar}
      >
        ☰
      </button>
      <div className="flex items-center space-x-4">
        <div className="text-white">User</div>
        <Image
          className="rounded-full"
          width={20}
          height={20}
          src="/path/to/user/avatar.jpg"
          alt="User Avatar"
        />
      </div>
      {isSidebarOpen && (
        <div className="fixed inset-0 bg-gray-800 bg-opacity-50 z-10">
          <div className="absolute left-0 h-full w-64 bg-blue-800 shadow-md">
            <nav className="p-4">
              <div className="flex items-enter justify-between mb-4 text-white">
                <h2 className="text-xl font-semibold">Menu</h2>
                <MdClose
                  className="text-xl cursor-pointer"
                  onClick={toggleSidebar}
                />
              </div>
              <ul className="space-y-2">
                <li>
                  <a href="#" className="text-white hover:underline">Menu 1</a>
                </li>
                <li>
                  <a href="#" className="text-white hover:underline">Menu 2</a>
                </li>
                <li>
                  <a href="#" className="text-white hover:underline">Menu 3</a>
                </li>
              </ul>
            </nav>
          </div>
        </div>
      )}
    </header>
  );
};

export default HeaderHome;
