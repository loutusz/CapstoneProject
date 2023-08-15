import React from 'react'
import Head from 'next/head'
import { FaRegEnvelope } from 'react-icons/fa'
import { useState } from 'react'

export const NotForgotPass = () => {
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
    <div>
        <div>
            <Head>
                <title>Reset your password</title>
            </Head>
        </div>
        <div className='bg-gradient-to-r from-cyan-500/10 via-teal-300/10 to-sky-200/10 block h-screen items-center justify-center p-4 md:flex'>
            <div className='bg-white flex flex-col items-center p-4 space-y-8 w-full md:w-3/5 text-slate-700 rounded-lg shadow-[0_3px_10px_rgb(0,0,0,1)]'>
                <img className='mt-10 h-24' src='https://img.freepik.com/free-icon/exclamation_318-579940.jpg?w=360'/>
                <h1 className='font-semibold text-3xl text-slate-700'>Forgot Password</h1>
                <p className='w-80 text-center'>Enter your registered email associated with your account.</p>
                <div className='relative py-2'>
                    <div className=' bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md flex items-center pl-4'>
                            <FaRegEnvelope className=' text-slate-700'/>
                            <input className='pl-4  py-1 w-72 h-10' type='email' name='email' value={formData.email} onChange={handleChange} placeholder='Email'/>
                    </div>
                </div>
                <div className='relative py-2'>
                    <button type="submit" className="bg-blue-700  text-zinc-100 text-center justify-center rounded-md hover:bg-blue-800 transition-colors duration-200 inline-flex items-center w-72 h-10 px-3 py-1">Submit</button>
                </div>
                <a href='/SignInPage' className='hover:underline decoration-slate-700 pb-10'>&lt; Back to Sign In</a>
            </div>
        </div>
    </div>
  )
}
export default NotForgotPass