import React, {useState} from "react";
import Link from "next/link";

export default function LoginPage() {
    const [formData, setFormData] = useState({
        email: '',
        password: ''
      });
    
      const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };
    
      const handleSubmit = (e) => {
        e.preventDefault();
        // Lakukan aksi login atau validasi form di sini
        console.log(formData);
        // Reset form setelah pengiriman
        setFormData({
          email: '',
          password: ''
        });
      };
    
      const handleGoogleLogin = () => {
        // Lakukan proses login dengan Google di sini
        console.log('Melakukan login dengan Google');
      };
    
      return (
        <div className="flex justify-center items-center h-screen">
          <form onSubmit={handleSubmit} className="max-w-md w-full px-6 py-8 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl font-semibold text-center mb-6">Login</h2>
            <div className="mb-4">
              <label htmlFor="email" className="block mb-2 font-medium text-black">Email</label>
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none"
                required
              />
            </div>
            <div className="mb-4">
              <label htmlFor="password" className="block mb-2 font-medium text-black">Password</label>
              <input
                type="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none"
                required
              />
            </div>
            <div className="text-center mb-4">
              <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors duration-200">Login</button>
            </div>
            <div className="flex justify-center">
              <button type="button" onClick={handleGoogleLogin} className="flex items-center border border-gray-300 rounded-lg px-4 py-2 hover:bg-gray-100 focus:outline-none">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
                <span className="text-black">Login with Google</span>
              </button>
            </div>
          </form>
        </div>
      );
}