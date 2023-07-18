import React, {useState} from 'react';
import Link from "next/link";

export default function RegisterPage() {
    const [formData, setFormData] = useState({
        name: '',
        username: '',
        email: '',
        password: ''
      });
    
      const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };
    
      const handleSubmit = (e) => {
        e.preventDefault();
        // Lakukan aksi pendaftaran atau validasi form di sini
        console.log(formData);
        // Reset form setelah pengiriman
        setFormData({
          name: '',
          username: '',
          email: '',
          password: ''
        });
      };
    
      const handleGoogleRegistration = () => {
        // Lakukan proses pendaftaran dengan Google di sini
        console.log('Melakukan pendaftaran dengan Google');
      };
    
      return (
        <div className="flex justify-center items-center h-screen">
          <form onSubmit={handleSubmit} className="max-w-md w-full px-6 py-8 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl font-semibold text-center mb-6">Form Pendaftaran</h2>
            <div className="mb-4">
              <label htmlFor="name" className="block mb-2 font-medium text-black">Nama</label>
              <input
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none"
                required
              />
            </div>
            <div className="mb-4">
              <label htmlFor="username" className="block mb-2 font-medium text-black">Username</label>
              <input
                type="text"
                name="username"
                value={formData.username}
                onChange={handleChange}
                className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none"
                required
              />
            </div>
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
              <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors duration-200">Register</button>
            </div>
            <div className="flex justify-center">
              <button type="button" onClick={handleGoogleRegistration} className="flex items-center border border-gray-300 rounded-lg px-4 py-2 hover:bg-gray-100 focus:outline-none">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
                <span className='text-black'>Register with Google</span>
              </button>
            </div>
          </form>
        </div>
      );
    };
