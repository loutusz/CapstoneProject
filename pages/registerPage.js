import React, {useState} from 'react';
import {FcGoogle, FcDiscord} from 'react-icons/fc'
import Link from "next/link";
import Head from 'next/head';

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

  
      return (
        
        <div className="flex flex-col justify-center items-center py-2 w-full h-full bg-slate-200">
          <div>
            <Head>
              <title>Register Page</title>
            </Head>
          </div>
          <form onSubmit={handleSubmit} className="flex flex-col items-center justify-center w-full flex-1 px-35 py-15 text-center" > 
            <div className="flex flex-col items-center justify-center w-full flex-1 px-35 py-15 text-center">
              <div className=" bg-gray-50  rounded-2xl flex w-2/3 max-w-4xl shadow-2xl">
       
              {/* Welcome */}
                <div className="text-black w-full md:w-2/5 bg-gradient-to-t from-sky-200 to-stone-50 rounded-tl-2xl rounded-bl-2xl p-10">
                  <h2 className="text-3xl font-bold mt-6 text-left">Welcome to JICO</h2>
                  <p className="text-xs mt-4 text-left">
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam nec ultricies nisi. 
                  Suspendisse pulvinar viverra nibh vel ultricies. Mauris tincidunt mollis diam,
                  at mollis enim aliquet eget. Fusce eros neque, pharetra eget tincidunt in,
                  tincidunt nec tellus. 
                  </p>
                  <div className='h-8 flex object-left-bottom'>
                        <img className='' src='https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png'/>
                        <img className='' src='https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png'/>
                    </div>
                </div>

              {/* SignUp */}
              {/* p-9 w-4/5 */}
                <div className="text-blackp-4 md:p-8 w-full md:w-3/5"> 
                  <h2 className="text-3xl font-semibold mb-3">Sign Up</h2>
                  
                  {/* Name */}
                     <div className="mt-6 mb-4 mx-auto w-4/5">
                        <label htmlFor="name" className="block mb-2 font-medium text-black"></label>
                        <input
                          type="text"
                          name="name"
                          placeholder="Name"
                          value={formData.name}
                          onChange={handleChange}
                          className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black"
                          required
                        />
                      </div>

                   {/* Username */}
                      <div className="mb-4 mx-auto w-4/5">
                        <label htmlFor="username" className="block mb-2 font-medium text-black"></label>
                        <input
                          type="text"
                          name="username"
                          placeholder="Username"
                          value={formData.username}
                          onChange={handleChange}
                          className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black"
                          required
                        />
                      </div>

                      {/* Email */}
                      <div className="mb-4 mx-auto w-4/5">
                        <label htmlFor="email" className="block mb-2 font-medium text-black"></label>
                        <input
                          type="email"
                          name="email"
                          placeholder="Email"
                          value={formData.email}
                          onChange={handleChange}
                          className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black"
                          required
                        />
                      </div>
                      
                    {/* Password */}
                    <div className="mb-4 mx-auto w-4/5">
                      <label htmlFor="password" className="block mb-2 font-medium text-black"></label>
                      <input
                        type="password"
                        name="password"
                        placeholder="Password"
                        value={formData.password}
                        onChange={handleChange}
                        className="w-full border-gray-300 border rounded-lg px-3 py-2 focus:outline-none shadow shadow-black"
                        required
                      />
                     </div>

                    {/* Checkbox */}
                    <div className="flex w-64 mb-5 items-center mx-auto">
                      <label htmlFor="remember" className="flex items-center mt-2 text-justify text-xs">
                        <input
                          type="checkbox"
                          id="remember"
                          name="remember"
                          className="mr-2"
                        />
                        I’ve read and agree with terms of service and our privacy policy
                      </label>                   
                    </div>

                     {/* Button Submit */}
                  <div className="text-center mb-3">
                    <button type="submit" className="bg-blue-500 text-white w-4/5 px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors duration-200">Register</button>
                  </div>

                  {/* Sign in */}
                  <div className="flex flex-col items-center mt-3">
                    <p>Already a member?
                      <a href='/loginPage' className="text-blue-700 font-semibold"> Sign In</a>
                    </p>                  
                  </div>
                  
                </div>
              </div>
            </div>
             </form>
        </div>
      );
    };
