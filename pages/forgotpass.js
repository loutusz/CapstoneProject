import { useState } from "react";
import Header from "@/components/header";

export default function Forgotpass() {
    const [formData, setFormData] = useState({
        email: '',
        newPassword: '',
        confirmPassword: ''
      });
    
      const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };
    
      const handleSubmit = (e) => {
        e.preventDefault();
        // Lakukan aksi reset password di sini
        console.log(formData);
        // Reset form setelah pengiriman
        setFormData({
          email: '',
          newPassword: '',
          confirmPassword: ''
        });
      };
    
      return (
        <div>
            <Header/>
        <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-zinc-100">
          <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
            <form onSubmit={handleSubmit} className="w-full max-w-md bg-white rounded-lg shadow-md p-8">
              <h2 className="text-2xl font-semibold mb-4">Forgot Password</h2>
              <div className="mb-4">
                <input
                  className="w-full px-3 py-2 placeholder-gray-400 border rounded-lg focus:outline-none"
                  type="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                  required
                  placeholder="Email"
                />
              </div>
              <div className="mb-4">
                <input
                  className="w-full px-3 py-2 placeholder-gray-400 border rounded-lg focus:outline-none"
                  type="password"
                  name="newPassword"
                  value={formData.newPassword}
                  onChange={handleChange}
                  required
                  placeholder="New Password"
                />
              </div>
              <div className="mb-4">
                <input
                  className="w-full px-3 py-2 placeholder-gray-400 border rounded-lg focus:outline-none"
                  type="password"
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  required
                  placeholder="Confirm Password"
                />
              </div>
              <div className="flex items-center justify-center">
                <button
                  type="submit"
                  className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-lg focus:outline-none"
                >
                  Reset Password
                </button>
              </div>
            </form>
          </main>
        </div>
        </div>

      );
}