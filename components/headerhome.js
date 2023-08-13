import { useState } from 'react';

const HeaderHome = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };

  return (
    <header className="py-4 px-6 flex justify-between items-center bg-blue-500">
      <button
        className="text-white text-xl focus:outline-none"
        onClick={toggleSidebar}
      >
        ☰
      </button>
      <div className="flex items-center space-x-4">
        <div className="text-white">Nama Pengguna</div>
        <img
          className="w-8 h-8 rounded-full"
          src="/path/to/user/avatar.jpg"
          alt="User Avatar"
        />
      </div>
      {isSidebarOpen && (
        <div className="fixed inset-0 bg-gray-800 bg-opacity-50 z-10">
          <div className="absolute left-0 h-full w-64 bg-blue-700 shadow-md">
            <nav className="p-4">
              <h2 className="text-xl font-semibold mb-4 text-white">Menu</h2>
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
