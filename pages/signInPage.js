import React from 'react'
import Head from 'next/head';
import { useState } from 'react';
import Link from 'next/link';
import { FaRegEnvelope } from 'react-icons/fa';
import { FcGoogle } from 'react-icons/fc';
import {MdLockOutline} from 'react-icons/md'
import { BiSolidHide } from "react-icons/bi";

export const notLoginPage = () => {
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

      const{show,showPass} = useState(false);
  
    return (
    <div>
        <div>
            <Head>
                <title>Sign In - Welcome to JICO</title>
            </Head>
        </div>
        {/* background */}
        <div className='bg-gradient-to-r from-cyan-500/10 via-teal-300/10 to-sky-200/10 block h-screen items-center justify-center p-4 md:flex'>
            {/* container */}
            <div className='bg-gradient-to-t from-sky-200 to-stone-50 flex flex-col items-center max-w-screen-lg overflow-hidden rounded-lg shadow-[0_3px_10px_rgb(0,0,0,1)] w-full md:flex-row text-white'>
                {/* form card */}
                <div className='bg-white flex flex-col items-center p-4 space-y-8 w-full md:w-3/5 text-slate-700'>
                    {/* TITLE Sign In */}
                    <div className='flex flex-col items-center'>
                        <h1 className=' text-slate-700 text-4xl font-semibold pt-14'>Sign in</h1>
                    </div>

                    {/* ISI FORM */}
                    <form onSubmit={handleSubmit} className='flex flex-col items-center space-y-4'>
                        <div className='relative py-2'>

                            {/* Username */}
                         <div className=' bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md flex items-center pl-4'>
                                <FaRegEnvelope className=' text-slate-700'/>
                                <input className='pl-4  py-1 w-72 h-10' type='username' name='username' value={formData.username} onChange={handleChange} placeholder='Username'/>
                            </div>
                        </div>

                        {/* Password */}
                        <div className='relative py-2'>
                            <div className=' bg-white shadow-[0_3px_10px_rgb(0,0,0,0.1)] rounded-md flex items-center pl-4'>
                                <MdLockOutline className=' text-slate-700'/>
                                <input className='pl-4 py-1 w-72 h-10' 
                                type='password' 
                                name='password' 
                                value={formData.password} 
                                onChange={handleChange} placeholder='Password'/>
                                <BiSolidHide className=""/>
                            </div>
                        </div>

                        {/* Submit */}
                        <div className='relative py-2'>
                            <button type="submit" className="bg-blue-700  text-zinc-100 text-center justify-center rounded-md hover:bg-blue-800 transition-colors duration-200 inline-flex items-center w-72 h-10 px-3 py-1">Sign In</button>
                        </div>
                    </form>
                    <div className='flex flex-col items-center'>
                        <label htmlFor="remember" className="flex items-center  text-justify text-xs">
                            <input type="checkbox" name="remember" className="mr-2"/>
                                Remember me
                        </label>
                    </div>
                    <div className='flex flex-col items-center pb-10'>
                        <a href='/NotForgotPass' className='text-slate-700 text-s font-semibold leading-snug hover:underline'>Forgot Password?</a>

                        <p className='SignUp text-slate-700 text-base font-normal leading-snug'>Not a member yet? <a href='/registerPage' className='text-blue-700 font-bold text-base leading-snug'>Sign Up</a></p>
                    </div>
                    <div className=''>
                        
                    </div>

                </div>

                {/* welcome card */}
                <div className='flex flex-col items-center justify-center text-right text-slate-700 p-4 w-full md:w-2/5'>
                    <div>

                        <h2 className='text-2xl font-bold mb-2'>Welcome to JICO</h2>
                        <p className='text-base font-normal mb-48'>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam nec ultricies nisi. Suspendisse pulvinar viverra nibh vel ultricies. Mauris tincidunt mollis diam, at mollis enim aliquet eget. Fusce eros neque, pharetra eget tincidunt in, tincidunt nec tellus. </p>

                        <div className='h-8 flex object-right-bottom'>
                        <img className='' src='https://upload.wikimedia.org/wikipedia/id/thumb/c/c4/Telkom_Indonesia_2013.svg/1200px-Telkom_Indonesia_2013.svg.png'/>
                        <img className='' src='https://cdn.icon-icons.com/icons2/2699/PNG/512/atlassian_jira_logo_icon_170512.png'/>
                        </div>
                    </div>
                </div>
                
            </div>        
        </div>
    </div>
    
  )
}
export default notLoginPage ;