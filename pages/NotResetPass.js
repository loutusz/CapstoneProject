import React from 'react'
import Head from 'next/head'
import {MdLockOutline} from 'react-icons/md'
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
                <h1 className='font-semibold text-3xl text text-slate-700'>Reset Password</h1>
                <p className='w-80 text-center'>Set your new Password that you will remember!!</p>
                <p className='w-56 text-sm text-center font-semibold'>Note : It must be different from your previous password</p>
                <div className='relative py-2'>
                    <div className=' bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md flex items-center pl-4 mb-6'>
                        <MdLockOutline className=' text-slate-700'/>
                        <input className='pl-4  py-1 w-72 h-10' type='password' name='password' value={formData.password} onChange={handleChange} placeholder='New Password'/>
                    </div>
                    <div className=' bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md flex items-center pl-4'>
                        <MdLockOutline className=' text-slate-700'/>
                        <input className='pl-4  py-1 w-72 h-10' type='password' name='password' value={formData.password} onChange={handleChange} placeholder='Confirm Password'/>
                    </div>
                </div>
                
                <div className='relative py-2 pb-10'>
                    <button type="submit" className="bg-blue-700  text-zinc-100 text-center justify-center rounded-md hover:bg-blue-800 transition-colors duration-200 inline-flex items-center w-72 h-10 px-3 py-1">Reset Password</button>
                </div>
            </div>
        </div>
    </div>
  )
}
export default NotForgotPass